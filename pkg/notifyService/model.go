package notifyService

import "github.com/google/uuid"

type TaskDeadlineInfo struct {
	UserID      uuid.UUID `json:"user_id"`
	Email       string    `json:"email"`
	ItemID      int       `json:"item_id"`
	Description string    `json:"description"`
}
