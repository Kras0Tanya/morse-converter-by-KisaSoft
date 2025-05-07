package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Kras0Tanya/morse-converter-by-KisaSoft/internal/handlers"
)

// эта часть задания вместе с main показалась уже не такой ужасной

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func CreateServer(logger *log.Logger) *Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: srv,
	}
}
