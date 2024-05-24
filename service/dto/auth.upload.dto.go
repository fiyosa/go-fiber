package dto

import "mime/multipart"

type AuthUploadPayload struct {
	File *multipart.FileHeader `form:"file" json:"file" binding:"required"`
	Name string                `form:"name" json:"name" binding:"required,min=3" example:""`
}

type AuthUploadResponse struct {
	Fieldname    string `json:"fieldname" example:""`
	OriginalName string `json:"originalname" example:""`
	Encoding     string `json:"encoding" example:""`
	MimeType     string `json:"mimetype" example:""`
	Size         int    `json:"size" example:""`
}
