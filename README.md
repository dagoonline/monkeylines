# MonkeyLines

An HTTP server that delivers Monkey Island sword fighting-style insults with an animated three-headed monkey.

> "You code like a dairy farmer!"
> "How appropriate. You debug like a cow."

## What is this?

MonkeyLines is a fun service that generates procedurally created insults and comebacks inspired by *The Secret of Monkey Island*'s iconic sword fighting mechanic. It features an animated pixel-art three-headed monkey web interface and a plain text API endpoint.

## Features

- **Web Interface** - Animated pixel-art three-headed monkey with typewriter speech bubble effect
- **Plain Text API** - `GET /line` returns just the text, perfect for scripting
- **Message Generator** - Thousands of unique combinations of insults and comebacks
- **Single Binary** - All assets (HTML, images) embedded at compile time
- **Security Hardened** - CSP headers, server timeouts, safe template rendering

## Installation

```bash
# Clone the repository
git clone https://github.com/dagoonline/monkeylines.git
cd monkeylines

# Build
go build -o monkeylines .

# Run
./monkeylines
```

## Usage

### Start the Server

```bash
./monkeylines
```

The server starts on port `8080` with two endpoints:
- `/` — Web interface
- `/line` — Plain text API

### Web Interface

Open your browser and navigate to:
```
http://localhost:8080
```

Features an animated three-headed monkey from Monkey Island that "talks" the insult via a typewriter effect in a speech bubble. Click "Draw your sword again..." to get a new insult.

### Plain Text API

```bash
curl http://localhost:8080/line
```

Returns a single line of text — useful for scripts, bots, or piping into other commands.

## Configuration

The port is configurable via environment variable:

| Variable | Default | Description |
|---|---|---|
| `MONKEYLINES_HTTP_PORT` | `8080` | HTTP server port |

```bash
MONKEYLINES_HTTP_PORT=3000 ./monkeylines
```

## Message Generator

The generator combines:
- **15 insult templates** - "You code like a {noun}!"
- **15 comeback templates** - "How appropriate. You {verb} like a {noun}."
- **20 insult nouns** - dairy farmer, null pointer, spaghetti monster...
- **20 comeback nouns** - cow, compiler, garbage collector...
- **18 verbs** - debug, compile, refactor, deploy...

This creates thousands of unique message combinations.

## Architecture

- **Embedded assets** - HTML template and images compiled into the binary via `go:embed`
- **Single binary** - No external dependencies or runtime files needed
- **Image caching** - Static assets served with one-week cache headers
- **Security headers** - CSP, X-Frame-Options, X-Content-Type-Options, Referrer-Policy

## Requirements

- Go 1.21+ (for building)
- No external dependencies - uses only Go standard library

## Motivation

This project was built as a hands-on test of [Claude Opus 4.6](https://www.anthropic.com/claude) — the entire codebase was written collaboratively with it using [Claude Code](https://claude.com/claude-code).

## License

Do whatever you want with it. You fight like a cow.

## Inspired By

*The Secret of Monkey Island* (1990) - LucasArts
"How appropriate. You fight like a cow."
