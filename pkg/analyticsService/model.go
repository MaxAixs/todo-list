package analyticsService

import (
	"github.com/google/uuid"
	"time"
)

type TaskDoneItem struct {
	UserID    uuid.UUID
	Email     string
	ItemID    int32
	CreatedAt time.Time
}
