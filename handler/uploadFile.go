package handler

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @BasePath /

// @Summary upload file

// @Description upload a file to encrypt
// @Tags Server operations
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} file_uploaded
// @Router /upload_file [post]
func UploadFile(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Open public key file
	filePath := "./key/public_key.pem"
	filePublicKey, err := os.Open(filePath)
	handleError("Error opening public key file", err)
	defer filePublicKey.Close()

	// Reads public key file
	publicKeyData, err := ioutil.ReadAll(filePublicKey)
	handleError("Error reading public key file", err)

	blockPublicKey, _ := pem.Decode(publicKeyData)
	handleError("Error deconding public key to block", err)

	publicKeyData = blockPublicKey.Bytes

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyData)
	handleError("Error converting public key bytes to public key object", err)

	publicKeyRsa := publicKey.(*rsa.PublicKey)

	// Abra o arquivo diretamente sem salvá-lo no disco
	uploadedFile, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer uploadedFile.Close()

	data, err := ioutil.ReadAll(uploadedFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Dividir os dados em blocos menores (tamanho máximo de bloco para criptografia RSA)
	maxBlockSize := 245
	var encryptedBlocks []byte
	for len(data) > 0 {
		blockSize := len(data)
		if blockSize > maxBlockSize {
			blockSize = maxBlockSize
		}

		// Criptografar o bloco e adicionar à lista de blocos criptografados
		encryptedBlock, err := rsa.EncryptPKCS1v15(rand.Reader, publicKeyRsa, data[:blockSize])
		if err != nil {
			fmt.Println("Erro ao criptografar o bloco:", err)
			return
		}
		encryptedBlocks = append(encryptedBlocks, encryptedBlock...)
		data = data[blockSize:]
	}

	tempDir := "./encrypted_files"
	err = os.MkdirAll(tempDir, os.ModePerm)
	handleError("Error creating 'encrypted_files' directory", err)

	tempfile, err := os.Create(tempDir + "/" + file.Filename)
	handleError("Error creating file", err)
	defer tempfile.Close()

	err = ioutil.WriteFile(tempfile.Name(), encryptedBlocks, 0644)
	if err != nil {
		fmt.Println("Erro ao escrever o arquivo criptografado:", err)
		return
	}

	fmt.Println("Arquivo recebido e armazenado com sucesso.")

	ctx.JSON(http.StatusOK, gin.H{"message": "Arquivo recebido e armazenado com sucesso"})
}
