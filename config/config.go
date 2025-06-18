package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config define la estructura para los parámetros de configuración
type Config struct {
	APIPort        string // Puerto para el servidor API
	StoragePath    string // Ruta donde se guarda la blockchain
	BlockchainID   string // Identificador único de la blockchain
	P2PAddress     string // Dirección para el nodo P2P
	P2PPort        string // Puerto para el nodo P2P
	LogLevel       string // Nivel de verbosidad para los logs
	MaxMessageSize int    // Tamaño máximo del mensaje en bytes
}

// LoadConfig carga la configuración desde variables de entorno o valores por defecto
func LoadConfig() (*Config, error) {
	// Cargar archivo .env si existe
	if err := godotenv.Load(); err != nil {
		fmt.Println("No se encontró archivo .env, usando valores por defecto")
	}

	config := &Config{
		APIPort:        getEnv("API_PORT", "8080"),
		StoragePath:    getEnv("STORAGE_PATH", "./storage/data/blockchain.json"),
		BlockchainID:   getEnv("BLOCKCHAIN_ID", "blockchat-mainnet"),
		P2PAddress:     getEnv("P2P_ADDRESS", "0.0.0.0"),
		P2PPort:        getEnv("P2P_PORT", "9000"),
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		MaxMessageSize: getEnvAsInt("MAX_MESSAGE_SIZE", 1024), // 1KB por defecto
	}

	return config, nil
}

// getEnv obtiene el valor de una variable de entorno o un valor por defecto
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt obtiene el valor de una variable de entorno como entero o un valor por defecto
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
