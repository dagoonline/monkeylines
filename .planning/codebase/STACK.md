# Technology Stack

**Analysis Date:** 2026-02-07

## Languages

**Primary:**
- Go 1.21 - HTTP server, message generation, and asset embedding

**Secondary:**
- HTML/CSS - Web interface styling and layout
- JavaScript (Vanilla) - Web Audio API integration, DOM manipulation, and interactive behavior

## Runtime

**Environment:**
- Go 1.21+ (required for building)

**Package Manager:**
- go (built-in with Go)
- Lockfile: Not applicable (go.mod with no external dependencies)

## Frameworks

**Core:**
- Go standard library `net/http` - HTTP server, routing, and handlers
- Go standard library `html/template` - HTML template parsing and rendering

**Build/Dev:**
- Docker - Multi-stage build using `golang:1.21-alpine` and `alpine:latest`

## Key Dependencies

**Critical:**
- None - Project uses only Go standard library (`bytes`, `context`, `encoding/json`, `fmt`, `html/template`, `log`, `net`, `os`, `strings`, `syscall`, `time`)

**Infrastructure:**
- None - All assets (HTML, images) embedded at compile time using `go:embed`

## Configuration

**Environment:**
- `MONKEYLINES_HTTP_PORT` - HTTP server port (default: `8080`)
- Located in: `main.go` line 101, function `getEnv()`

**Build:**
- `Dockerfile` - Multi-stage build configuration
- `go.mod` - Go module definition (minimal, no external dependencies)

## Platform Requirements

**Development:**
- Go 1.21 or later
- Unix-like environment (bash scripts in GitHub Actions)

**Production:**
- Alpine Linux base image (lightweight container deployment)
- Runs as single binary with embedded assets
- No external runtime dependencies

---

*Stack analysis: 2026-02-07*
