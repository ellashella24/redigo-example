package entity

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func UserSeeder(db *gorm.DB) {
	for i := 1; i <= 1000; i++ {
		db.Create(&User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: faker.Password(),
		})
	}
}
