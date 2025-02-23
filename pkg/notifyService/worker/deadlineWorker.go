package worker

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
	notifyServices "todo-list/pkg/notifyService/service"
	"todo-list/todo/repository"
)

type DeadlineWorker struct {
	repo    *repository.Repository
	service notifyServices.NotifyService
}

func NewDeadlineWorker(repo *repository.Repository, service notifyServices.NotifyService) *DeadlineWorker {
	return &DeadlineWorker{repo: repo, service: service}
}

func (c *DeadlineWorker) Start(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logrus.Println("deadlineWorker shutting down")
			return ctx.Err()

		case <-ticker.C:
			usersWithDeadlines, err := c.repo.GetDeadlineItems()
			if err != nil {
				return fmt.Errorf("error fetching users with expiring deadlines: %v", err)
			}

			if len(usersWithDeadlines) > 0 {
				logrus.Printf("Found %d users with expiring deadlines", len(usersWithDeadlines))
			}

			if err := c.service.PushToNotifyService(usersWithDeadlines); err != nil {
				return fmt.Errorf("error sending deadline to notification service: %v", err)
			}

			logrus.Info("successfully sent data to  notification service")
		}
	}
}
