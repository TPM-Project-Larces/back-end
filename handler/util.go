package handler

import "fmt"

func handleError(message string, err error) {
	if err != nil {
		fmt.Println(message+":", err)
		panic((err))
	}
}
