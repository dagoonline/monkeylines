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
        @import url('https://fonts.googleapis.com/css2?family=Pirata+One&family=IM+Fell+English:ital@0;1&display=swap');

        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: 'IM Fell English', Georgia, serif;
            background: linear-gradient(180deg,
                #0a0a1a 0%,
                #0d1b2a 20%,
                #1b263b 40%,
                #274c77 70%,
                #1b4332 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 20px;
            position: relative;
            overflow-x: hidden;
        }

        body::before {
            content: '';
            position: fixed;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            background:
                radial-gradient(ellipse at 20% 20%, rgba(255, 255, 150, 0.03) 0%, transparent 50%),
                radial-gradient(ellipse at 80% 30%, rgba(255, 255, 200, 0.02) 0%, transparent 40%),
                radial-gradient(ellipse at 50% 80%, rgba(0, 180, 216, 0.05) 0%, transparent 50%);
            pointer-events: none;
        }

        .stars {
            position: fixed;
            top: 0;
            left: 0;
            right: 0;
            height: 60%;
            background-image:
                radial-gradient(2px 2px at 20px 30px, #fff, transparent),
                radial-gradient(2px 2px at 40px 70px, rgba(255,255,255,0.8), transparent),
                radial-gradient(1px 1px at 90px 40px, #fff, transparent),
                radial-gradient(2px 2px at 130px 80px, rgba(255,255,255,0.6), transparent),
                radial-gradient(1px 1px at 160px 120px, #fff, transparent),
                radial-gradient(2px 2px at 200px 50px, rgba(255,255,255,0.7), transparent),
                radial-gradient(1px 1px at 250px 160px, #fff, transparent),
                radial-gradient(2px 2px at 300px 90px, rgba(255,255,255,0.5), transparent),
                radial-gradient(1px 1px at 350px 30px, #fff, transparent),
                radial-gradient(2px 2px at 400px 140px, rgba(255,255,255,0.8), transparent);
            background-size: 400px 200px;
            animation: twinkle 4s ease-in-out infinite;
            pointer-events: none;
        }

        @keyframes twinkle {
            0%, 100% { opacity: 1; }
            50% { opacity: 0.7; }
        }

        .container {
            max-width: 800px;
            width: 100%;
            position: relative;
            z-index: 1;
        }

        .card {
            background: linear-gradient(145deg,
                #d4a574 0%,
                #c9a067 20%,
                #e6d5b8 50%,
                #c9a067 80%,
                #b8860b 100%);
            border-radius: 8px;
            padding: 50px 40px;
            box-shadow:
                0 0 0 4px #8b4513,
                0 0 0 8px #654321,
                0 20px 60px rgba(0, 0, 0, 0.6),
                inset 0 0 30px rgba(139, 69, 19, 0.2);
            text-align: center;
            position: relative;
        }

        .card::before {
            content: '';
            position: absolute;
            top: 15px;
            left: 15px;
            right: 15px;
            bottom: 15px;
            border: 2px solid rgba(139, 69, 19, 0.3);
            border-radius: 4px;
            pointer-events: none;
        }

        .header {
            margin-bottom: 30px;
        }

        h1 {
            font-family: 'Pirata One', cursive;
            font-size: 3.5em;
            color: #2c1810;
            margin-bottom: 10px;
            text-shadow:
                2px 2px 0px #8b4513,
                4px 4px 8px rgba(0, 0, 0, 0.3);
            letter-spacing: 2px;
        }

        .skull {
            display: inline-block;
            margin: 0 10px;
        }

        .subtitle {
            color: #4a3728;
            font-size: 1.3em;
            font-style: italic;
            text-shadow: 1px 1px 2px rgba(255, 255, 255, 0.3);
        }

        .quote-container {
            position: relative;
            margin: 40px 0;
            padding: 35px 30px;
            background: linear-gradient(135deg,
                #1b4332 0%,
                #2d6a4f 30%,
                #40916c 50%,
                #2d6a4f 70%,
                #1b4332 100%);
            border-radius: 6px;
            border: 3px solid #0d260d;
            box-shadow:
                inset 0 0 20px rgba(0, 0, 0, 0.4),
                0 8px 20px rgba(0, 0, 0, 0.3);
        }

        .quote-mark {
            font-size: 3.5em;
            color: rgba(212, 175, 55, 0.5);
            position: absolute;
            font-family: 'Pirata One', cursive;
        }

        .quote-mark.open {
            top: 5px;
            left: 15px;
        }

        .quote-mark.close {
            bottom: 5px;
            right: 15px;
        }

        .quote {
            font-size: 1.7em;
            color: #f4d03f;
            line-height: 1.6;
            position: relative;
            z-index: 1;
            text-shadow:
                2px 2px 4px rgba(0, 0, 0, 0.8),
                0 0 10px rgba(244, 208, 63, 0.3);
            font-style: italic;
        }

        .actions {
            margin-top: 35px;
            display: flex;
            gap: 20px;
            justify-content: center;
            flex-wrap: wrap;
        }

        .btn {
            padding: 15px 35px;
            font-size: 1.2em;
            font-family: 'Pirata One', cursive;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            transition: all 0.3s ease;
            text-decoration: none;
            display: inline-block;
            letter-spacing: 1px;
        }

        .btn-primary {
            background: linear-gradient(180deg,
                #00b4d8 0%,
                #0096c7 50%,
                #0077b6 100%);
            color: #fff;
            border: 3px solid #023e8a;
            box-shadow:
                0 4px 15px rgba(0, 119, 182, 0.4),
                inset 0 1px 0 rgba(255, 255, 255, 0.3);
            text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.5);
        }

        .btn-primary:hover {
            background: linear-gradient(180deg,
                #48cae4 0%,
                #00b4d8 50%,
                #0096c7 100%);
            transform: translateY(-2px);
            box-shadow:
                0 6px 20px rgba(0, 180, 216, 0.5),
                inset 0 1px 0 rgba(255, 255, 255, 0.4);
        }

        .info {
            margin-top: 35px;
            padding: 20px;
            background: rgba(44, 24, 16, 0.15);
            border-radius: 6px;
            border: 2px solid rgba(139, 69, 19, 0.4);
        }

        .info h3 {
            color: #2c1810;
            margin-bottom: 12px;
            font-family: 'Pirata One', cursive;
            font-size: 1.4em;
            letter-spacing: 1px;
        }

        .info p {
            color: #4a3728;
        }

        .info code {
            background: rgba(27, 67, 50, 0.8);
            padding: 4px 10px;
            border-radius: 3px;
            font-family: 'Courier New', monospace;
            color: #48cae4;
            border: 1px solid #0d260d;
        }

        .footer {
            margin-top: 35px;
            color: #5a4a3a;
            font-size: 1em;
        }

        .footer p {
            margin: 5px 0;
        }

        .mi-quote {
            font-style: italic;
            color: #6b5a4a;
        }

        .decorative-line {
            width: 60%;
            height: 2px;
            background: linear-gradient(90deg,
                transparent 0%,
                #8b4513 20%,
                #d4a574 50%,
                #8b4513 80%,
                transparent 100%);
            margin: 20px auto;
        }

        @media (max-width: 600px) {
            h1 {
                font-size: 2.2em;
            }

            .quote {
                font-size: 1.3em;
            }

            .card {
                padding: 35px 20px;
            }

            .quote-container {
                padding: 25px 20px;
            }
        }
    </style>
</head>
<body>
    <div class="stars"></div>
    <div class="container">
        <div class="card">
            <div class="header">
                <h1><span class="skull">&#9760;</span> MonkeyLines <span class="skull">&#9760;</span></h1>
                <div class="decorative-line"></div>
                <p class="subtitle">Insult Sword Fighting</p>
            </div>

            <div class="quote-container">
                <span class="quote-mark open">"</span>
                <p class="quote">{{.Message}}</p>
                <span class="quote-mark close">"</span>
            </div>

            <div class="actions">
                <a href="/" class="btn btn-primary">&#9876; Draw Another Insult!</a>
            </div>

            <div class="info">
                <h3>&#9875; TCP Access</h3>
                <p>Connect via telnet or netcat:</p>
                <p style="margin-top: 10px;"><code>telnet localhost {{.TCPPort}}</code></p>
                <p style="margin-top: 5px;"><code>nc localhost {{.TCPPort}}</code></p>
            </div>

            <div class="decorative-line"></div>

            <div class="footer">
                <p>Inspired by The Secret of Monkey Island &copy; 1990</p>
                <p class="mi-quote">"You fight like a dairy farmer!"</p>
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
