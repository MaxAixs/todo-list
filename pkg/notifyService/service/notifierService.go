package service

import "todo-list/pkg/notifyService"

type NotifyService interface {
	PushToNotifyService(users []notifyService.TaskDeadlineInfo) error
}
