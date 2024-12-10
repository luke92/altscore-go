package api

import (
	"fmt"
)

func HandleRequestError(err error) {
	if err != nil {
		fmt.Println("Error en la solicitud API:", err)
	}
}
