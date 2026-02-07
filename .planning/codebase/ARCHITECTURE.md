# Architecture

**Analysis Date:** 2026-02-07

## Pattern Overview

**Overall:** Single-server embedded assets pattern with stateless HTTP API

**Key Characteristics:**
- Monolithic Go server with no external dependencies
- All static assets (HTML, images) embedded at compile time via `go:embed`
- Stateless request handling with template rendering
- Procedural content generation (no databases or persistent state)
- Security-hardened with strict CSP and timeout controls

## Layers

**HTTP Server Layer:**
- Purpose: Handle incoming HTTP requests and manage lifecycle
- Location: `main.go` lines 94-160
- Contains: Server initialization, middleware, shutdown logic
- Depends on: Go standard library `net/http`
- Used by: All client requests

**Request Handler Layer:**
- Purpose: Route requests to appropriate handlers and coordinate responses
- Location: `main.go` lines 28-61
- Contains: Three handler functions (`handleExchange`, `handlePlain`, `handleHTTP`)
- Depends on: Message generation layer, template engine
- Used by: HTTP server via ServeMux

**Message Generation Layer:**
- Purpose: Generate themed insult/comeback pairs deterministically
- Location: `messages.go` lines 1-207
- Contains: Theme definitions, randomization logic, template substitution
- Depends on: Go standard library `math/rand`, `fmt`
- Used by: Request handlers

**Template & Asset Rendering Layer:**
- Purpose: Parse and serve embedded HTML with dynamic base URL injection
- Location: `main.go` lines 26, 96-97 (parsing), `index.html` (template)
- Contains: HTML template with inline CSS and JavaScript, image references
- Depends on: Go `html/template` package, embedded `index.html`
- Used by: `handleHTTP` handler

**Security Middleware Layer:**
- Purpose: Enforce security headers on all responses
- Location: `main.go` lines 109-117
- Contains: Header injection function that wraps the main router
- Depends on: HTTP middleware pattern
- Used by: HTTP server wraps all handlers

## Data Flow

**Web Page Load:**

1. Client requests `/`
2. `handleHTTP` receives request
3. Template engine parses embedded `index.html` and injects `{{.}}` (base URL)
4. HTML with inline CSS/JS is written to response
5. Client browser executes JavaScript (Immediately Invoked Function Expression)

**Insult/Comeback Generation (Exchange API):**

1. Client clicks ⚔️ button or auto-play triggers fetch to `/exchange`
2. `handleExchange` calls `generateExchange()`
3. `generateExchange()` selects random theme from 8 available themes
4. For each noun position, random noun from theme selected via `randomChoice()`
5. Template strings from theme substituted with selected nouns via `fmt.Sprintf`
6. `exchange` struct marshaled to JSON and returned
7. Client receives JSON, animates monkey, types text, plays hooting audio

**Plain Text Generation (Line API):**

1. Client requests `/line` (typically via curl or programmatic access)
2. `handlePlain` calls `generateMessage()`
3. `generateMessage()` generates exchange, randomly picks Insult or Comeback
4. Plain text returned with Content-Type `text/plain`

**Static Asset Serving:**

1. Client requests `/images/*` (favicon.ico, base.png, l.png, c.png, r.png, preview.jpg)
2. FileServer handles embedded `images` FS
3. Cache-Control header set to one week (604800 seconds)
4. Image served from embedded filesystem

**State Management:**

- No server-side state; each request is independent
- Client-side state in JavaScript: `busy`, `autoPlay`, `muted`, `audioCtx`
- Randomization seeded by Go runtime (not explicitly seeded, uses default)
- Request headers logged but not persisted

## Key Abstractions

**Theme:**
- Purpose: Groups thematically-related insults, comebacks, and nouns
- Examples: `messages.go` lines 16-185 (8 theme definitions)
- Pattern: Struct with three string slices (`insults`, `comebacks`, `nouns`)

**Exchange:**
- Purpose: Represents a paired insult and comeback
- Examples: `messages.go` lines 8-11
- Pattern: Simple struct with two exported string fields (`Insult`, `Comeback`)

**Client IP Detection:**
- Purpose: Extract true client IP from request, accounting for proxies
- Examples: `main.go` lines 63-78 (`clientIP` function)
- Pattern: Checks X-Forwarded-For, X-Real-IP, then falls back to RemoteAddr

**Protocol Detection:**
- Purpose: Determine if request came over HTTPS
- Examples: `main.go` lines 80-85 (`scheme` function)
- Pattern: Checks TLS presence or X-Forwarded-Proto header

## Entry Points

**Main Function:**
- Location: `main.go` lines 94-160
- Triggers: Server startup via `go run` or binary execution
- Responsibilities:
  - Parse environment variables (HTTP port)
  - Initialize HTML template from embedded asset
  - Create HTTP server with timeouts and header limits
  - Set up middleware chain (security headers)
  - Register request handlers to routes
  - Implement graceful shutdown on SIGINT/SIGTERM

**HTTP Handlers:**
- `/` (`handleHTTP` - `main.go` lines 42-61): Serves rendered HTML page
- `/line` (`handlePlain` - `main.go` lines 35-40): Serves plain text insult/comeback
- `/exchange` (`handleExchange` - `main.go` lines 28-33): Serves JSON exchange pair
- `/images/*`: Served by FileServer from embedded filesystem

**Client-Side Entry Point:**
- Location: `index.html` lines 161-378 (IIFE script block)
- Triggers: Browser page load, user button clicks
- Responsibilities:
  - Initialize audio context
  - Bind click handlers to buttons
  - Fetch and animate exchanges
  - Manage monkey voice generation
  - Handle UI state (busy, autoPlay, muted)

## Error Handling

**Strategy:** Graceful degradation with informative logging

**Patterns:**

- **Template parsing failure** (`main.go` lines 96-99): Fatal error at startup, prevents server from running with broken template
- **Handler errors** (`main.go` lines 51-54): HTTP 500 response with generic message, detailed error logged
- **Shutdown timeout** (`main.go` lines 152-157): Logs error but doesn't block—server stops either way
- **Client-side fetch failure** (`index.html` lines 343-350): Catches promise rejection, resets UI state, user can retry

## Cross-Cutting Concerns

**Logging:** Structured text logging to stdout via `log` package
- Format: Timestamp auto-prefixed by `log` package
- Logged events: Server start, handler requests (method, scheme, host, path, client IP, operation)
- Location: `main.go` lines 32, 39, 60, 103-105, 140-141, 150-151

**Validation:** Minimal; content generated from static theme data
- Input validation: Only the HTTP path is checked (`main.go` line 43, only `/` is valid HTML handler)
- Invalid paths result in 404 response via `http.NotFound`
- No user-supplied data reaches template engine

**Security:**
- Content Security Policy: `default-src 'self'; script-src 'unsafe-inline'; style-src 'unsafe-inline'; img-src 'self'`
- X-Content-Type-Options: `nosniff`
- X-Frame-Options: `DENY`
- Referrer-Policy: `no-referrer`
- Template rendered safely via `html/template` (not `text/template`)
- Client IP detection accounts for proxy headers
- Server timeouts: 5s read-header, 10s read, 10s write, 60s idle

**Audio Handling (Client-Side):**
- Web Audio API context lazily initialized on first user interaction
- Requires user gesture to unlock audio on iOS/Android (`ensureAudio()` at `index.html` lines 179-192)
- Silent buffer playback to unlock AudioContext state on mobile
- Per-monkey voice frequency ranges defined at `index.html` lines 193-197

---

*Architecture analysis: 2026-02-07*
