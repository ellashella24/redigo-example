package repository

import (
	"redigo-example/entity"

	"gorm.io/gorm"
)

type UserRepositoryIn interface {
	GetAll() (data []entity.User, err error)
	GetDetail(id int) (data entity.User, err error)
	Create(data entity.User) (err error)
	Update(id int, data entity.User) (err error)
	Delete(id int) (err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll() (data []entity.User, err error) {
	err = r.db.Find(&data).Error

	return
}

func (r *userRepository) GetDetail(id int) (data entity.User, err error) {
	err = r.db.Where("id = ?", id).Find(&data).Error

	return
}

func (r *userRepository) Create(data entity.User) (err error) {
	err = r.db.Create(&data).Error

	return
}

func (r *userRepository) Update(id int, data entity.User) (err error) {
	err = r.db.Where("id = ?", id).Updates(&data).Error

	return
}

func (r *userRepository) Delete(id int) (err error) {
	user := entity.User{}
	err = r.db.Where("id = ?", id).Delete(&user).Error

	return
}
