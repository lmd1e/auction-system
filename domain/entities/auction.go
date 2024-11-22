package entities

import "time"

type Auction struct {
	ID        int
	LotID     int
	StartTime time.Time
	EndTime   time.Time
}
