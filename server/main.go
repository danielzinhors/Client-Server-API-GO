package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/danielzinhors/Client-Server-API-GO/server/cambio"
	//_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println(os.Getenv("CGO_ENABLED"))
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
	//ctx := r.Context()
	cambio.Cotar()

}

// func SalvaCotacao(cotacao *Cotacao) {
// 	db, err := sql.Open("sqlite3", "mydatabase.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	if err = db.Ping(); err != nil {
// 		panic(err)
// 	}

// 	createTableSQL := `
//         CREATE TABLE IF NOT EXISTS cotacoes (
//             id INTEGER PRIMARY KEY AUTOINCREMENT,
//             moeda TEXT,
//             valor Real
//         )
//     `
// 	_, err = db.Exec(createTableSQL)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Insertar datos
// 	insertSQL := "INSERT INTO cotacoes (moeda, valor) VALUES (?, ?)"
// 	_, err = db.Exec(insertSQL, cotacao.Usdbrl.Name, cotacao.Usdbrl.Bid)
// 	if err != nil {
// 		panic(err)
// 	}
// }
