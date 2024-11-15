package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"
	"todo-list/todo"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
		logrus.Printf("SERVER_PORT not set, using default %s", port)
	}

	srv := new(todo.)
	handler := &todo.Handler{}

	go func() {
		if err := srv.RunServer(port, handler.MapRoutes()); err != nil {
			logrus.Fatalf("cant run server: %s", err)
		}
	}()

	logrus.Println("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.ShutDown(ctx); err != nil {
		logrus.Fatal("cant shutdown server: %s", err)
	}

	logrus.Println("server shutdown")
}
