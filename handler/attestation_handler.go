package handler

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/TPM-Project-Larces/back-end.git/config"
	"github.com/TPM-Project-Larces/back-end.git/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"os"
	"time"
)

// @BasePath /
// @Summary Upload challenge
// @Description Upload a challenge to TPM
// @Tags Attestation
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Challenge"
// @Success 200 {string} string "challenge_uploaded"
// @Failure 400 {string} string "bad_request"
// @Failure 500 {string} string "internal_server_error"
// @Router /attestation/upload_challenge [post]
func UploadChallenge(ctx *gin.Context) {
	username, err := MiddlewaveVerifyToken(ctx)
	if err != nil || username == "" {
		response(ctx, 403, "invalid_token", err)
		return
	}
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	tempDir := "./challenge"
	err = os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	tempFile, err := os.Create(tempDir + "/" + file.Filename)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}
	defer tempFile.Close()

	err = ctx.SaveUploadedFile(file, "./challenge/"+file.Filename)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	ChallengeBytes, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	challengeCollection := config.GetMongoDB().Collection("challenge")

	// Checks if the user has a challenge in database
	existingSig := model.PublicKey{}
	err = challengeCollection.FindOne(context.Background(),
		bson.M{"username": username}).Decode(&existingSig)
	if err == nil {
		// If the user already has a challenge, delete the previous challenge
		_, err := challengeCollection.DeleteOne(context.Background(), bson.M{"username": username})
		if err != nil {
			response(ctx, 500, "internal_server_error", err)
			return
		}
	}

	challenge := model.Challenge{
		Username:  username,
		Data:      ChallengeBytes,
		CreatedAt: time.Now(),
	}

	_, err = challengeCollection.InsertOne(context.Background(), challenge)
	if err != nil {
		response(ctx, 500, "challenge_not_regenerated", err)
	}

	err = os.Remove(tempFile.Name())
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	response(ctx, 200, "challenge_created", err)
}

// @BasePath /
// @Summary Upload signature
// @Description Upload a signature from TPM
// @Tags Attestation
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Signature"
// @Success 200 {string} string "Signature_uploaded"
// @Failure 400 {string} string "bad_request"
// @Failure 500 {string} string "internal_server_error"
// @Router /attestation/upload_signature [post]
func UploadSignature(ctx *gin.Context) {
	username, err := MiddlewaveVerifyToken(ctx)
	if err != nil || username == "" {
		response(ctx, 403, "invalid_token", err)
		return
	}
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 400, "bad_request", nil)
		return
	}

	tempDir := "./Signature"
	err = os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	tempFile, err := os.Create(tempDir + "/" + file.Filename)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}
	defer tempFile.Close()

	err = ctx.SaveUploadedFile(file, tempFile.Name())
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	SignatureBytes, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	signatureCollection := config.GetMongoDB().Collection("signature")

	// Checks if the user has a signature in database
	existingSig := model.PublicKey{}
	err = signatureCollection.FindOne(context.Background(),
		bson.M{"username": username}).Decode(&existingSig)
	if err == nil {
		// If the user already has a signature, delete the previous signature
		_, err := signatureCollection.DeleteOne(context.Background(), bson.M{"username": username})
		if err != nil {
			response(ctx, 500, "internal_server_error", err)
			return
		}
	}

	signature := model.Signature{
		Username:  username,
		Data:      SignatureBytes,
		CreatedAt: time.Now(),
	}

	_, err = signatureCollection.InsertOne(context.Background(), signature)
	if err != nil {
		response(ctx, 500, "signature_not_regenerated", err)
	}

	err = os.Remove(tempFile.Name())
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	response(ctx, 200, "signature_created", err)
}

// @BasePath /
// @Summary Upload attestation key
// @Description Uploads a public attestation key
// @Tags Attestation
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} string "attestation_key_uploaded"
// @Failure 400 {string} string "bad_request"
// @Failure 500 {string} string "internal_server_error"
// @Router /attestation/upload_attestation_key [post]
func UploadAttestationKey(ctx *gin.Context) {
	username, err := MiddlewaveVerifyToken(ctx)
	if err != nil || username == "" {
		response(ctx, 403, "invalid_token", err)
		return
	}
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 400, "bad_request", nil)
		return
	}

	tempDir := "./key"
	err = os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	tempFile, err := os.Create(tempDir + "/" + file.Filename)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}
	defer tempFile.Close()

	err = ctx.SaveUploadedFile(file, tempFile.Name())
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	publicKeyPEM, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	attestationCollection := config.GetMongoDB().Collection("attestationkey")

	// Checks if the user has an attestation key in database
	existingKey := model.AttestationKey{}
	err = attestationCollection.FindOne(context.Background(),
		bson.M{"username": username}).Decode(&existingKey)
	if err == nil {
		// If the user already has an attestation key, delete the previous key
		_, err := attestationCollection.DeleteOne(context.Background(), bson.M{"username": username})
		if err != nil {
			response(ctx, 500, "internal_server_error", err)
			return
		}
	}

	key := model.AttestationKey{
		Username:  username,
		KeyBytes:  publicKeyPEM,
		CreatedAt: time.Now(),
	}

	_, err = attestationCollection.InsertOne(context.Background(), key)
	if err != nil {
		response(ctx, 500, "attestation_key_not_created", err)
		return
	}

	ctx.JSON(200, gin.H{"message": "attestation_key_uploaded", "username": username})
}

// @BasePath /
// @Summary Confirm tpm attestation
// @Description Confirm tpm attestation
// @Tags Attestation
// @Accept json
// @Produce json
// @Success 200 {string} string "attestation_successfully"
// @Failure 500 {string} string "internal_server_rror"
// @Router /attestation/make_attestation [post]
func UploadAttestation(ctx *gin.Context) error {
	username, err := MiddlewaveVerifyToken(ctx)
	if err != nil || username == "" {
		response(ctx, 403, "invalid_token", err)
		return err
	}
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return err
	}

	keyCollection := config.GetMongoDB().Collection("attestationkey")

	// Checks if the user has a key in database
	existingKey := model.AttestationKey{}
	err = keyCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&existingKey)

	blockPublicKey, _ := pem.Decode(existingKey.KeyBytes)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return err
	}

	publicKey, err := x509.ParsePKIXPublicKey(blockPublicKey.Bytes)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return err
	}

	publicKeyRsa := publicKey.(*rsa.PublicKey)

	// Open the file directly without saving it to disk
	uploadedFile, err := file.Open()
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return err
	}
	defer uploadedFile.Close()

	//Find User Signature
	signatureCollection := config.GetMongoDB().Collection("signature")
	Signature := model.Signature{}
	err = signatureCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&Signature)

	//Find User Challenge
	challengeCollection := config.GetMongoDB().Collection("challenge")
	Challenge := model.Signature{}
	err = challengeCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&Challenge)

	//Make Attestation
	err = rsa.VerifyPKCS1v15(publicKeyRsa, crypto.SHA256, Challenge.Data, Signature.Data)
	if err != nil {
		response(ctx, 500, "attestation_failed", nil)
		return err
	}

	response(ctx, 200, "attestation_successfully", nil)
	return nil
}
