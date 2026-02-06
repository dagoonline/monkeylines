package main

import (
	"bufio"
	"bytes"
	"context"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:embed index.html
var indexHTML string

//go:embed images
var imagesFS embed.FS

var tmpl *template.Template

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()
	log.Printf("TCP connection from %s", clientAddr)

	message := generateMessage()

	writer := bufio.NewWriter(conn)
	fmt.Fprintf(writer, "%s\n", message)
	writer.Flush()

	log.Printf("Sent to %s: %s", clientAddr, message)
}

func startTCPServer(ctx context.Context, port string) {
	lc := net.ListenConfig{}
	listener, err := lc.Listen(ctx, "tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to start TCP server: %v", err)
	}
	defer listener.Close()

	log.Printf("TCP server listening on port %s", port)

	go func() {
		<-ctx.Done()
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleTCPConnection(conn)
	}
}

func handlePlain(w http.ResponseWriter, r *http.Request) {
	message := generateMessage()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, message)
	log.Printf("HTTP %s %s from %s - Served: %s", r.Method, r.URL.Path, r.RemoteAddr, message)
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	message := generateMessage()

	data := struct{ Message string }{Message: message}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)

	log.Printf("HTTP %s %s from %s - Served: %s", r.Method, r.URL.Path, r.RemoteAddr, message)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	var err error
	tmpl, err = template.New("index").Parse(indexHTML)
	if err != nil {
		log.Fatal("Failed to parse template:", err)
	}

	tcpPort := getEnv("MONKEYLINES_TCP_PORT", "8023")
	httpPort := getEnv("MONKEYLINES_HTTP_PORT", "8080")

	log.Println("MonkeyLines Server Starting...")
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go startTCPServer(ctx, tcpPort)

	imageServer := http.FileServer(http.FS(imagesFS))

	mux := http.NewServeMux()
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=604800")
		imageServer.ServeHTTP(w, r)
	})
	mux.HandleFunc("/line", handlePlain)
	mux.HandleFunc("/", handleHTTP)

	httpServer := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	log.Printf("HTTP: http://localhost:%s", httpPort)
	log.Printf("TCP:  telnet localhost %s", tcpPort)
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down servers...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	log.Println("Server stopped.")
}
