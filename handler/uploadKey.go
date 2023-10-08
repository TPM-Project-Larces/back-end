package handler

import (
	"github.com/gin-gonic/gin"
)

// @BasePath /

// @Summary upload key

// @Description uploads a public key
// @Tags Server operations
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} string "key_uploaded"
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /upload_key [post]
func UploadKey(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 400, "bad_request", err)
	}

	// Salva a chave no servidor
	err = ctx.SaveUploadedFile(file, "./key/"+file.Filename)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
	}

	response(ctx, 200, "keys_generated", err)
}
