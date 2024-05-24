package dto

type AuthUserResponse struct {
	Data    AuthUserDataResponse `json:"data"`
	Message string               `json:"message" example:""`
}

type AuthUserDataResponse struct {
	Id        string `json:"id" example:""`
	Username  string `json:"username" example:""`
	Name      string `json:"name" example:""`
	CreatedAt string `json:"created_at" example:""`
	UpdatedAt string `json:"updated_at" example:""`
}
