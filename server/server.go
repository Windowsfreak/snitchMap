package server

import (
	"compress/gzip"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/Windowsfreak/go-mc/bot/repository"
	"github.com/Windowsfreak/go-mc/domain"
	"github.com/Windowsfreak/go-mc/endpoint/world"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Run starts the HTTP server
func Run() {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&domain.Config)
	if err != nil {
		log.Fatal(err)
	}

	if len(domain.Config.PreSharedKey) < 1 {
		log.Fatal("pre-shared key missing in config")
	}

	handler := setUpServer()
	var srv *http.Server
	if domain.Config.Https {
		config := &tls.Config{MinVersion: tls.VersionTLS10}
		srv = &http.Server{Addr: domain.Config.ServerAddr, Handler: handler, TLSConfig: config}
	} else {
		srv = &http.Server{Addr: domain.Config.ServerAddr, Handler: handler}
	}
	go func() {
		log.Println("Starting server, port:", srv.Addr, ", TLS enabled:", domain.Config.Https)

		if domain.Config.Https {
			err = srv.ListenAndServeTLS("/var/lib/dehydrated/certs/windowsfreak.de/fullchain.pem", "/var/lib/dehydrated/certs/windowsfreak.de/privkey.pem")
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil {
			if err == http.ErrServerClosed {
				log.Println("Server shut down. Waiting for connections to drain.")
			} else {
				log.Fatal("failed to start server, port:", srv.Addr, err)
			}
		}
	}()

	// Wait for an interrupt
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)    // interrupt signal sent from terminal
	signal.Notify(sigint, syscall.SIGTERM) // sigterm signal sent from system
	<-sigint

	log.Println("Shutting down server")

	attemptGracefulShutdown(srv)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// https://gist.github.com/the42/1956518
func makeGzipHandler(fn http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		fn.ServeHTTP(gzr, r)
	}
}

func makeDumpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := world.DecodeGetRequest(context.Background(), r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, err.Error())
		} else {
			http.ServeFile(w, r, "./snitchLogs.db")
		}
	}
}

func setUpServer() http.Handler {
	mux := http.NewServeMux()

	r, err := repository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}
	worldService := world.NewService(r)
	mux.Handle("/world/events/after", world.MakeGetEventsAfterHandler(worldService))
	mux.Handle("/world/events/user/after", world.MakeGetEventsByUserAfterHandler(worldService))
	mux.Handle("/world/events/region/after", world.MakeGetEventsByRegionAfterHandler(worldService))
	mux.Handle("/world/events/id", world.MakeGetLastEventRowIdBeforeHandler(worldService))
	mux.Handle("/world/events/user/id", world.MakeGetLastEventRowIdByUserBeforeHandler(worldService))
	mux.Handle("/world/events/region/id", world.MakeGetLastEventRowIdByRegionBeforeHandler(worldService))
	mux.Handle("/world/chats/after", world.MakeGetChatsAfterHandler(worldService))
	mux.Handle("/world/chats/id", world.MakeGetLastChatRowIdBeforeHandler(worldService))
	mux.Handle("/world/users/after", world.MakeGetUsersSeenAfterHandler(worldService))
	mux.Handle("/world/user", world.MakeGetUserHandler(worldService))
	mux.Handle("/world/snitches/all", world.MakeGetSnitchesHandler(worldService))
	mux.Handle("/world/alerts/set", world.MakeSetSnitchAlertByRegionHandler(worldService))
	mux.Handle("/world/everything", makeDumpHandler())
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)

	return accessControl(makeGzipHandler(mux))
}

func attemptGracefulShutdown(srv *http.Server) {
	if err := shutdownServer(srv, 25*time.Second); err != nil {
		log.Println("failed to shutdown server", err)
	}
}

func shutdownServer(srv *http.Server, maximumTime time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), maximumTime)
	defer cancel()
	return srv.Shutdown(ctx)
}
