package api

import (
	"blockchating/blockchain"
	"blockchating/config"
	"github.com/gin-gonic/gin"
)

// SetupRouter configura el enrutador Gin y mapea los endpoints
func SetupRouter(bc *blockchain.Blockchain, cfg *config.Config, storagePath string) *gin.Engine {
	r := gin.Default() // Crea un enrutador Gin con middleware predeterminado

	// Definir endpoints
	r.POST("/message", PostMessage(bc, cfg, storagePath))
	r.GET("/messages", GetMessages(bc))

	return r
}
