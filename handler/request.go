package handler

type DeleteFileRequest struct {
	Filename string `bson:"filename"`
}
