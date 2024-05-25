package dto

type GuestUploadRequest struct {
	Name string `json:"name" form:"name" validate:"required,min=3" example:""`
}

type GuestUploadResponse struct {
	Fieldname string `json:"fieldname" example:""`
	MimeType  string `json:"mimetype" example:""`
	Size      int    `json:"size" example:"0"`
	Name      string `json:"name" example:""`
}
