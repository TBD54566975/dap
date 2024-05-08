package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/tbd54566975/web5-go/dids/didweb"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("provide a domain name as an argument")
	}

	domain := os.Args[1]
	registryURL, err := url.Parse(domain + "/dap-registry")
	if err != nil {
		log.Fatalf("failed to parse domain into registry service endpoint url: %v", err)
	}

	if registryURL.Scheme == "" {
		registryURL.Scheme = "https"
	}

	bearerDID, err := didweb.Create(domain, didweb.Service(
		"dap-registry",
		"dap-registry",
		registryURL.String()),
	)

	if err != nil {
		log.Fatalf("failed to generate did:web %v", err.Error())
	}

	portableDID, err := bearerDID.ToPortableDID()
	if err != nil {
		log.Fatalf("failed to export portable did: %v", err.Error())
	}

	dd, err := json.MarshalIndent(portableDID.Document, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal did document: %v", err)
	}

	filePath := "./did-document.json"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("failed to open/create file to store portable did: %v", err)
	}

	defer file.Close()

	_, err = file.Write(dd)
	if err != nil {
		log.Fatalf("failed to write portable did to %s: %v", filePath, err)
	}

	log.Printf("created did:web and wrote did document to %s! This will be read by the registry when it starts\n", filePath)

	pd, err := json.Marshal(portableDID)
	if err != nil {
		log.Fatalf("failed to marshal portable did: %v", err)
	}

	fmt.Println("store this somewhere safe. it will need to be loaded by the registry")
	fmt.Println(string(pd))
}
