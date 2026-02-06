# MonkeyLines

A dual-protocol server that delivers Monkey Island sword fighting-style insults via both TCP and HTTP.

> "You code like a dairy farmer!"
> "How appropriate. You debug like a cow."

## What is this?

MonkeyLines is a fun service that generates procedurally created insults and comebacks inspired by *The Secret of Monkey Island*'s iconic sword fighting mechanic. It runs as both a TCP server (for terminal/telnet access) and an HTTP server (for web browser access).

## Features

- **TCP Server** - Connect via telnet/netcat and receive instant insults
- **HTTP Server** - Beautiful web interface with gradient styling
- **Message Generator** - Thousands of unique combinations of insults and comebacks
- **Concurrent** - Handles multiple connections simultaneously
- **Logging** - Tracks all connections and served messages

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

The server will start both:
- **HTTP server** on port 8080
- **TCP server** on port 8023

### Access via Web Browser

Open your browser and navigate to:
```
http://localhost:8080
```

Each refresh generates a new insult. The page features a beautiful gradient design with the insult prominently displayed.

### Access via TCP/Telnet

```bash
# Using telnet
telnet localhost 8023

# Using netcat
nc localhost 8023

# Using echo/pipeline
echo "" | nc localhost 8023
```

Each connection receives a fresh insult and automatically disconnects.

## Examples

### TCP Session
```
$ telnet localhost 8023
Trying 127.0.0.1...
Connected to localhost.
Your functions are as useless as a merge conflict!
Connection closed by foreign host.
```

### Web Interface
Access `http://localhost:8080` to see a beautifully styled page with:
- Gradient purple background
- Centered card with the insult
- Button to generate new insults
- Instructions for TCP access

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
- **Thread-safe** - Protected shared state with mutexes
- **Single binary** - No external dependencies
- **Embedded template** - HTML template compiled into binary

## Configuration

Default ports can be modified in `main.go`:
- `tcpPort := "8023"` - TCP server port
- `httpPort := "8080"` - HTTP server port

## Requirements

- Go 1.21+ (for building)
- No external dependencies - uses only Go standard library

## Use Cases

- Fun addition to your local dev environment
- Testing TCP/HTTP client implementations
- Learning concurrent server design in Go
- Adding personality to monitoring dashboards
- Party trick at developer meetups

## License

Do whatever you want with it. You fight like a cow.

## Inspired By

*The Secret of Monkey Island* (1990) - LucasArts
"How appropriate. You fight like a cow."
