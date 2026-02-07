# Codebase Concerns

**Analysis Date:** 2026-02-07

## Test Coverage Gaps

**No automated tests:**
- What's not tested: All core functionality (message generation, HTTP handlers, template rendering, client IP extraction)
- Files: `main.go`, `messages.go`
- Risk: Changes to message generation logic, server routing, or IP extraction can break unnoticed. Particularly risky for the `clientIP()` function which handles proxy headers.
- Priority: High

**Browser/Client-side code untested:**
- What's not tested: Audio context initialization, event handling, state management, async exchange flow
- Files: `index.html` (lines 161-378)
- Risk: Mobile audio unlocking workaround (lines 186-191) is fragile and could silently fail on new browser versions. Auto-play logic (lines 359-368) lacks verification.
- Priority: Medium

## Security Considerations

**CSP allows unsafe-inline scripts:**
- Risk: While CSP includes `default-src 'self'`, script-src and style-src both allow `'unsafe-inline'` (line 114 in `main.go`). This weakens XSS protection for inline scripts embedded in HTML.
- Files: `main.go` (line 114), `index.html` (lines 20-142)
- Current mitigation: Template is compiled at build time and cannot be modified at runtime; embedded assets prevent path traversal
- Recommendations: If HTML must remain inline, consider migrating to external stylesheets and script files with stricter CSP. Add `nonce` attributes for inline scripts if staying with current approach.

**X-Forwarded-For header parsing without validation:**
- Risk: `clientIP()` function trusts X-Forwarded-For and X-Real-IP headers without rate limiting or sanitization. Could be exploited to log false IPs or cause logging spikes.
- Files: `main.go` (lines 63-78)
- Current mitigation: Only first IP extracted from comma-separated list; falls back to RemoteAddr
- Recommendations: Add IP validation (must be valid CIDR or known proxy), consider rate limiting on repeated suspicious IPs, or add `X-Forwarded-For` validation against known proxy IPs only

**No validation on Exchange/Line output:**
- Risk: While message templates are hardcoded, there's no runtime validation that generated strings don't exceed reasonable bounds. Extremely long insults could cause browser/client issues.
- Files: `messages.go` (lines 191-207)
- Current mitigation: Template system is closed (themes defined in code)
- Recommendations: Add maximum string length checks before serving, or document and monitor max generation size

## Performance Bottlenecks

**Random number generation not seeded per-request:**
- Problem: `rand.Intn()` is called directly without explicit seeding in `messages.go`
- Files: `messages.go` (lines 188-189)
- Cause: Go's default random source is seeded once at startup. In a heavily loaded service, this could create predictable patterns or collisions across concurrent requests.
- Improvement path: Seed `rand.Seed(time.Now().UnixNano())` per-request, or use `rand/v2` (Go 1.22+) which has better concurrency properties

**Web Audio API initialization repeated on button clicks:**
- Problem: `ensureAudio()` called on every button click (lines 355, 360, 372 in `index.html`)
- Files: `index.html` (lines 179-192)
- Cause: No caching of AudioContext state across event handlers. Each click re-enters the audio setup logic.
- Improvement path: Initialize AudioContext once on page load rather than on first interaction

## Fragile Areas

**Mobile audio unlock workaround:**
- Files: `index.html` (lines 186-191)
- Why fragile: Playing a silent buffer to unlock AudioContext is a browser-specific workaround. Safari, Chrome, and Firefox handle AudioContext state differently. Future browser updates could invalidate this approach without warning.
- Safe modification: Test on iOS Safari, Android Chrome, and Android Firefox after any change. Add feature detection to verify buffer playback works.
- Test coverage: Untested - requires manual testing on actual mobile devices

**Text fitting algorithm:**
- Files: `index.html` (lines 260-267)
- Why fragile: `fitText()` reduces font size by 0.5px in a loop while `el.scrollHeight > el.clientHeight`. This can cause excessive reflows. Changing padding, margin, or speech bubble dimensions breaks the algorithm.
- Safe modification: Use CSS `max-width` and `overflow` instead of JavaScript manipulation. Current approach defeats browser layout optimization.
- Test coverage: No unit tests; only manual visual verification

