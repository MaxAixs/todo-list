package service

import "todo-list/pkg/analyticsService"

type AnalyticsGrpcService interface {
	SendDoneItems(tasks []analyticsService.TaskDoneItem) error
	CloseConn()
}
