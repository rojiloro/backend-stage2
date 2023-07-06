package handlers

import (
	usersdto "LandTicket-Backend/dto/auth"
	dto "LandTicket-Backend/dto/result"
	"LandTicket-Backend/models"
	"LandTicket-Backend/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type authhandler struct {
	userRepository repositories.UserRepository
}

func (h *handler) CreateUser(c echo.Context) error{
	request := new(usersdto.userRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code : http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Email: request.Email,
		Password: request.Password,
	}

	data, err := h.UserRepository.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data) })

}