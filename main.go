package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"sync"
)

var insults = []string{
	"You code like a %s!",
	"Your commits smell like %s!",
	"You fight bugs like a %s!",
	"I've seen %s write better code!",
	"Your code is as ugly as a %s!",
	"You debug slower than a %s!",
	"My %s could merge better than you!",
	"Your functions are as useless as a %s!",
	"You refactor like a %s!",
	"Soon your code will be as dead as a %s!",
	"Your variables smell like a %s!",
	"I've taught %s to code better!",
	"Every bug you fix spawns %s!",
	"Your pull requests look like %s!",
	"You handle exceptions like a %s!",
}

var comebacks = []string{
	"How appropriate. You %s like a %s.",
	"First you'd better stop %s like a %s.",
	"I'm glad you attended your %s reunion.",
	"And I thought you smelled like a %s.",
	"At least I know how to %s a %s.",
	"You make me want to %s my %s.",
	"I'd %s, but I don't want to dirty my %s.",
	"Then you better stop %s your %s.",
	"Even so, my %s can still %s.",
	"Too bad nobody will %s your %s.",
	"I've %s worse %s than you.",
	"Yet you still can't %s a simple %s.",
	"Funny, your %s said the same about your %s.",
	"That explains the %s in your repository.",
	"I'll %s that into your %s.",
}

var insultNouns = []string{
	"dairy farmer", "rubber duck", "segfault", "null pointer",
	"deprecated function", "memory leak", "infinite loop", "stack overflow",
	"merge conflict", "legacy codebase", "untested module", "spaghetti monster",
	"floating point", "race condition", "dead code", "code monkey",
	"keyboard warrior", "copy-paster", "tab user", "vim user",
}

var comebackNouns = []string{
	"cow", "compiler", "debugger", "garbage collector",
	"exception handler", "code reviewer", "unit test", "documentation",
	"git history", "production server", "staging environment", "bug tracker",
	"coffee machine", "rubber duck", "stack trace", "error log",
	"keyboard", "monitor", "semicolon", "curly brace",
}

var verbs = []string{
	"debug", "compile", "refactor", "deploy", "merge", "commit",
	"push", "pull", "branch", "rebase", "squash", "cherry-pick",
	"rollback", "hotfix", "optimize", "minify", "lint", "test",
}

var gerunds = []string{
	"debugging", "compiling", "refactoring", "deploying", "merging", "committing",
	"pushing", "pulling", "branching", "rebasing", "squashing", "cherry-picking",
	"rolling back", "hotfixing", "optimizing", "minifying", "linting", "testing",
}

var pastVerbs = []string{
	"debugged", "compiled", "refactored", "deployed", "merged", "committed",
	"pushed", "pulled", "branched", "rebased", "squashed", "cherry-picked",
	"rolled back", "hotfixed", "optimized", "minified", "linted", "tested",
}

func randomChoice[T any](slice []T) T {
	return slice[rand.Intn(len(slice))]
}

func generateMessage() string {
	if rand.Intn(2) == 0 {
		template := randomChoice(insults)
		return fmt.Sprintf(template, randomChoice(insultNouns))
	}

	template := randomChoice(comebacks)
	noun := randomChoice(comebackNouns)
	verb := randomChoice(verbs)
	gerund := randomChoice(gerunds)
	pastVerb := randomChoice(pastVerbs)

	// Handle different template patterns
	switch {
	case strings.Contains(template, "%s reunion"):
		return fmt.Sprintf(template, noun)
	case strings.Contains(template, "smelled like"):
		return fmt.Sprintf(template, noun)
	case strings.Contains(template, "stop %s like"):
		return fmt.Sprintf(template, gerund, noun)
	case strings.Contains(template, "stop %s your"):
		return fmt.Sprintf(template, gerund, noun)
	case strings.Contains(template, "can still"):
		return fmt.Sprintf(template, noun, verb)
	case strings.Contains(template, "I've %s worse"):
		return fmt.Sprintf(template, pastVerb, noun)
	case strings.Contains(template, "said the same"):
		return fmt.Sprintf(template, noun, gerund)
	case strings.Contains(template, "in your repository"):
		return fmt.Sprintf(template, noun)
	default:
		return fmt.Sprintf(template, verb, noun)
	}
}

var htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>MonkeyLines - Sword Fighting Insults</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: 'Georgia', serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 20px;
        }

        .container {
            max-width: 800px;
            width: 100%;
        }

        .card {
            background: rgba(255, 255, 255, 0.95);
            border-radius: 20px;
            padding: 60px 40px;
            box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
            text-align: center;
            backdrop-filter: blur(10px);
        }

        .header {
            margin-bottom: 40px;
        }

        h1 {
            font-size: 3em;
            color: #2d3748;
            margin-bottom: 10px;
            text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
        }

        .subtitle {
            color: #718096;
            font-size: 1.2em;
            font-style: italic;
        }

        .quote-container {
            position: relative;
            margin: 50px 0;
            padding: 40px;
            background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
            border-radius: 15px;
            box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
        }

        .quote-mark {
            font-size: 4em;
            color: rgba(255, 255, 255, 0.3);
            position: absolute;
            font-family: 'Times New Roman', serif;
        }

        .quote-mark.open {
            top: 10px;
            left: 20px;
        }

        .quote-mark.close {
            bottom: 10px;
            right: 20px;
        }

        .quote {
            font-size: 1.8em;
            color: white;
            line-height: 1.6;
            position: relative;
            z-index: 1;
            text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);
        }

        .actions {
            margin-top: 40px;
            display: flex;
            gap: 20px;
            justify-content: center;
            flex-wrap: wrap;
        }

        .btn {
            padding: 15px 30px;
            font-size: 1.1em;
            border: none;
            border-radius: 50px;
            cursor: pointer;
            transition: all 0.3s ease;
            font-weight: bold;
            text-decoration: none;
            display: inline-block;
        }

        .btn-primary {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            box-shadow: 0 5px 15px rgba(102, 126, 234, 0.4);
        }

        .btn-primary:hover {
            transform: translateY(-2px);
            box-shadow: 0 7px 20px rgba(102, 126, 234, 0.6);
        }

        .info {
            margin-top: 40px;
            padding: 20px;
            background: rgba(237, 242, 247, 0.5);
            border-radius: 10px;
            border-left: 4px solid #667eea;
        }

        .info h3 {
            color: #2d3748;
            margin-bottom: 10px;
        }

        .info code {
            background: rgba(102, 126, 234, 0.1);
            padding: 2px 8px;
            border-radius: 4px;
            font-family: 'Courier New', monospace;
            color: #667eea;
        }

        .footer {
            margin-top: 40px;
            color: #718096;
            font-size: 0.9em;
        }

        @media (max-width: 600px) {
            h1 {
                font-size: 2em;
            }

            .quote {
                font-size: 1.3em;
            }

            .card {
                padding: 40px 20px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="card">
            <div class="header">
                <h1>⚔️ MonkeyLines</h1>
                <p class="subtitle">Sword Fighting Insults Generator</p>
            </div>

            <div class="quote-container">
                <span class="quote-mark open">"</span>
                <p class="quote">{{.Message}}</p>
                <span class="quote-mark close">"</span>
            </div>

            <div class="actions">
                <a href="/" class="btn btn-primary">⚔️ Another Insult!</a>
            </div>

            <div class="info">
                <h3>🐵 TCP Access</h3>
                <p>Connect via telnet or netcat to get insults via TCP:</p>
                <p style="margin-top: 10px;"><code>telnet localhost {{.TCPPort}}</code></p>
                <p style="margin-top: 5px;"><code>nc localhost {{.TCPPort}}</code></p>
            </div>

            <div class="footer">
                <p>Inspired by The Secret of Monkey Island</p>
                <p style="margin-top: 5px;">🏴‍☠️ "You fight like a dairy farmer!"</p>
            </div>
        </div>
    </div>
</body>
</html>
`

type PageData struct {
	Message string
	TCPPort string
}

var (
	tmpl     *template.Template
	mu       sync.Mutex
	lastLine string
)

func init() {
	var err error
	tmpl, err = template.New("index").Parse(htmlTemplate)
	if err != nil {
		log.Fatal("Failed to parse template:", err)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()
	log.Printf("TCP connection from %s", clientAddr)

	message := generateMessage()

	mu.Lock()
	lastLine = message
	mu.Unlock()

	writer := bufio.NewWriter(conn)
	fmt.Fprintf(writer, "%s\n", message)
	writer.Flush()

	log.Printf("Sent to %s: %s", clientAddr, message)
}

func startTCPServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Failed to start TCP server:", err)
	}
	defer listener.Close()

	log.Printf("TCP server listening on port %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleTCPConnection(conn)
	}
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	message := generateMessage()

	mu.Lock()
	lastLine = message
	mu.Unlock()

	data := PageData{
		Message: message,
		TCPPort: "8023",
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	log.Printf("HTTP %s %s from %s - Served: %s", r.Method, r.URL.Path, r.RemoteAddr, message)
}

func startHTTPServer(port string) {
	http.HandleFunc("/", handleHTTP)

	log.Printf("HTTP server listening on port %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}
}

func main() {
	tcpPort := "8023"
	httpPort := "8080"

	log.Println("🐵 MonkeyLines Server Starting...")
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// Start TCP server in a goroutine
	go startTCPServer(tcpPort)

	// Start HTTP server in main goroutine
	log.Printf("🌐 HTTP: http://localhost:%s", httpPort)
	log.Printf("🔌 TCP:  telnet localhost %s", tcpPort)
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	startHTTPServer(httpPort)
}
