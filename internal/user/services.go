package user

import "log"

type Service interface {
	Create( first_name, last_name, email, phone string) (User, error)
}

type service struct {}

//Este es un patrón fabrica. Encapsula la creación de un objeto service.
func NewService() Service {
	return &service{}
}

func (s *service) Create(first_name, last_name, email, phone string) (User, error) {
	log.Print("Creating user service")
	return User{
		ID:        "1",
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Phone:     phone,
	}, nil
}