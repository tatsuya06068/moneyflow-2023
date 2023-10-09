package entity

import (
	"context"
	"time"
)

type BoPList struct {
	BoPID       string
	PaymentName string
	PaymentDate time.Time
	TotalAmount int
	CategoryId  int
}

type Token struct {
	UserId string
}

type IBoPInteractor interface {
	BoPList(ctx context.Context, token string) ([]BoPList, error)
}
