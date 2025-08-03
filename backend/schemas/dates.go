package schemas

type DateChoiceRequest struct {
	Dates []string `json:"dates" validate:"required,dive,datetime=2006-01-02"`
}

type UserDateChoiceResponse struct {
	Dates []string `json:"dates"`
}