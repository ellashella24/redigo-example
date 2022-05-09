package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"redigo-example/config"
	"redigo-example/entity"
	"redigo-example/repository"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserServiceIn interface {
	GetAll() (data []entity.User, httpStatus int, err error)
	GetAllWithoutRedis() (data []entity.User, httpStatus int, err error)
	GetDetail(c echo.Context) (data entity.User, httpStatus int, err error)
	Create(c echo.Context) (httpStatus int, err error)
	Update(c echo.Context) (httpStatus int, err error)
	Delete(c echo.Context) (httpStatus int, err error)
}

type userService struct {
	userRepository  repository.UserRepositoryIn
	redisRepository repository.RedisRepositoryIn
}

func NewUserService(userRepository repository.UserRepositoryIn, redisRepository repository.RedisRepositoryIn) *userService {
	return &userService{userRepository, redisRepository}
}

func (s *userService) GetAll() (data []entity.User, httpStatus int, err error) {
	_, _, redisConf := config.InitConfig()
	key := redisConf.UserKey
	exist, _ := s.redisRepository.CheckExist(key)

	if exist {
		cacheData, err := s.redisRepository.GetCache(key)
		if err != nil {
			httpStatus = http.StatusInternalServerError
			return data, httpStatus, errors.New("can't get data from cache")
		}

		err = json.Unmarshal(cacheData, &data)
		if err != nil {
			httpStatus = http.StatusInternalServerError
			return data, httpStatus, errors.New("can't unmarshal data from cache")
		}

		httpStatus = http.StatusOK
		return data, httpStatus, err
	}

	data, err = s.userRepository.GetAll()
	if err != nil {
		err = errors.New("can't get all user data")
		log.Error("can't get all user data")
		httpStatus = http.StatusInternalServerError

		return
	}

	cacheData, err := json.Marshal(data)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return nil, httpStatus, err
	}

	err = s.redisRepository.WriteCache(key, cacheData)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return nil, httpStatus, err
	}

	httpStatus = http.StatusOK
	return
}

func (s *userService) GetAllWithoutRedis() (data []entity.User, httpStatus int, err error) {
	data, err = s.userRepository.GetAll()
	if err != nil {
		err = errors.New("can't get all user data")
		log.Error("can't get all user data")
		httpStatus = http.StatusInternalServerError

		return
	}

	httpStatus = http.StatusOK
	return
}

func (s *userService) GetDetail(c echo.Context) (data entity.User, httpStatus int, err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("can't get user id")
		log.Error("can't get user id")
		httpStatus = http.StatusBadRequest

		return
	}

	data, err = s.userRepository.GetDetail(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("user not found")
			log.Error("user not found")
			httpStatus = http.StatusNotFound

			return
		}
		err = errors.New("can't get user data")
		log.Error("can't get user data")
		httpStatus = http.StatusInternalServerError

		return
	}

	httpStatus = http.StatusOK
	return
}

func (s *userService) Create(c echo.Context) (httpStatus int, err error) {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	data := entity.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err = s.userRepository.Create(data)
	if err != nil {
		err = errors.New("can't create user data")
		log.Error("can't create user data")
		httpStatus = http.StatusInternalServerError
		return
	}

	_, _, redisConf := config.InitConfig()
	key := redisConf.UserKey
	err = s.redisRepository.DeleteCache(key)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return httpStatus, err
	}

	httpStatus = http.StatusOK
	return
}

func (s *userService) Update(c echo.Context) (httpStatus int, err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("can't get user id")
		log.Error("can't get user id")
		httpStatus = http.StatusBadRequest

		return
	}

	_, err = s.userRepository.GetDetail(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("user not found")
			log.Error("user not found")
			httpStatus = http.StatusNotFound

			return
		}
		err = errors.New("can't get user data")
		log.Error("can't get user data")
		httpStatus = http.StatusInternalServerError

		return
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	data := entity.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err = s.userRepository.Update(id, data)
	if err != nil {
		err = errors.New("can't update user data")
		log.Error("can't update user data")
		httpStatus = http.StatusInternalServerError
		return httpStatus, err
	}

	_, _, redisConf := config.InitConfig()
	key := redisConf.UserKey
	err = s.redisRepository.DeleteCache(key)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return httpStatus, err
	}

	httpStatus = http.StatusOK
	return
}

func (s *userService) Delete(c echo.Context) (httpStatus int, err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New("can't get user id")
		log.Error("can't get user id")
		httpStatus = http.StatusBadRequest

		return
	}

	_, err = s.userRepository.GetDetail(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("user not found")
			log.Error("user not found")
			httpStatus = http.StatusNotFound

			return
		}
		err = errors.New("can't get user data")
		log.Error("can't get user data")
		httpStatus = http.StatusInternalServerError

		return
	}

	err = s.userRepository.Delete(id)
	if err != nil {
		err = errors.New("can't delete user data")
		log.Error("can't delete user data")
		httpStatus = http.StatusInternalServerError
		return httpStatus, err
	}

	_, _, redisConf := config.InitConfig()
	key := redisConf.UserKey
	err = s.redisRepository.DeleteCache(key)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return httpStatus, err
	}

	httpStatus = http.StatusOK
	return
}
