package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/tbd54566975/web5-go/dids/didweb"
)

func main() {
	bearerDID, err := didweb.Create("http://localhost:8080", didweb.Service(
		"dap-registry",
		"dap-registry",
		"http://localhost:8080"),
	)

	if err != nil {
		log.Fatalf("failed to generate did:web %v", err.Error())
	}

	portableDID, err := bearerDID.ToPortableDID()
	if err != nil {
		log.Fatalf("failed to export portable did: %v", err.Error())
	}
	marshaled, err := json.MarshalIndent(portableDID, "", "  ")

	if err != nil {
		log.Fatalf("failed to marshal portable did: %v", err.Error())
	}

	dirPath := "priv"
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		log.Fatalf("failed to create %s dir: %v", dirPath, err)
	}

	filePath := dirPath + "/portable-did.json"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open/create file to store portable did: %v")
	}

	defer file.Close()

	_, err = file.Write(marshaled)
	if err != nil {
		log.Fatalf("failed to write portable did to %s: %v", filePath, err)
	}

	log.Printf("created did:web and exported portable did to %s!", filePath)
}
