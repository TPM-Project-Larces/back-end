package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
)

func response(ctx *gin.Context, code int, message string, err error) {
	response := gin.H{
		"code": code,
	}

	if message != "" {
		response["message"] = message
	}

	if err != nil {
		response["error"] = err.Error()
	}

	ctx.JSON(code, response)
}

func sendFile(fileName string, url string) error {

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Println("teste 1")
	// Create a buffer to store the request body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create a form file part in the request body
	part, err := writer.CreateFormFile("arquivo", fileName)
	if err != nil {
		return err
	}

	// Copy the contents of the file to the form file part
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	writer.Close()

	// Create a new HTTP POST request with the request body
	request, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType()) // Set the request content type

	// Send the HTTP request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	// Remove the file if it's not the public key file
	//if fileName != "./key/public_key.pem" {
	//	os.Remove(fileName)
	//}

	// Check the HTTP response status code
	if response.StatusCode == http.StatusOK {
		return nil
	} else {
		return errors.New("file not sentt")
	}
}

func sendString(data string, url string) error {
	stringData := schemas.StringData{Data: data}

	requestBody, err := json.Marshal(stringData)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json") // Define o tipo de conteúdo como JSON

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		return nil
	} else {
		return errors.New("string not sent")
	}
}
