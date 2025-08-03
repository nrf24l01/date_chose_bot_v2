package schemas

type DateChoiceRequest struct {
	Dates []string `json:"dates" validate:"required,dive,datetime=2006-01-02"`
}

type UserDateChoiceResponse struct {
	Dates []string `json:"dates"`
}

type VotedUser struct {
	UserID   int64    `json:"user_id"`
	UserName string   `json:"user_name"`
	Dates    []string `json:"dates"`
}