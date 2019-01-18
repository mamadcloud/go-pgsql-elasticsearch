package entities

import (
	"time"
)

type DeltaUpdate struct {
	ID int
	ItemID int
	Action string
	TableName string
	Timestamp time.Time
}