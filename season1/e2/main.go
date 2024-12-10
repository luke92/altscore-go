package main

import (
	"altscore-go/api"
	"altscore-go/utils"
	"fmt"
)

func main() {
	fmt.Println("¡El Enigma Cósmico de Kepler-452b! 🌌")

	client := utils.LoadAPIKeyAndClient()

	stars, err := client.GetStars()
	if err != nil {
		fmt.Printf("Error al obtener estrellas: %v\n", err)
		return
	}

	fmt.Printf("Cantidad de estrellas: %d\n", len(stars))

	averageResonance := api.CalculateAverageResonance(stars)
	fmt.Printf("Resonancia promedio de las estrellas: %f\n", averageResonance)
}
