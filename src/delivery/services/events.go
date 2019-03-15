package services

import (
	"time"
)

type Event struct {
	timestamp time.Time
	event     string
}
