package route

import (
	"go-fiber/pkg/helper"
	"go-fiber/service/dto"

	"github.com/gofiber/fiber/v2"
)

// @Summary 	Upload
// @Description Upload
// @Tags 		Guest
// @Accept 		json
// @Produce 	json
// @Param   	file formData  file true "file"
// @Param   	name formData string true "name"
// @Success 	200 {object} dto.GuestUploadResponse "ok"
// @Router 		/auth/upload [post]
func GuestUpload(c *fiber.Ctx) error {
	validated := dto.GuestUploadRequest{}
	if check, err := helper.Validate(c, &validated); check {
		return err
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return helper.SendError(
			c,
			"File is required",
			fiber.StatusBadRequest,
		)
	}

	file, err := fileHeader.Open()
	if err != nil {
		return helper.SendError(
			c,
			"Unable to open file",
			fiber.StatusBadRequest,
		)
	}
	defer file.Close()

	mimeType, err := helper.GetMimeType(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to detect MimeType",
		})
	}

	res := dto.GuestUploadResponse{
		Fieldname: fileHeader.Filename,
		MimeType:  mimeType,
		Size:      int(fileHeader.Size),
		Name:      validated.Name,
	}

	return helper.SendCustom(c, res)
}
