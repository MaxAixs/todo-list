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

//	@host		localhost:8080
//	@BasePath  /

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %v", err)
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
		logrus.Fatalf("Cant run DB: %v", err)
	}
	defer db.Close()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = viper.GetString("server.port")
		logrus.Infof("SERVER_PORT not set, using default %s", port)
	}

	repo := repository.NewRepository(db)
	s := service.NewService(repo)
	h := handler.NewHandler(s)

	srv := &server.Server{}
	serverConfig := server.SrvConfig{
		Port:              port,
		ReadHeaderTimeout: viper.GetDuration("server.read_header_timeout"),
		WriteTimeout:      viper.GetDuration("server.write_timeout"),
		IdleTimeout:       viper.GetDuration("server.idle_timeout"),
	}

	go func() {
		if err := srv.RunServer(serverConfig, h.MapRoutes()); err != nil {
			logrus.Fatalf("cant run server: %v", err)
		}
	}()

	logrus.Println("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.ShutDown(ctx); err != nil {
		logrus.Fatalf("cant shutdown server: %v", err)
	}

	logrus.Println("server shutdown")
}

func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
