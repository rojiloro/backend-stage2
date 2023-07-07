package stationdto

type CreateStationRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type UpdateStationRequest struct {
	Name string `json:"name" form:"name"`
}

