package helper

import (
	"mime/multipart"
	"net/http"
)

func GetMimeType(file multipart.File) (string, error) {
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return "", err
	}
	// Reset the read pointer
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}
	return http.DetectContentType(buffer), nil
}
