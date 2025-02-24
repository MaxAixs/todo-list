package service

import (
	"fmt"
	"github.com/MaxAixs/protos/gen/api/gen/api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"todo-list/pkg/analyticsService"
)

type AnalyticsClient struct {
	client api.AnalyticsDataClient
	conn   *grpc.ClientConn
}

func NewAnalyticsClient(port string) (*AnalyticsClient, error) {
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("NewAnalyticsClient: %v", err)
	}

	gRPCClient := api.NewAnalyticsDataClient(conn)

	return &AnalyticsClient{conn: conn, client: gRPCClient}, nil
}

func (c *AnalyticsClient) SendDoneItems(tasks []analyticsService.TaskDoneItem) error {
	gRPCModels, err := convertToGRPCModel(tasks)
	if err != nil {
		return fmt.Errorf("could not convert tasks to GRPC model: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.client.SaveDoneTasks(ctx, gRPCModels)
	if err != nil {
		return fmt.Errorf("could not send done tasks data: %v", err)
	}

	logrus.Printf("Response from analyticsService: %v", resp)

	return nil
}

func convertToGRPCModel(tasks []analyticsService.TaskDoneItem) (*api.TaskDoneItems, error) {
	var gRPCTasks []*api.TaskDoneItem

	for _, task := range tasks {
		gRPCTasks = append(gRPCTasks, &api.TaskDoneItem{
			UserId:    task.UserID.String(),
			Email:     task.Email,
			ItemId:    task.ItemID,
			CreatedAt: timestamppb.New(task.CreatedAt),
		})
	}

	req := &api.TaskDoneItems{
		Items: gRPCTasks,
	}

	return req, nil
}

func (c *AnalyticsClient) CloseConn() {
	c.conn.Close()
}
