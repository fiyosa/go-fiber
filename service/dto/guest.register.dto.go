package dto

type GuestRegisterRequest struct {
	Username string `json:"username" form:"username" from:"username" validate:"required,min=3" example:""`
	Password string `json:"password" form:"password" from:"password" validate:"required,min=3" example:""`
	Name     string `json:"name" form:"name" from:"name" validate:"required,min=3" example:""`
}

type GuestRegisterResponse struct {
	Data    GuestRegisterDataResponse `json:"data"`
	Message string                    `json:"message" example:""`
}

type GuestRegisterDataResponse struct {
	Id        string `json:"id" example:""`
	Username  string `json:"username" example:""`
	Name      string `json:"name" example:""`
	CreatedAt string `json:"created_at" example:""`
	UpdatedAt string `json:"updated_at" example:""`
}
