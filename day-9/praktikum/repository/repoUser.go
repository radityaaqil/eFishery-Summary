package repository

import (
	"day9/model"
	"log"

	"gorm.io/gorm"
)

type Repository_User interface {
	GetAllUsers() ([]model.User, error)
	GetUserById(id int) (model.User, error)
	CreateUser(name string, age int) (model.User, error)
	UpdateUser(id int, name string, age int) (model.User, error)
}

type RepositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *RepositoryUser {
	return &RepositoryUser{db}
}

func (r *RepositoryUser) GetAllUsers() ([]model.User, error) {
	var users []model.User

	result := r.db.Order("id asc").Find(&users)

	err := result.Error

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func (r *RepositoryUser) GetUserById(id int) (model.User, error) {
	var user model.User

	result := r.db.First(&user, id)

	err := result.Error

	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func (r *RepositoryUser) CreateUser(name string, age int) (model.User, error) {
	user := model.User{Name: name, Age: age}

	result := r.db.Create(&user)

	err := result.Error

	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func (r *RepositoryUser) UpdateUser(id int, name string, age int) (model.User, error) {
	var user model.User

	r.db.First(&user, id)

	user.Name = name
	user.Age = age

	result := r.db.Updates(&user)

	err := result.Error

	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func (r *RepositoryUser) DeleteUser(id int) (model.User, error) {
	var user model.User

	r.db.First(&user, id)

	result := r.db.Delete(&user)

	err := result.Error

	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}
