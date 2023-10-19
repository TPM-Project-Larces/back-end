package main

import (
	"github.com/TPM-Project-Larces/back-end.git/handler"
	"github.com/TPM-Project-Larces/back-end.git/router"
)

func main() {
	handler.Init()
	// Initialize Router
	router.Initialize()
}
