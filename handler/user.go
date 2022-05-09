package handler

import (
	"redigo-example/response"
	"redigo-example/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserServiceIn
}

func NewUserHandler(userService service.UserServiceIn) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetAll(c echo.Context) (err error) {
	data, httpStatus, err := h.userService.GetAll()

	if err != nil {
		return c.JSON(httpStatus, err)
	}

	responses := []response.UserResponse{}
	for _, unitData := range data {
		responses = append(responses, response.UserResponse{
			ID:       int(unitData.ID),
			Name:     unitData.Name,
			Email:    unitData.Email,
			Password: unitData.Password,
		})
	}

	return c.JSON(httpStatus, responses)
}

func (h *UserHandler) GetAllWithoutRedis(c echo.Context) (err error) {
	data, httpStatus, err := h.userService.GetAllWithoutRedis()

	if err != nil {
		return c.JSON(httpStatus, err)
	}

	responses := []response.UserResponse{}
	for _, unitData := range data {
		responses = append(responses, response.UserResponse{
			ID:       int(unitData.ID),
			Name:     unitData.Name,
			Email:    unitData.Email,
			Password: unitData.Password,
		})
	}

	return c.JSON(httpStatus, responses)
}

func (h *UserHandler) GetDetail(c echo.Context) (err error) {
	data, httpStatus, err := h.userService.GetDetail(c)

	if err != nil {
		return c.JSON(httpStatus, err)
	}

	responses := response.UserResponse{
		ID:       int(data.ID),
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	return c.JSON(httpStatus, responses)
}

func (h *UserHandler) Create(c echo.Context) (err error) {
	httpStatus, err := h.userService.Create(c)
	if err != nil {
		return c.JSON(httpStatus, response.MessageResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(httpStatus, response.MessageResponse{
		Message: "User berhasil dibuat",
	})
}

func (h *UserHandler) Update(c echo.Context) (err error) {
	httpStatus, err := h.userService.Update(c)
	if err != nil {
		return c.JSON(httpStatus, response.MessageResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(httpStatus, response.MessageResponse{
		Message: "User berhasil diupdate",
	})
}

func (h *UserHandler) Delete(c echo.Context) (err error) {
	httpStatus, err := h.userService.Delete(c)
	if err != nil {
		return c.JSON(httpStatus, response.MessageResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(httpStatus, response.MessageResponse{
		Message: "User berhasil dihapus",
	})
}
