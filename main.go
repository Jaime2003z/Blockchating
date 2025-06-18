package main

import (
	"blockchating/api"
	"blockchating/config"
	"blockchating/storage"
	"fmt"
)

func main() {
	// Cargar configuración
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error cargando configuración: %v\n", err)
		return
	}

	// Cargar blockchain
	bc, err := storage.LoadBlockchain(cfg.StoragePath)
	if err != nil {
		fmt.Printf("Error cargando blockchain: %v\n", err)
		return
	}

	// Validar blockchain
	if !bc.IsValid() {
		fmt.Println("Error: La blockchain cargada no es válida")
		return
	}

	// Configurar el enrutador Gin
	router := api.SetupRouter(bc, cfg, cfg.StoragePath)

	// Iniciar el servidor
	addr := ":" + cfg.APIPort
	err = router.Run(addr)
	if err != nil {
		fmt.Printf("Error iniciando el servidor: %v\n", err)
		return
	}
}
