package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"time"
	"todo-list/cmd/server"
	"todo-list/internal/config"
	analyticsService "todo-list/pkg/analyticsService/service"
	analyticworker "todo-list/pkg/analyticsService/worker"
	"todo-list/pkg/database"
	notifyServices "todo-list/pkg/notifyService/service"
	"todo-list/pkg/notifyService/worker"
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

	if err := config.InitCfg(); err != nil {
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
	log.Printf("init db on %s", viper.GetString("db.host"))

	repo := repository.NewRepository(db)
	s := service.NewService(repo)
	h := handler.NewHandler(s)

	srv := &server.Server{}
	serverConfig := server.SrvConfig{
		Port:              config.GetPort(),
		ReadHeaderTimeout: viper.GetDuration("server.read_header_timeout"),
		WriteTimeout:      viper.GetDuration("server.write_timeout"),
		IdleTimeout:       viper.GetDuration("server.idle_timeout"),
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		if err := srv.RunServer(serverConfig, h.MapRoutes()); err != nil {
			logrus.Fatalf("cant run server: %v", err)
		}
	}()

	logrus.Println("server started")

	notifierClient := notifyServices.NewNotifyClient(viper.GetString("notification.address"))
	deadlineWorker := worker.NewDeadlineWorker(repo, notifierClient)

	go func() {
		if err := deadlineWorker.Start(ctx); err != nil {
			logrus.Fatalf("cant start deadline worker: %v", err)
		}
	}()

	analyticClient, err := analyticsService.NewAnalyticsClient(viper.GetString("analytics.grpc_port"))
	if err != nil {
		logrus.Printf("Can't create analytics client: %v", err)
	}
	defer analyticClient.CloseConn()

	analyticWorker := analyticworker.NewAnalyticWorker(repo, analyticClient)

	go func() {
		if err := analyticWorker.Start(ctx); err != nil {
			logrus.Fatalf("cant start analytics worker: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Println("Shutting down server...")

	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.ShutDown(shutdownCtx); err != nil {
		logrus.Fatalf("cant shutdown server: %v", err)
	}

	logrus.Println("server shutdown")

}
