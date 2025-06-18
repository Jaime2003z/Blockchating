package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Creamos una estructura con la cual van a trabajar los bloques
type Block struct {
	Index     int    `json:"index"`     //indice indica el numero de este bloque
	Timestamp string `json:"timestamp"` //La hora a la que se creo
	Sender    string `json:"sender"`    //Nos informa quien envió el mensaje
	Message   string `json:"message"`   //El contenido del mensaje
	PrevHash  string `json:"prevHash"`  //Nos indica que bloque lo firmó antes
	Hash      string `json:"hash"`      //Nos indica el hash propio
}

// Creamos una función para firmar los bloques usando el hash
func CalculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Sender + block.Message + block.PrevHash //Concatenamos toda la info del bloque
	h := sha256.New()                                                                                     //Instanciamos sha256
	h.Write([]byte(record))                                                                               //volvemos lel texto a bytes
	hashed := h.Sum(nil)                                                                                  //finalizamos el calculo del hash como una secuencia de bytes
	return hex.EncodeToString(hashed)                                                                     //volvemos el chat una cadena hexadecimal
}

func NewBlock(prev Block, sender string, msg string) Block {
	newB := Block{
		Index:     prev.Index + 1,
		Timestamp: time.Now().Format(time.RFC3339),
		Sender:    sender,
		Message:   msg,
		PrevHash:  prev.Hash,
	}
	newB.Hash = CalculateHash(newB)
	return newB
}
