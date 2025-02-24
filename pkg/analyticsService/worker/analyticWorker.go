package worker

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
	"todo-list/pkg/analyticsService/service"
	"todo-list/todo/repository"
)

type AnalyticWorker struct {
	repo       *repository.Repository
	gRPCClient service.AnalyticsGrpcClient
}

func NewAnalyticWorker(repo *repository.Repository, gRPCClient service.AnalyticsGrpcClient) *AnalyticWorker {
	return &AnalyticWorker{repo: repo, gRPCClient: gRPCClient}
}

func (a *AnalyticWorker) Start(ctx context.Context) error {
	ticker := time.NewTicker(60 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logrus.Print("analyticWorker shutting down")
			return ctx.Err()

		case <-ticker.C:
			usersWithCompletedTasks, err := a.repo.GetUsersWithDoneItem()
			if err != nil {
				return fmt.Errorf("error fetching users with done item: %w", err)
			}

			if len(usersWithCompletedTasks) > 0 {
				logrus.Printf("analyticWorker got users %v with completed tasks", len(usersWithCompletedTasks))

				if err := a.gRPCClient.SendDoneItems(usersWithCompletedTasks); err != nil {
					return fmt.Errorf("error sending data to analytics service: %w", err)
				}
			}
		}
	}
}
