# MonkeyLines

A dual-protocol server that delivers Monkey Island sword fighting-style insults via TCP, HTTP, and a plain text API.

> "You code like a dairy farmer!"
> "How appropriate. You debug like a cow."

## What is this?

MonkeyLines is a fun service that generates procedurally created insults and comebacks inspired by *The Secret of Monkey Island*'s iconic sword fighting mechanic. It features an animated three-headed monkey web interface, a TCP server for terminal access, and a plain text API endpoint.

## Features

- **HTTP Server** - Animated pixel-art three-headed monkey with typewriter speech bubble effect
- **TCP Server** - Connect via telnet/netcat and receive instant insults
- **Plain Text API** - `GET /line` returns just the text, perfect for scripting
- **Message Generator** - Thousands of unique combinations of insults and comebacks
- **Concurrent** - Handles multiple connections simultaneously
- **Single Binary** - All assets (HTML, images) embedded at compile time

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

The server starts three endpoints:
- **HTTP server** on port `8080` (web interface)
- **Plain text API** on port `8080` at `/line`
- **TCP server** on port `8023`

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

### TCP/Telnet

```bash
# Using telnet
telnet localhost 8023

# Using netcat
nc localhost 8023

# Using echo/pipeline
echo "" | nc localhost 8023
```

Each connection receives a fresh insult and automatically disconnects.

## Configuration

Ports are configurable via environment variables:

| Variable | Default | Description |
|---|---|---|
| `MONKEYLINES_TCP_PORT` | `8023` | TCP server port |
| `MONKEYLINES_HTTP_PORT` | `8080` | HTTP server port (web + API) |

```bash
MONKEYLINES_HTTP_PORT=3000 MONKEYLINES_TCP_PORT=2323 ./monkeylines
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

- **Concurrent design** - TCP connections handled in goroutines
- **Embedded assets** - HTML template and images compiled into the binary via `go:embed`
- **Single binary** - No external dependencies or runtime files needed
- **Image caching** - Static assets served with one-week cache headers

## Requirements

- Go 1.21+ (for building)
- No external dependencies - uses only Go standard library

## License

Do whatever you want with it. You fight like a cow.

## Inspired By

*The Secret of Monkey Island* (1990) - LucasArts
"How appropriate. You fight like a cow."
