package main

import (
	"github.com/TPM-Project-Larces/back-end.git/config"
	"github.com/TPM-Project-Larces/back-end.git/router"
)

// @title Server API
// @description Server Operations
// @version 1.0.0
//
//	@contact {
//	  name: "Computer Networks and Security Laboratory (LARCES)",
//	  url: "https://larces.uece.br/",
//	  email: "larces@uece.br
//	}
//
// @BasePath /
func main() {
	// Initailize Database
	config.DatabaseInitialize()

	// Initailize Routes
	router.RouterInitialize()
}
