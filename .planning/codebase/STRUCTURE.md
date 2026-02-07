# Codebase Structure

**Analysis Date:** 2026-02-07

## Directory Layout

```
monkeylines/
├── .github/
│   └── workflows/
│       └── deploy.yml              # CI/CD deployment to Fly.io
├── .planning/
│   └── codebase/                   # Codebase analysis documents
├── images/                         # Embedded pixel-art assets
│   ├── base.png                    # Three-headed monkey base scene
│   ├── l.png                       # Left monkey overlay
│   ├── c.png                       # Center monkey overlay
│   ├── r.png                       # Right monkey overlay
│   ├── preview.jpg                 # OG preview image for sharing
│   └── favicon.ico                 # Browser tab icon
├── main.go                         # HTTP server, handlers, middleware
├── messages.go                     # Message generation, themes, templates
├── index.html                      # Embedded HTML template + CSS + JS
├── go.mod                          # Go module definition
├── go.sum                          # (not present) - no external dependencies
├── Dockerfile                      # Multi-stage build for containerization
├── fly.toml                        # Fly.io deployment configuration
├── .gitignore                      # Git ignore rules
└── README.md                       # Project documentation
```

## Directory Purposes

**`.github/workflows/`:**
- Purpose: GitHub Actions CI/CD pipeline
- Contains: Deployment automation
- Key files: `deploy.yml` - Triggers Fly.io deployment on push to main/master

**`images/`:**
- Purpose: Pixel-art assets for web interface
- Contains: PNG overlays for three-headed monkey animation, favicon, preview image
- Generated: No, all hand-crafted pixel art
- Committed: Yes, all assets versioned in git
- Embedded: Yes, bundled into binary via `go:embed`

**`.planning/codebase/`:**
- Purpose: Codebase analysis and planning documents
- Contains: Architecture, structure, conventions, testing, tech stack analysis
- Generated: Yes, created by GSD mapping tools
- Committed: Yes, for team reference

## Key File Locations

**Entry Points:**
- `main.go`: Server startup (main function, HTTP setup, handler registration)

**Configuration:**
- `fly.toml`: Fly.io deployment settings (app name, region, HTTP port)
- `.github/workflows/deploy.yml`: CI/CD configuration
- Environment variables read at `main.go` line 101 (`MONKEYLINES_HTTP_PORT`)

**Core Logic:**
- `main.go`: HTTP server, request routing, security middleware, logging
- `messages.go`: Content generation engine, theme definitions, template substitution

**Frontend:**
- `index.html`: Single-page application with embedded CSS and JavaScript
  - CSS: Lines 20-142
  - JavaScript: Lines 161-378 (IIFE with no external dependencies)
  - HTML structure: Lines 144-159

**Assets:**
- `images/base.png`: 1x1 aspect ratio pixel art (main scene)
- `images/l.png`, `c.png`, `r.png`: Overlay sprites for monkey expressions
- `images/favicon.ico`: 16x16 icon
- `images/preview.jpg`: Social media preview card (512x512)

**Build & Deployment:**
- `Dockerfile`: Multi-stage build (Go 1.21 alpine builder → alpine runtime)
- `go.mod`: Module declaration (no dependencies listed)
- `README.md`: Comprehensive documentation

## Naming Conventions

**Files:**
- `main.go`: Single entry point, contains HTTP layer
- `messages.go`: Message generation and content definitions
- `index.html`: Main HTML template (Go embed reference uses filename)
- `[l|c|r].png`: Named by position in three-headed monkey (left, center, right)

**Directories:**
- Lowercase with no underscores: `.github`, `images`, `workflows`
- Follows Go project conventions (no vendor directory, no src/ wrapper)

**Functions:**
- camelCase: `handleExchange`, `handlePlain`, `handleHTTP`, `generateExchange`, `generateMessage`, `randomChoice`, `clientIP`, `scheme`, `getEnv`
- No underscores in function names
- Handler functions prefixed with `handle`
- Generator functions prefixed with `generate`
- Helper functions use descriptive action verbs

