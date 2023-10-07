package database

import (
	"context"

	"github.com/tatsuya06068/moneyflow-2023/internal/domain/entity"
)

type BoPDBGateway struct {
	ISqlHandler
}

func (bdg BoPDBGateway) GetBoPList(ctx context.Context, userId int) ([]entity.BoPList, error) {

	row, err := bdg.Query("SELECT t_bop_id, payment_naem, payment_date, total_amount, m_bop_category_id WHERE user_id = ?", userId)

	if err != nil {
		return nil, err
	}

	bopList := []entity.BoPList{}

	defer row.Close()

	for row.Next() {
		var bop entity.BoPList

		err := row.Scan(&bop.BoPID, &bop.PaymentName, &bop.PaymentDate, &bop.TotalAmount, &bop.CategoryId)

		if err != nil {
			return nil, err
		}

		bopList = append(bopList, bop)
	}
	return bopList, nil
}
