package blockchain

type Blockchain struct {
	Blocks []Block `json:"blocks"`
}

// Creamos una funcion con valores definidos para generar el primer bloque
func GenesisBlock() Block {
	genesisBlock := Block{
		Index:     0,
		Timestamp: "2025-06-17 20:00:00.0000",
		Sender:    "system",
		Message:   "This is the first block",
		PrevHash:  "0",
		Hash:      "",
	}
	genesisBlock.Hash = CalculateHash(genesisBlock)
	return genesisBlock
}

// Iniciamos la cadena usando el bloque genesis
func NewBlockchain() Blockchain {
	genesis := GenesisBlock()
	return Blockchain{Blocks: []Block{genesis}}
}

// Metodo para usar el ultimo bloque y crear con ello n nuevo bloque en la blockchain
func (bc *Blockchain) AddBlock(sender, msg string) error {
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(lastBlock, sender, msg)
	bc.Blocks = append(bc.Blocks, newBlock)
	return nil
}

// IsValid verifica la integridad de toda la blockchain
func (bc Blockchain) IsValid() bool {
	// Empezamos desde el segundo bloque (índice 1), ya que el génesis no tiene PrevHash
	for i := 1; i < len(bc.Blocks); i++ {
		current := bc.Blocks[i]
		prev := bc.Blocks[i-1]

		// Verificar que el hash del bloque actual sea correcto
		if current.Hash != CalculateHash(current) {
			return false
		}

		// Verificar que PrevHash coincida con el hash del bloque anterior
		if current.PrevHash != prev.Hash {
			return false
		}
	}
	return true
}
