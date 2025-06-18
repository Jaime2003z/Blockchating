package storage

import (
	"blockchating/blockchain"
	"encoding/json"
	"os"
	"path/filepath"
)

// Metodo para guardar la blockchain en un archivo JSON
func SaveBlockchain(chain *blockchain.Blockchain, path string) error {
	// Asegurar que el directorio existe
	dir := filepath.Dir(path) // Extrae el directorio de la ruta (e.g., "./data")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err // Retorna error si no se puede crear el directorio
	}

	// Convertir la blockchain a JSON
	data, err := json.MarshalIndent(chain, "", "  ")
	if err != nil {
		return err
	}

	// Escribir el JSON en el archivo
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// LoadBlockchain carga la blockchain desde un archivo JSON
func LoadBlockchain(path string) (*blockchain.Blockchain, error) {
	// Leer el archivo
	data, err := os.ReadFile(path)
	if err != nil {
		// Si el archivo no existe, retornar una nueva blockchain
		if os.IsNotExist(err) {
			return &blockchain.Blockchain{Blocks: []blockchain.Block{blockchain.GenesisBlock()}}, nil
		}
		return nil, err
	}

	// Convertir el JSON a una estructura Blockchain
	var chain blockchain.Blockchain
	err = json.Unmarshal(data, &chain)
	if err != nil {
		return nil, err
	}

	return &chain, nil
}
