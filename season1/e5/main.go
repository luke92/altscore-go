package main

import (
	"altscore-go/utils"
	"fmt"
)

func main() {
	fmt.Println("La Última Defensa de la Valiant - ¡Cuenta Regresiva!")

	client := utils.LoadAPIKeyAndClient()
	if err := client.Valiant(); err != nil {
		fmt.Printf("Error executing valiant: %v\n", err)
		return
	}
}
