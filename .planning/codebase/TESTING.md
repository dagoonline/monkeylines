# Testing Patterns

**Analysis Date:** 2026-02-07

## Test Framework

**Runner:**
- Go: `go test` (standard Go testing package)
- Config: `go.mod` specifies Go 1.21; no separate test config file
- No external testing framework detected

**Assertion Library:**
- Go: Standard `testing` package with manual assertions
- JavaScript: No test framework or assertions detected

**Run Commands:**
```bash
go test ./...           # Run all tests (if tests exist)
go test -v ./...        # Verbose test output
go test -cover ./...    # Show coverage
```

## Test File Organization

**Location:**
- **Status:** No test files currently present in codebase
- **Pattern:** Go convention uses `_test.go` suffix for test files
- **Expected Location:** Tests would be co-located with source files:
  - `main_test.go` alongside `main.go`
  - `messages_test.go` alongside `messages.go`

**Naming:**
- Go tests follow pattern: `Test[FunctionName](t *testing.T)`
- Test helper functions: `Test[FunctionName]_[Scenario](t *testing.T)`

**Structure:**
```
/home/david/Projects/Github/dagoonline/monkeylines/
├── main.go
├── main_test.go          # (Does not exist; would go here)
├── messages.go
├── messages_test.go      # (Does not exist; would go here)
└── index.html
```

## Test Structure

**Suite Organization:**
```go
// Example pattern (not currently in codebase):
func TestHandleExchange(t *testing.T) {
    // Arrange
    req, err := http.NewRequest("GET", "/exchange", nil)
    if err != nil {
        t.Fatal(err)
    }
    w := httptest.NewRecorder()

    // Act
    handleExchange(w, req)

    // Assert
    if w.Code != http.StatusOK {
        t.Errorf("expected 200, got %d", w.Code)
    }
}
```

**Patterns:**
- Setup: Manual initialization in test functions or `TestMain(m *testing.M)` for shared setup
- Teardown: Deferred cleanup with `defer`
- Assertion: Manual comparison with `if condition != expected { t.Errorf(...) }`
- Subtests: `t.Run("subtest name", func(t *testing.T) { ... })`

## Mocking

**Framework:**
- Go: No external mocking library detected
- Manual mocking using standard library: `httptest.NewRecorder()`, `httptest.NewRequest()`

**Patterns:**
```go
// Example pattern (not currently in codebase):
// Using httptest for HTTP handler testing
func TestHandleHTTP(t *testing.T) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest("GET", "/", nil)
    handleHTTP(w, r)

    if w.Code != http.StatusOK {
        t.Errorf("expected 200, got %d", w.Code)
    }
}
```

**What to Mock:**
- External HTTP handlers using `httptest` package
- Template rendering should use actual embedded template or separate test template
- Client IP extraction should test various header combinations

**What NOT to Mock:**
- Standard library functions (use them directly)
- Template parsing (test with actual templates)
- Random number generation (if testing randomness, seed predictably)

## Fixtures and Factories

**Test Data:**
```go
// Example pattern (not currently in codebase):
func createTestTheme() theme {
    return theme{
        insults:   []string{"Test insult %s"},
        comebacks: []string{"Test comeback %s"},
        nouns:     []string{"test noun"},
    }
}

func createTestExchange() exchange {
    return exchange{
        Insult:   "You code like a rubber duck!",
        Comeback: "How appropriate. You fight like a cow.",
    }
}
```

**Location:**
- Fixtures: `*_test.go` files, typically defined in `TestMain()` or helper functions
- Shared fixtures could go in `test_helpers.go` or similar

## Coverage

**Requirements:**
- Not enforced (no `.github/workflows/coverage.yml` or similar)
- No codecov or coveralls integration detected

**View Coverage:**
```bash
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Test Types

**Unit Tests:**
- **Scope:** Individual functions in isolation
- **Approach:** Test handlers, generators, and utility functions separately
- **Examples to write:**
  - `TestGenerateExchange()` - validates exchange struct creation
  - `TestGenerateMessage()` - ensures message selection works
  - `TestClientIP()` - tests IP extraction from various header combinations
  - `TestScheme()` - validates HTTPS detection logic
  - `TestGetEnv()` - tests environment variable fallback
  - `TestRandomChoice()` - validates random selection from slices

**Integration Tests:**
- **Scope:** HTTP handlers with template rendering and file serving
- **Approach:** Use `httptest` to simulate full request/response cycle
- **Examples to write:**
  - `TestHandleHTTP()` - template rendering, status codes, headers
  - `TestHandleExchange()` - JSON response structure and content-type
  - `TestHandlePlain()` - plain text response and format
  - `TestImageServer()` - static file serving and cache headers
  - `TestSecurityHeaders()` - middleware sets correct headers

**E2E Tests:**
- **Framework:** Not currently in use
- **Recommendation:** Could use `curl` with shell scripts or Go's `os/exec` package for testing server startup and shutdown

## Common Patterns

**Async Testing:**
- Go testing is synchronous by default
- Goroutines tested indirectly through integration tests
- Use channels and `time.After()` to test timeout scenarios

**Error Testing:**
```go
// Example pattern (not currently in codebase):
func TestHandleHTTPWithBadTemplate(t *testing.T) {
    // Arrange: Mock template parsing failure
    // This would require refactoring to inject template

    // Act & Assert: Ensure proper error response
}
```

**Graceful Shutdown Testing:**
- Test server shutdown with context cancellation
- Verify timeout enforcement during shutdown
- Ensure goroutines are properly cleaned up

## Current Testing Status

**Coverage Gaps:**
- No HTTP handler tests
- No JSON response structure validation
- No IP extraction logic tests
- No template rendering tests
- No middleware security header tests
- No error handling tests
- No concurrent request tests

**Untested Areas:**
- `handleHTTP()` - `main.go` lines 42-61
- `handleExchange()` - `main.go` lines 28-33
- `handlePlain()` - `main.go` lines 35-40
- `clientIP()` - `main.go` lines 63-78
- `scheme()` - `main.go` lines 80-85
- `getEnv()` - `main.go` lines 87-92
- `main()` - `main.go` lines 94-160 (server startup and shutdown)
- `generateExchange()` - `messages.go` lines 191-199
- `generateMessage()` - `messages.go` lines 201-207
- `randomChoice()` - `messages.go` lines 187-189

**Recommendation:** Implement unit tests for utility functions first, then add integration tests for HTTP handlers.

---

*Testing analysis: 2026-02-07*
