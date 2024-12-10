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

	orbitalSpeed := api.CalculateOrbitalSpeed(measurement)
	fmt.Printf("Velocidad orbital instantánea: %d km/h\n", orbitalSpeed)
}
