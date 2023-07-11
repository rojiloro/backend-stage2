package handlers

import (
	dto "LandTicket-Backend/dto/result"
	stationdto "LandTicket-Backend/dto/station"
	"LandTicket-Backend/models"
	"LandTicket-Backend/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerStation struct {
	StationRepository repositories.StationRepository
}

func HandlerStation(StationRepository repositories.StationRepository) *handlerStation {
	return &handlerStation{StationRepository}
}

func (h *handlerStation) CreateStation(c echo.Context) error {
	
	request := new(stationdto.CreateStationRequest)
	if err:=c.Bind(request) ;err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	station := models.Station{
		Name: request.Name,
	}

	data, err := h.StationRepository.CreateStation(station)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseStation(data)})
}

func (h *handlerStation) FindStation(c echo.Context) error {
	station, err := h.StationRepository.FindStation()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: station})
}

func convertResponseStation(s models.Station) stationdto.StationResponse{
	return stationdto.StationResponse{
		ID: s.ID,
		Name: s.Name,
	}
}