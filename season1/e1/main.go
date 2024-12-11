package main

import (
	"altscore-go/api"
	"altscore-go/utils"
	"fmt"
)

func main() {
	fmt.Println("¡La Sonda Silenciosa! 🛰️")

	client := utils.LoadAPIKeyAndClient()

	measurement, err := client.GetMeasurement()
	if err != nil {
		fmt.Printf("Error al obtener medición: %v\n", err)
		return
	}

	orbitalSpeed, err := api.CalculateOrbitalSpeed(measurement)
	if err != nil {
		fmt.Printf("Error al calcular la velocidad de la orbita: %v\n", err)
		return
	}

	fmt.Printf("Velocidad orbital instantánea: %d km/h\n", orbitalSpeed)
}
