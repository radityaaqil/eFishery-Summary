package service

import (
	"day7/model"
	"day7/repository"
	"log"
)

type Service_User interface {
	GetAllUsersService() ([]model.User, error)
	GetUserByIdService(id int) (model.User, error)
	CreateUserService(name string, age int) (model.User, error)
	UpdateUserService(id int, name string, age int) (model.User, error)
	DeleteUserService(id int) (model.User, error)
}

type ServiceUser struct {
	repository1 repository.RepositoryUser
}

func NewUserService(repository1 repository.RepositoryUser) *ServiceUser {
	return &ServiceUser{repository1}
}

func (s *ServiceUser) GetAllUsersService() ([]model.User, error) {
	users, err := s.repository1.GetAllUsers()

	if err != nil {
		log.Println("Get Users Service Error", err)
		return nil, err
	}

	return users, nil
}

func (s *ServiceUser) GetUserByIdService(id int) (model.User, error) {
	user, err := s.repository1.GetUserById(id)

	if err != nil {
		log.Println("Get User By Id Service Error", err)
		return user, err
	}

	return user, nil
}

func (s *ServiceUser) CreateUserService(name string, age int) (model.User, error) {
	user, err := s.repository1.CreateUser(name, age)

	if err != nil {
		log.Println("Create User Service Error", err)
		return user, err
	}

	return user, nil
}

func (s *ServiceUser) UpdateUserService(id int, name string, age int) (model.User, error) {
	user, err := s.repository1.UpdateUser(id, name, age)

	if err != nil {
		log.Println("Update User Service Error", err)
		return user, err
	}

	return user, nil
}

func (s *ServiceUser) DeleteUserService(id int) (model.User, error) {
	user, err := s.repository1.DeleteUser(id)

	if err != nil {
		log.Println("Delete User Service Error", err)
		return user, err
	}

	return user, nil
}
