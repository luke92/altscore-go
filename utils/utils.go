package utils

import (
	"altscore-go/api"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadAPIKeyAndClient loads the API key from the environment and returns a configured client.
func LoadAPIKeyAndClient() *api.Client {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Printf("Advertencia: No se pudo cargar el archivo .env: %v", err)
	}

	// Obtener la API Key
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("La API Key no está configurada. Asegúrate de configurarla en el archivo .env o como variable de entorno.")
	}

	// Crear y devolver el cliente
	return api.NewClient(apiKey)
}
