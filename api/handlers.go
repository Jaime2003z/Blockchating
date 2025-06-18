package api

import (
	"blockchating/blockchain"
	"blockchating/config"
	"blockchating/storage"
	"blockchating/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostMessage maneja la solicitud POST /message para agregar un nuevo mensaje a la blockchain
func PostMessage(bc *blockchain.Blockchain, cfg *config.Config, storagePath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Estructura para el cuerpo de la solicitud
		type MessageRequest struct {
			Sender  string `json:"sender" binding:"required"`
			Message string `json:"message" binding:"required"`
		}

		var req MessageRequest
		// Deserializar el JSON de la solicitud
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Solicitud inválida, se requiere 'sender' y 'message'"})
			return
		}

		if !utils.ValidateLength(req.Message, cfg.MaxMessageSize) { // Usamos el limite de la configuracion
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Mensaje excede el límite de %d caracteres", cfg.MaxMessageSize)})
			return
		}

		// Agregar el bloque a la blockchain
		if err := bc.AddBlock(req.Sender, req.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo agregar el bloque: " + err.Error()})
			return
		}

		// Guardar la blockchain
		if err := storage.SaveBlockchain(bc, storagePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar la blockchain: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Mensaje agregado a la blockchain"})
	}
}

// GetMessages maneja la solicitud GET /messages para obtener todos los mensajes
func GetMessages(bc *blockchain.Blockchain) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"blocks": bc.Blocks})
	}
}
