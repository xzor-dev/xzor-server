package user

// Service controls various functions related to users.
type Service struct {
	user *User
}

// NewService creates a new instance of the user service.
func NewService() *Service {
	return &Service{}
}