**Template rendering with user-derived baseURL:**
- Files: `main.go` (lines 48-58)
- Why fragile: `baseURL` computed from `r.Host` and `scheme()` is passed to Go template. While not directly injectable (HTML escaping in use), r.Host could contain unusual values from proxies or malformed requests.
- Safe modification: Validate r.Host against a whitelist for known deployment domains or use a config file instead of deriving from request
- Test coverage: No tests for malformed Host headers

## Scaling Limits

**No connection pooling or rate limiting:**
- Current capacity: Single Go process handles all concurrent connections
- Limit: Server has ReadTimeout=10s, WriteTimeout=10s, but no per-IP rate limiting. A single client hammering `/exchange` can tie up goroutines.
- Scaling path: Add middleware to rate limit per IP (token bucket), use `github.com/prometheus/client_golang` for metrics, deploy multiple instances behind load balancer

**Client-side hooting audio generation**:
- Current capacity: Each hooting sound spawns oscillators and gain nodes
- Limit: On slow devices or with many simultaneous words, AudioContext could become resource-constrained. No throttling of audio node creation.
- Scaling path: Add audio context node pooling, limit max hoots per word, add device capability detection

**Embedded assets grow with image size:**
- Current capacity: `monkeylines` binary is ~18MB (mostly embedded images)
- Limit: Each image added increases binary size linearly. Container image size compounds this.
- Scaling path: If adding more images, consider lazy-loading from CDN; keep embedded assets for critical path only

## Dependencies at Risk

**Go 1.21 requirement:**
- Risk: Go 1.21 was released August 2023. Security fixes may not apply to earlier versions. No upper bound on Go version used to build.
- Impact: If a security fix ships in Go 1.23+, builds on Go 1.21 won't include it
- Migration plan: Upgrade to Go 1.22+ (released Feb 2024), verify all tests pass, update go.mod and Dockerfile

**No dependency lock mechanism for images:**
- Risk: Dockerfile uses `alpine:latest` which is a floating tag
- Impact: Multi-stage builds could produce different images if new Alpine versions introduce breaking changes
- Migration plan: Pin to specific Alpine version (e.g., `alpine:3.19`), review and rotate every 6 months

## Known Issues

**Auto-play state not preserved on page reload:**
- Symptoms: If user enables auto-play (🔁) then refreshes page, auto-play is off
- Files: `index.html` (lines 175-176 - state is local variable, no localStorage)
- Workaround: User must click 🔁 again
- Severity: Low - expected behavior, but poor UX for long-running sessions

**No error handling for `/exchange` fetch failures:**
- Symptoms: If server is down, auto-play stops silently and buttons re-enable without user feedback
- Files: `index.html` (lines 343-350)
- Workaround: Buttons become clickable again but silently fail on retry
- Severity: Medium - users unaware of service degradation

**Server timeouts not configurable:**
- Symptoms: Hard-coded 10s read/write timeouts cannot be adjusted without code change
- Files: `main.go` (lines 133-136)
- Workaround: Rebuild binary with new timeouts
- Severity: Low - reasonable defaults, but inflexible for different deployments

## Missing Critical Features

**No health check endpoint:**
- Problem: Fly.io and other orchestrators need a dedicated health check
- Blocks: Proper rolling deployments, automatic instance restart on service death
- Current state: `/` and `/exchange` could be used but they generate random content (not idempotent), making health checks unreliable

**No metrics or logging for monitoring:**
- Problem: No Prometheus metrics, structured logs, or error tracking integration
- Blocks: Observability into error rates, response times, audio failures, or repeated IP patterns
- Current state: Server logs to stdout (good for containers) but no differentiation between INFO, WARN, ERROR levels

**No input validation or contract testing:**
- Problem: No validation that theme data structures are well-formed at compile time or startup
- Blocks: Early detection if themes are misconfigured or nouns become nil
- Current state: Relies on Go's initialization; would panic at runtime if data is malformed

---

*Concerns audit: 2026-02-07*
