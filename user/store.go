package main

import (
	"errors"

	"github.com/dongsu8142/blog-common/database"
	"gorm.io/gorm"
)

type store struct {
	db *gorm.DB
}

func NewStore(host, user, password, dbname, port string) *store {
	db, err := database.ConnectDatabase(host, user, password, dbname, port)
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(&database.User{})
	db.AutoMigrate(&database.User{})
	return &store{db}
}

func (s *store) Register(username, email, password string) error {
	result := s.db.Create(&database.User{
		Username: username,
		Email:    email,
		Password: password,
	})
	if result.Error != nil {
		return errors.New("username already exists")
	}
	return nil
}

func (s *store) Login(username string) (*database.User, error) {
	var user database.User
	result := s.db.First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, errors.New("not found, invalid username")
	}
	return &user, nil
}

func (s *store) GetUserByID(id int) (*database.User, error) {
	var user database.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		return nil, errors.New("not found, invalid id")
	}
	return &user, nil
}