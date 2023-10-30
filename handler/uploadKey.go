package handler

import (
	"github.com/gin-gonic/gin"
)

// @BasePath /
// @Summary Upload a public key
// @Description Uploads a public key
// @Tags Encryption
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} string "key_uploaded"
// @Failure 400 {string} string "bad_request"
// @Failure 500 {string} string "internal_server_error"
// @Router /encryption/upload_key [post]
func UploadKey(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	err = ctx.SaveUploadedFile(file, "./key/"+file.Filename)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	response(ctx, 200, "key_uploaded", err)
}
