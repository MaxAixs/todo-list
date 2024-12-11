package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"time"
	"todo-list/cmd/server"
	"todo-list/pkg/database"
	"todo-list/todo/handler"
	"todo-list/todo/repository"
	"todo-list/todo/service"
)

//	@title			Todo App API
//	@version		1.0
//	@description	API for TodoList Application

//	@host		localhost:8000
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatal("error init config:", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal("error loading env variables:", err)
	}

	db, err := database.NewPostgresDB(database.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatal("Cant run DB", err)
	}
	defer db.Close()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = viper.GetString("port")
		logrus.Printf("SERVER_PORT not set, using default %s", port)
	}

	repo := repository.NewRepository(db)
	s := service.NewService(repo)
	h := handler.NewHandler(s)

	srv := &server.Server{}

	go func() {
		if err := srv.RunServer(port, h.MapRoutes()); err != nil {
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

func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
