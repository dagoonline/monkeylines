package main

import (
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
	"strings"
	"syscall"
	"time"
)

//go:embed index.html
var indexHTML string

//go:embed images
var imagesFS embed.FS

var tmpl *template.Template

func handleInsult(w http.ResponseWriter, r *http.Request) {
	message := generateInsult()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, message)
	log.Printf("HTTP %s %s://%s%s from %s - Served insult: %s", r.Method, scheme(r), r.Host, r.URL.Path, clientIP(r), message)
}

func handlePlain(w http.ResponseWriter, r *http.Request) {
	message := generateMessage()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, message)
	log.Printf("HTTP %s %s://%s%s from %s - Served: %s", r.Method, scheme(r), r.Host, r.URL.Path, clientIP(r), message)
}

func handleComeback(w http.ResponseWriter, r *http.Request) {
	message := generateComeback()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, message)
	log.Printf("HTTP %s %s://%s%s from %s - Served comeback: %s", r.Method, scheme(r), r.Host, r.URL.Path, clientIP(r), message)
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	insult := generateInsult()
	monkey := randomSideMonkey()

	data := struct {
		Insult string
		Monkey string
	}{Insult: insult, Monkey: monkey}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)

	log.Printf("HTTP %s %s://%s%s from %s - Served insult (%s): %s", r.Method, scheme(r), r.Host, r.URL.Path, clientIP(r), monkey, insult)
}

func clientIP(r *http.Request) string {
	if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
		// X-Forwarded-For may contain a comma-separated list; first entry is the client
		if ip := strings.TrimSpace(strings.SplitN(fwd, ",", 2)[0]); ip != "" {
			return ip
		}
	}
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return strings.TrimSpace(ip)
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

func scheme(r *http.Request) string {
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		return "https"
	}
	return "http"
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

	httpPort := getEnv("MONKEYLINES_HTTP_PORT", "8080")

	log.Println("MonkeyLines Server Starting...")
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	securityHeaders := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("X-Frame-Options", "DENY")
			w.Header().Set("Referrer-Policy", "no-referrer")
			w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'unsafe-inline'; style-src 'unsafe-inline'; img-src 'self'")
			next.ServeHTTP(w, r)
		})
	}

	imageServer := http.FileServer(http.FS(imagesFS))

	mux := http.NewServeMux()
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=604800")
		imageServer.ServeHTTP(w, r)
	})
	mux.HandleFunc("/line", handlePlain)
	mux.HandleFunc("/insult", handleInsult)
	mux.HandleFunc("/comeback", handleComeback)
	mux.HandleFunc("/", handleHTTP)

	httpServer := &http.Server{
		Addr:              ":" + httpPort,
		Handler:           securityHeaders(mux),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	log.Printf("HTTP: port %s", httpPort)
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	log.Println("Server stopped.")
}
