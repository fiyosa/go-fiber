package dto

type GuestLoginRequest struct {
	Username string `json:"username" form:"username" validate:"required,min=3" example:""`
	Password string `json:"password" form:"password" validate:"required,min=3" example:""`
}

type GuestLoginResponse struct {
	Token string `json:"token" example:""`
}
