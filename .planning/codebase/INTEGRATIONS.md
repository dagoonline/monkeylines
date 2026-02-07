# External Integrations

**Analysis Date:** 2026-02-07

## APIs & External Services

**None detected** - This application is self-contained and does not integrate with external APIs.

## Data Storage

**Databases:**
- Not applicable - Application is stateless and generates data procedurally

**File Storage:**
- Local filesystem only - Static assets (images, HTML) embedded in binary via `go:embed` directive in `main.go`
- Assets located in: `/images/` directory (embedded at compile time)

**Caching:**
- HTTP header-based caching only
- Static image assets cached for one week: `Cache-Control: public, max-age=604800` (line 123 in `main.go`)

## Authentication & Identity

**Auth Provider:**
- None - Application is public and requires no authentication

## Monitoring & Observability

**Error Tracking:**
- None detected

**Logs:**
- Standard Go `log` package (line 11 in `main.go`)
- HTTP request logging: `log.Printf()` calls record method, scheme, host, path, client IP, and served content
- Errors logged to stdout via `log.Fatalf()` and `log.Printf()` (lines 31-32, 39, 52, 145-146, 155-156 in `main.go`)

## CI/CD & Deployment

**Hosting:**
- Fly.io
- Configuration: `fly.toml` (lines 1-9)
- Primary region: cdg (Paris)
- Force HTTPS enabled
- Auto-start/stop machines configured
- Internal port: 8080

**CI Pipeline:**
- GitHub Actions
- Workflow file: `.github/workflows/deploy.yml`
- Triggers: Push to `main` or `master` branches
- Action: Deploys via flyctl on remote Fly.io infrastructure
- Requires: `FLY_API_TOKEN` secret in repository settings

**Secrets Required:**
- `FLY_API_TOKEN` - GitHub repository secret for Fly.io deployment authentication

## Environment Configuration

**Required env vars:**
- `MONKEYLINES_HTTP_PORT` - HTTP server port (default: `8080`)
  - Read in: `main.go` line 101, function `getEnv()`

**Secrets location:**
- GitHub repository secrets (for `FLY_API_TOKEN`)
- No `.env` files used; environment variables passed at runtime

## Webhooks & Callbacks

**Incoming:**
- None - Application does not accept webhooks

**Outgoing:**
- None - Application does not make outbound API calls to external services

## Client-Side Integrations

**Web Audio API:**
- Used in `index.html` lines 179-218
- Browser's native AudioContext for hooting sound generation
- No external audio library dependency
- Features: Oscillator synthesis, gain control, frequency modulation
- State: Initialized on first user interaction to support mobile browser requirements

---

*Integration audit: 2026-02-07*
