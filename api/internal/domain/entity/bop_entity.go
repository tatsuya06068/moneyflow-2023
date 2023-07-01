package entity

import (
	"context"
	"time"
)

type BoPList struct {
	bop_id       string
	payment_name string
	payment_date time.Time
	total_amount int
	category_id  int
}

type IBoPInteractor interface {
	BoPList(ctx context.Context, jwt string)
}
