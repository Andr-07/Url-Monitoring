package url

type RegisterRequest struct {
	Address string `json:"address" validate:"required,url"`
	Interval int `json:"interval" validate:"required"`
}

type PostResponse struct {
	Message string `json: "message"`
}