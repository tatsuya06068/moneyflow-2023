package jwt

import (
	"fmt"

	"github.com/tatsuya06068/moneyflow-2023/internal/constants"
)

type BopJwt struct {
	IJwtHandler
}

func (bj *BopJwt) GetUserId(token string) (int, error) {
	claim, verErr := bj.IJwtHandler.VerifyToken(token)
	if verErr != nil {
		return 0, fmt.Errorf(constants.ErrorFormat, "aj.IJwtHandler.VerifyToken", verErr)
	}

	mClaim, payErr := claim.GetPayload()
	if payErr != nil {
		return 0, fmt.Errorf(constants.ErrorFormat, "claim.GetPayload()", payErr)
	}
	fmt.Printf("claim %#v\n", mClaim)
	if userId, ok := mClaim["user_id"]; ok {
		return int(userId.(float64)), nil
	}
	return 0, fmt.Errorf(constants.ErrorFormat, "claim.GetPayload()", "Not user_id")
}
