package dto

import "go-fiber/pkg/helper"

type UserIndexResponse struct {
	Data       []UserIndexDataResponse `json:"data"`
	Pagination helper.Paginate         `json:"pagination"`
	Message    string                  `json:"message" example:""`
}

type UserIndexDataResponse struct {
	Id        string `json:"id" example:""`
	Username  string `json:"username" example:""`
	Name      string `json:"name" example:""`
	CreatedAt string `json:"created_at" example:""`
	UpdatedAt string `json:"updated_at" example:""`
}
