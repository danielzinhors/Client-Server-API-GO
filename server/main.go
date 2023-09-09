package main

import (
	"fmt"
	"net/http"

	"github.com/danielzinhors/Client-Server-API-GO/server/cambio"
)

func main() {
	http.HandleFunc("/cotacao", handler)
	fmt.Println("Ouvindo na porta 8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("Request iniciada")
	defer fmt.Println("Request finalizada")
	cambio.Cotar(w, r)
}
