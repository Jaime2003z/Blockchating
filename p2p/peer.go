package p2p

import (
	"blockchating/blockchain"
	"blockchating/storage"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Con esta estructura definimos los nodos de la red
type Peer struct {
	Address string // Dirección del nodo (por ejemplo, "localhost:8080")
}

// SyncFromPeer sincroniza la blockchain desde un peer
func SyncFromPeer(bc *blockchain.Blockchain, peer Peer, storagePath string) error {
	// Construir la URL del endpoint /sync del peer
	url := fmt.Sprintf("http://%s/sync", peer.Address)

	// Hacer una solicitud GET al peer
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error conectando al peer %s: %v", peer.Address, err)
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error leyendo respuesta del peer %s: %v", peer.Address, err)
	}

	// Deserializar la blockchain recibida
	var remoteBC blockchain.Blockchain
	if err := json.Unmarshal(body, &remoteBC); err != nil {
		return fmt.Errorf("error deserializando blockchain del peer %s: %v", peer.Address, err)
	}

	// Validar la blockchain recibida
	if !remoteBC.IsValid() {
		return fmt.Errorf("blockchain recibida del peer %s no es válida", peer.Address)
	}

	// Si la blockchain remota es más larga y válida, actualizar la local
	if len(remoteBC.Blocks) > len(bc.Blocks) {
		*bc = remoteBC // Sobrescribir la blockchain local
		if err := storage.SaveBlockchain(bc, storagePath); err != nil {
			return fmt.Errorf("error guardando blockchain sincronizada: %v", err)
		}
		fmt.Printf("Blockchain sincronizada desde %s, nueva longitud: %d\n", peer.Address, len(bc.Blocks))
	}

	return nil
}
