package main

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunServer(port string) error {
	s.httpServer = &http.Server{
		Addr:              ":" + port,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func main() {
	server := &Server{}
	if err := server.RunServer("8080"); err != nil {
		log.Fatalf("cant run server: %s", err)
	}
}
