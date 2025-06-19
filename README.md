🧱 BlockChat – Una mini blockchain en Go

Este es un proyecto experimental donde implementé una blockchain básica en Go para entender sus conceptos fundamentales.

Está construida con:
- Go v1.24.4
- Gin para el servidor HTTP/API
- godotenv para manejar configuración personalizada vía archivo .env

📦 ¿Qué hace este proyecto?
- Simular un chat descentralizado (inspirado en blockchain), donde cada mensaje se almacena como un bloque encadenado.
- Puedes interactuar con él usando herramientas como Postman o Hoppscotch.

⚙️ Variables de configuración (Estas variables no son impositivas pero pueden definirse en un archivo .env en la raíz del proyecto):
- API_PORT=8080
- STORAGE_PATH=./data/blockchain.json
- BLOCKCHAIN_ID=blockchat-mainnet
- P2P_ADDRESS=0.0.0.0
- P2P_PORT=9000
- LOG_LEVEL=info
- MAX_MESSAGE_SIZE=1024

📡 Endpoints disponibles
- GET	/messages	Devuelve todos los mensajes del chat
- POST	/message	Envía un nuevo mensaje al blockchain

📝 Ejemplo de cuerpo para el POST /message:
{
  "sender": "user",
  "message": "hola mundo"
}

📂 Estructura del proyecto

blockchat/
│
├── main.go                  → Punto de entrada de la aplicación
├── go.mod                   → Módulo de dependencias de Go
│
├── config/                  → Carga de configuración desde .env
│   └── config.go
│
├── blockchain/              → Lógica central de blockchain
│   ├── block.go             → Definición de bloques y hash
│   └── chain.go             → Manejo de la cadena (agregar, validar)
│
├── storage/                 → Persistencia en disco (JSON)
│   └── json_store.go
│
├── api/                     → API HTTP con Gin
│   ├── handlers.go          → Rutas /message y /messages
│   └── router.go            → Inicialización del servidor
│
├── p2p/                     → (Futuro) Comunicación entre nodos
│   └── peer.go              → Lógica de red y sincronización
│
└── utils/                   → Funciones auxiliares (tiempo, hashing, etc.)
    └── helpers.go

Estado actual
✅ Blockchain local funcional
✅ API REST para enviar y consultar mensajes
🔜 Integración P2P (próximamente)
🔜 Pruebas automáticas y validación extendida

📥 Cómo usar
- Clona el repositorio
- Crea un archivo .env con tu configuración
- Ejecuta el proyecto: "go run main.go"
  
Puedes probarlo desde Postman, Hoppscotch o cURL.
