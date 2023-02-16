package admin

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterInputAdmin) (Admin, error)
	Login(input LoginInputAdmin) (Admin, error)
	FindByID(ID int) (Admin, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input RegisterInputAdmin) (Admin, error) {
	admin := Admin{}
	admin.Name.String = input.Name
	admin.Email.String = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return admin, err
	}
	admin.Password.String = string(passwordHash)
	newAdmin, err := s.repository.Save(admin)
	if err != nil {
		return newAdmin, err
	}
	return newAdmin, nil
}

func (s *service) Login(input LoginInputAdmin) (Admin, error) {

	admin, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return admin, err
	}
	if admin.ID.Int64 == 0 {
		return admin, errors.New("no admin found on that email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password.String), []byte(input.Password))
	if err != nil {
		return admin, err
	}
	return admin, nil
}

func (s *service) FindByID(ID int) (Admin, error) {

	admin, err := s.repository.FindByID(ID)
	if err != nil {
		return admin, err
	}
	
	return admin, nil
}

