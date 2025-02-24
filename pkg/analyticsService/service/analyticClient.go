package service

import "todo-list/pkg/analyticsService"

type AnalyticsGrpcClient interface {
	SendDoneItems(tasks []analyticsService.TaskDoneItem) error
	CloseConn()
}
