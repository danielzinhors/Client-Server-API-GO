package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danielzinhors/Client-Server-API-GO/server/cambio"
	"github.com/joho/godotenv"
)

func main() {
	http.HandleFunc("/cotacao", handler)
	err := godotenv.Load()
	if err != nil {
		log.Printf("Some error occured. Err: %s", err)
	}
	port := os.Getenv("PORT")
	log.Println("Ouvindo na porta " + port)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("Request iniciada")
	defer log.Println("Request finalizada")
	ctx := context.Background()
	cotacao, err := cambio.Cotar(ctx)
	if err != nil {
		log.Printf("Error %v", err)
	}
	json.NewEncoder(w).Encode(&cotacao.Usdbrl.Bid)
}
