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
