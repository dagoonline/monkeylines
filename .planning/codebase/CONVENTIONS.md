# Coding Conventions

**Analysis Date:** 2026-02-07

## Naming Patterns

**Files:**
- Go source files: lowercase with underscores: `main.go`, `messages.go`
- No file extensions in package declarations

**Functions:**
- Exported functions: PascalCase (Go standard): `generateExchange()`, `generateMessage()`
- Unexported functions: camelCase: `clientIP()`, `scheme()`, `handleHTTP()`
- HTTP handlers follow pattern `handleX()` where X is descriptive: `handleExchange()`, `handlePlain()`, `handleHTTP()`

**Variables:**
- Local variables: camelCase: `tmpl`, `ex`, `buf`, `baseURL`
- Constants: UPPERCASE: typically used implicitly with timeouts (`10 * time.Second`)
- Receiver variables in methods: single letters (Go convention): `w http.ResponseWriter`, `r *http.Request`

**Types:**
- Structs: PascalCase: `exchange` (unexported)
- Type names map to exported/unexported convention: `theme` (unexported), capitalized for exported

## Code Style

**Formatting:**
- Go code is gofmt-formatted (standard Go formatter)
- Standard Go indentation: tabs (Go convention)
- Line length: implicit adherence to Go standards (typically ~80-100 chars practically)

**Linting:**
- No explicit linting configuration detected
- Code follows Go idioms and best practices observed in stdlib patterns
- No `.golangci.yml` or similar linting config present

## Import Organization

**Order:**
1. Standard library imports (grouped): `"bytes"`, `"context"`, `"encoding/json"`, `"fmt"`, `"html/template"`, `"log"`, `"net"`, `"net/http"`, `"os"`, `"os/signal"`, `"strings"`, `"syscall"`, `"time"`
2. External imports (none in this codebase)

**Path Aliases:**
- No aliases used in current codebase
- Imports use full standard package paths

## Error Handling

**Patterns:**
- Explicit error checking with `if err != nil` pattern (Go standard): `if err := httpServer.Shutdown(shutdownCtx); err != nil { log.Printf(...) }`
- Some errors logged as warnings, not fatal: HTTP shutdown errors
- Fatal errors use `log.Fatal()`: `log.Fatal("Failed to parse template:", err)`
- Some errors silently ignored or deferred: audio context errors in JavaScript (`catch(e) { return; }`)

**HTTP Error Handling:**
- Uses `http.NotFound(w, r)` for 404s
- Uses `http.Error()` for 500s with descriptive messages
- Errors in template rendering logged and responded with `http.StatusInternalServerError`

## Logging

**Framework:** Standard Go `log` package

**Patterns:**
- All server operations logged with `log.Printf()`: startup messages, HTTP requests
- Logging includes: HTTP method, scheme, host, path, client IP, operation result
- Log lines prefixed with timestamp and level by default (Go log package)
- Decorative logging for visual separation: `log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")`
- Client IP extraction logged during request handling
- Template rendering failures logged with context

**JavaScript:**
- No explicit logging in JavaScript code (index.html)
- Browser console available implicitly but not used

## Comments

**When to Comment:**
- Comments explain WHY, not WHAT: `// X-Forwarded-For may contain a comma-separated list; first entry is the client`
- Used for non-obvious logic: IP parsing from headers
- Used for type explanations: `// A theme groups insult/comeback templates that are topically related.`
- Not overused; most code is self-documenting

**JSDoc/TSDoc:**
- Not used in this codebase
- Go lacks formal doc comment structure but single-line comments used
- Comments explain intent and edge cases

## Function Design

**Size:**
- Functions are concise and focused
- HTTP handlers (`handleHTTP`, `handleExchange`, `handlePlain`): 10-30 lines
- Helper functions (`clientIP`, `scheme`, `getEnv`): 5-15 lines
- Generator functions (`generateExchange`, `generateMessage`): 10-15 lines

**Parameters:**
- HTTP handlers follow Go convention: `(w http.ResponseWriter, r *http.Request)`
- Helper functions take minimal parameters
- No variadic functions used
- Generic function used for array choice: `randomChoice[T any](slice []T)`

**Return Values:**
- Single return values for simple operations
- Error returns for operations that can fail (template parsing, template execution)
- Named returns not used in this codebase
- Implicit returns for template execution: `json.NewEncoder(w).Encode(ex)` writes directly to response

## Module Design

**Exports:**
- Only functions and types that need external use are exported (capitalized)
- Unexported functions handle internal logic: `clientIP()`, `scheme()`, `randomChoice()`
- Unexported types: `theme` struct containing theme data
- Exported types: `exchange` struct (used for JSON encoding to clients)

**Barrel Files:**
- Not applicable; Go packages don't use barrel exports
- All code in `main` package uses direct imports

## Go-Specific Conventions

**Defer Pattern:**
- Used for cleanup: `defer stop()`, `defer cancel()`, `defer buf.WriteTo(w)`
- Ensures operations complete even if early returns occur

**Context Usage:**
- `context.Background()` for root context
- `signal.NotifyContext()` for graceful shutdown
- `context.WithTimeout()` for operation timeouts

**Goroutine Handling:**
- Used for async operations: `go func() { httpServer.ListenAndServe() }()`
- Channels used for signaling: `<-ctx.Done()`

**Template Embedding:**
- `//go:embed` directives for embedding static files: `index.html`, `images/*`
- Embedded files served directly without file I/O

## JavaScript Conventions (index.html)

**Naming:**
- camelCase for variables: `sideMonkeys`, `autoPlay`, `monkeyVoices`
- Single-letter IDs for monkeys: `l`, `c`, `r`
- Prefixed function names: `showOverlay()`, `playHoot()`, `typeText()`

**Style:**
- Inline `<style>` tag with modern CSS (CSS Grid, Flexbox)
- CSS uses kebab-case for class names: `monkey-scene`, `speech-text`, `dialogue-btn`
- HTML uses semantic attributes: `aria-label`

**Error Handling:**
- Try-catch for audio context creation: `try { audioCtx = new ... } catch(e) { return; }`
- Graceful degradation if audio fails
- Promise-based flow with `.then()` chains

---

*Convention analysis: 2026-02-07*
