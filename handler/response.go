package handler

import "github.com/TPM-Project-Larces/back-end.git/schemas"

type ListFilesResponse struct {
	Message string                          `bson:"message"`
	Data    []schemas.EncryptedFileResponse `bson:"data"`
}
type ShowFileResponse struct {
	Message string                        `bson:"message"`
	Data    schemas.EncryptedFileResponse `bson:"data"`
}
type DeleteFileResponse struct {
	Message string                        `bson:"message"`
	Data    schemas.EncryptedFileResponse `bson:"data"`
}
