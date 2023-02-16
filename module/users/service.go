package users

import "golang.org/x/crypto/bcrypt"

type Service interface {
	Create(input InputUsers) (Users, error)
	Read() ([]Users, error)
	Update(input UpdateUsers) (Users, error)
	Delete(ID int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input InputUsers) (Users, error) {
	var users Users
	users.Name.String=input.Name
	users.Email.String=input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return users, err
	}
	users.Password.String = string(passwordHash)
	data, err := s.repository.Create(users)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *service) Read() ([]Users, error) {
	data, err := s.repository.Read()
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *service) Update(input UpdateUsers) (Users, error) {
	var users Users
	users.ID.Int64 = int64(input.ID)
	users.Name.String=input.Name
	users.Email.String=input.Email
	if input.Password!=""{
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return users, err
		}
		users.Password.String = string(passwordHash)
	}
	data, err := s.repository.Update(users)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *service) Delete(ID int) error {
	err := s.repository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
