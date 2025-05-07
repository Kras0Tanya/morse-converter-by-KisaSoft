package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Kras0Tanya/morse-converter-by-KisaSoft/internal/server"
)

// это даже в чёт-то даже приятная часть работы, но я всё ещё учусь обработке ошибок

func main() {

	logger := log.New(os.Stderr, "SERVER: ", log.LstdFlags)

	myServer := server.CreateServer(logger)

	logger.Printf("Starting server on %s\n", myServer.Server.Addr)

	err := myServer.Server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Server error: %v", err)
	}
}
