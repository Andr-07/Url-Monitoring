package url

type RegisterRequest struct {
	Address string `json:"address" validate:"required,url"`
}

type PostResponse struct {
	Message string `json: "message"`
}
