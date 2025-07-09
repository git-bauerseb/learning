package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	ID string `json:"product_id"`
}

func main() {

	prod := Product{
		ID: "ABC-123",
	}

	jsonData, err := json.MarshalIndent(prod, "", "  ")
	if err != nil {
		log.Fatalf("Error")
	}

	fmt.Printf("%+v\n", string(jsonData))
}
