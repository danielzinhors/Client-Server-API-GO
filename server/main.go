package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/danielzinhors/Client-Server-API-GO/server/cambio"
	"github.com/joho/godotenv"
)

func main() {
	http.HandleFunc("/cotacao", handler)
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Some error occured. Err: %s", err)
	}
	port := os.Getenv("PORT")
	fmt.Println("Ouvindo na porta " + port)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("Request iniciada")
	defer fmt.Println("Request finalizada")
	ctx := context.Background()
	cotacao, err := cambio.Cotar(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v", err)
	}
	json.NewEncoder(w).Encode(&cotacao.Usdbrl.Bid)
	fmt.Println(cotacao.Usdbrl.Bid)
}
