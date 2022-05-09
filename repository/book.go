package repository

import (
	"redigo-example/entity"

	"gorm.io/gorm"
)

type BookRepositoryIn interface {
	GetAll() (data []entity.Book, err error)
	GetDetail(id int) (data entity.Book, err error)
	Create(data entity.Book) (err error)
	Update(id int, data entity.Book) (err error)
	Delete(id int) (err error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) bookRepository {
	return bookRepository{db}
}

func (r *bookRepository) GetAll() (data []entity.Book, err error) {
	err = r.db.Find(&data).Error

	return
}

func (r *bookRepository) GetDetail(id int) (data entity.Book, err error) {
	err = r.db.Where("id = ?", id).Find(&data).Error

	return
}

func (r *bookRepository) Create(data entity.Book) (err error) {
	err = r.db.Create(&data).Error

	return
}

func (r *bookRepository) Update(id int, data entity.Book) (err error) {
	err = r.db.Where("id = ?", id).Updates(&data).Error

	return
}

func (r *bookRepository) Delete(id int) (err error) {
	book := entity.Book{}
	err = r.db.Where("id = ?", id).Delete(&book).Error

	return
}
