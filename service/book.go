package service

import (
	"errors"
	"net/http"
	"redigo-example/entity"
	"redigo-example/repository"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type BookServiceIn interface {
	GetAll() (data []entity.Book, httpStatus int, err error)
	GetDetail(c echo.Context) (data entity.Book, httpStatus int, err error)
	Create(c echo.Context) (httpStatus int, err error)
	Update(c echo.Context) (httpStatus int, err error)
	Delete(c echo.Context) (httpStatus int, err error)
}

type bookService struct {
	bookRepository repository.BookRepositoryIn
}

func NewBookService(bookRepository repository.BookRepositoryIn) *bookService {
	return &bookService{bookRepository}
}

func (s *bookService) GetAll() (data []entity.Book, httpStatus int, err error) {
	data, err = s.bookRepository.GetAll()

	if err != nil {
		err = errors.New("can't get all book data")
		log.Error("can't get all book data")
		httpStatus = http.StatusInternalServerError

		return
	}

	httpStatus = http.StatusOK
	return
}

func (s *bookService) GetDetail(c echo.Context) (data entity.Book, httpStatus int, err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("can't get book id")
		log.Error("can't get book id")
		httpStatus = http.StatusBadRequest

		return
	}

	data, err = s.bookRepository.GetDetail(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("book not found")
			log.Error("book not found")
			httpStatus = http.StatusNotFound

			return
		}
		err = errors.New("can't get book data")
		log.Error("can't get book data")
		httpStatus = http.StatusInternalServerError

		return
	}

	httpStatus = http.StatusOK
	return
}

func (s *bookService) Create(c echo.Context) (httpStatus int, err error) {
	title := c.FormValue("title")
	author := c.FormValue("author")

	data := entity.Book{
		Title:  title,
		Author: author,
	}

	err = s.bookRepository.Create(data)
	if err != nil {
		err = errors.New("can't create book data")
		log.Error("can't create book data")
		httpStatus = http.StatusInternalServerError
	}

	httpStatus = http.StatusOK
	return
}

func (s *bookService) Update(c echo.Context) (httpStatus int, err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("can't get book id")
		log.Error("can't get book id")
		httpStatus = http.StatusBadRequest

		return
	}

	_, err = s.bookRepository.GetDetail(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("book not found")
			log.Error("book not found")
			httpStatus = http.StatusNotFound

			return
		}
		err = errors.New("can't get book data")
		log.Error("can't get book data")
		httpStatus = http.StatusInternalServerError

		return
	}

	title := c.FormValue("title")
	author := c.FormValue("author")

	data := entity.Book{
		Title:  title,
		Author: author,
	}

	err = s.bookRepository.Update(id, data)
	if err != nil {
		err = errors.New("can't update book data")
		log.Error("can't update book data")
		httpStatus = http.StatusInternalServerError
	}

	httpStatus = http.StatusOK
	return
}

func (s *bookService) Delete(c echo.Context) (httpStatus int, err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("can't get book id")
		log.Error("can't get book id")
		httpStatus = http.StatusBadRequest

		return
	}

	_, err = s.bookRepository.GetDetail(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("book not found")
			log.Error("book not found")
			httpStatus = http.StatusNotFound

			return
		}
		err = errors.New("can't get book data")
		log.Error("can't get book data")
		httpStatus = http.StatusInternalServerError

		return
	}

	err = s.bookRepository.Delete(id)
	if err != nil {
		err = errors.New("can't delete book data")
		log.Error("can't delete book data")
		httpStatus = http.StatusInternalServerError
	}

	httpStatus = http.StatusOK
	return
}
