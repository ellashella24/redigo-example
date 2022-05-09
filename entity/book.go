package entity

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string
	Author string
}

func BookSeeder(db *gorm.DB) {
	for i := 1; i <= 1000; i++ {
		db.Create(&Book{
			Title:  faker.Word(),
			Author: faker.Name(),
		})
	}
}