**Variables:**
- camelCase for all variables: `tmpl`, `indexHTML`, `imagesFS`, `ex`, `message`, `themes`, `sideMonkeys`, `audioCtx`, `monkeyVoices`
- Struct fields PascalCase (exported): `Insult`, `Comeback`
- Private fields lowercase: `insults`, `comebacks`, `nouns`, `low`, `high`

**Types:**
- PascalCase structs: `exchange` (private), `theme` (private)
- Constants/package-level: `indexHTML`, `imagesFS` (implicit exports)
- Embedded filesystem: `imagesFS`

## Where to Add New Code

**New Feature (e.g., additional endpoints):**
- HTTP handler: Add new function in `main.go` following pattern of `handleExchange`, `handlePlain`, `handleHTTP` (lines 28-61)
- Route registration: Register in `main.go` ServeMux setup (lines 126-128)
- Tests: Add `*_test.go` file in root directory (tests co-located with source)

**New Message Theme:**
- Theme definition: Add to `themes` slice in `messages.go` (lines 22-185)
- Structure: Follow existing theme pattern with `insults`, `comebacks`, `nouns` string slices
- Example: 4 insult templates, 4 comeback templates, 7-10 nouns per theme

**New Utility Function:**
- Location: Add to `main.go` if related to HTTP (scheme detection, IP parsing), or `messages.go` if related to message generation
- Pattern: Follow existing helper functions (`clientIP`, `scheme`, `getEnv`, `randomChoice`)

**Frontend Enhancement:**
- Location: Modify `index.html` (lines 1-380)
- CSS updates: Lines 20-142 (inline `<style>` tag)
- JavaScript updates: Lines 161-378 (IIFE, no external imports)
- HTML structure: Lines 144-159

**Assets:**
- Location: Add image files to `images/` directory
- Embed: Image will be automatically included via `//go:embed images` at `main.go` line 23
- Caching: Update cache headers in FileServer route if needed (`main.go` line 123)

**Configuration:**
- Runtime port: `MONKEYLINES_HTTP_PORT` environment variable (default 8080)
- Deployment: Update `fly.toml` for Fly.io settings or `.github/workflows/deploy.yml` for CI/CD
- Docker: Update `Dockerfile` if base image or build steps change

## Special Directories

**`images/`:**
- Purpose: Embedded static assets
- Generated: No (all hand-crafted)
- Committed: Yes
- Embedded: Yes, included in binary via `go:embed images` directive
- Cache behavior: Served with 1-week Cache-Control header

**`.github/workflows/`:**
- Purpose: GitHub Actions automation
- Generated: No
- Committed: Yes
- Special handling: Requires `FLY_API_TOKEN` secret in GitHub repository settings

**`.planning/codebase/`:**
- Purpose: GSD mapping output
- Generated: Yes
- Committed: Yes, for team reference
- Contents: ARCHITECTURE.md, STRUCTURE.md, CONVENTIONS.md, TESTING.md, STACK.md, INTEGRATIONS.md, CONCERNS.md

## File Relationship Map

```
main.go
├── imports: "embed", "html/template", "net/http", "fmt", "log", "context", "os", "net", "strings"
├── uses: messages.go (generateExchange, generateMessage)
├── embeds: index.html (template) and images/* (filesystem)
└── exports: main() entry point

messages.go
├── imports: "fmt", "math/rand"
├── defines: exchange struct, theme struct, themes array
└── exports: generateExchange(), generateMessage(), randomChoice()

index.html
├── embedded in: main.go (as indexHTML string)
├── references: /images/base.png, /images/l.png, /images/c.png, /images/r.png, /images/favicon.ico
├── fetches from: /exchange API, /images/* URLs
└── contains: CSS, HTML, JavaScript (no external dependencies)

Dockerfile
├── builds: main.go + messages.go + index.html + images/*
├── produces: monkeylines binary
└── runs: ./monkeylines on port 8080

fly.toml
├── deploys: Docker image (Dockerfile)
└── configures: Fly.io app settings (app name, region, internal port)

.github/workflows/deploy.yml
├── triggers: on push to main or master
├── builds: Dockerfile
└── deploys: via Fly.io API
```

---

*Structure analysis: 2026-02-07*
