# MonkeyLines

An HTTP server that delivers Monkey Island sword fighting-style insults with an animated three-headed monkey.

> "You code like a dairy farmer!"
> "How appropriate. You fight like a cow."

## What is this?

MonkeyLines is a fun service that generates procedurally created insults and comebacks inspired by *The Secret of Monkey Island*'s iconic sword fighting mechanic. It features an animated pixel-art three-headed monkey web interface and a plain text API endpoint.

## Features

- **Web Interface** - Animated pixel-art three-headed monkey with typewriter speech bubble effect
- **Insult Sword Fighting** - Side monkeys throw insults, center monkey delivers themed comebacks
- **Auto-play Mode** - Toggle continuous exchanges with the 🙊 button
- **Plain Text API** - `GET /line` returns just the text, perfect for scripting
- **Exchange API** - `GET /exchange` returns a paired insult/comeback as JSON
- **Themed Generator** - 8 themes with thousands of unique insult/comeback combinations
- **Single Binary** - All assets (HTML, images) embedded at compile time
- **Security Hardened** - CSP headers, server timeouts, safe template rendering
- **CI/CD** - GitHub Actions workflow for automatic deployment to Fly.io

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

The server starts on port `8080` with the following endpoints:
- `/` — Web interface
- `/line` — Plain text API (random insult or comeback)
- `/exchange` — JSON API (paired insult + comeback)

### Web Interface

Open your browser and navigate to:
```
http://localhost:8080
```

Features an animated three-headed monkey from Monkey Island. Press ⚔️ to trigger an insult sword fight — a side monkey delivers the insult, then the center monkey fires back with a themed comeback. Press 🙊 to enable auto-play mode for continuous exchanges.

### Plain Text API

```bash
curl http://localhost:8080/line
```

Returns a single line of text — useful for scripts, bots, or piping into other commands.

### Exchange API

```bash
curl http://localhost:8080/exchange
```

Returns a JSON object with a paired insult and comeback:

```json
{"Insult":"You code like a dairy farmer!","Comeback":"How appropriate. You fight like a cow."}
```

## Configuration

The port is configurable via environment variable:

| Variable | Default | Description |
|---|---|---|
| `MONKEYLINES_HTTP_PORT` | `8080` | HTTP server port |

```bash
MONKEYLINES_HTTP_PORT=3000 ./monkeylines
```

## Message Generator

The generator uses a themed template system. Each theme groups related insult templates, comeback templates, and a shared word list:

- **8 themes** - Classic Monkey Island, code quality, debugging, deployments, git history, error handling, architecture, testing
- **4 insult templates per theme** - "You code like a %s!"
- **4 comeback templates per theme** - "How appropriate. You fight like a %s."
- **7-10 nouns per theme** - dairy farmer, merge conflict, infinite loop...

Insults and comebacks are always drawn from the same theme, keeping pairs coherent. The combinatorics produce thousands of unique exchanges.

## Deployment

### Fly.io

The project includes a `fly.toml` and a GitHub Actions workflow that deploys on push to `main`/`master`. Set the `FLY_API_TOKEN` secret in your GitHub repository settings.

### Docker

```bash
docker build -t monkeylines .
docker run -p 8080:8080 monkeylines
```

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
