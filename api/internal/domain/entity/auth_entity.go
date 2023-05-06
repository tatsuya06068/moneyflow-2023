package entity

type SignupRequest struct {
	UserName string `json:"name"`
	Password string `json:"password"`
}

// type Input func(o port.UserOutputPort, u port.UserRepository) port.UserInputPort
// type Repo func(c *sql.DB) port.UserRepository
