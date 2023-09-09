package banco

import (
	"database/sql"

	"github.com/danielzinhors/Client-Server-API-GO/server/model"
	_ "github.com/mattn/go-sqlite3"
)

func SalvaCotacao(cotacao *model.Cotacao) {
	db, err := sql.Open("sqlite3", "mydatabase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		panic(err)
	}

	createTableSQL := `
        CREATE TABLE IF NOT EXISTS cotacoes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            moeda TEXT,
            valor Real
        )
    `
	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}

	// Insertar datos
	insertSQL := "INSERT INTO cotacoes (moeda, valor) VALUES (?, ?)"
	_, err = db.Exec(insertSQL, cotacao.Usdbrl.Name, cotacao.Usdbrl.Bid)
	if err != nil {
		panic(err)
	}
}
