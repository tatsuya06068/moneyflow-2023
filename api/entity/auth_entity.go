package entity

type SignupRequest struct {
	UserName string `json:"name"`
	Password string `json:"password"`
}
