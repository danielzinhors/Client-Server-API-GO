package banco

import (
	"context"
	"database/sql"
	"time"

	"github.com/danielzinhors/Client-Server-API-GO/server/model"
	_ "github.com/mattn/go-sqlite3"
)

func SalvaCotacao(ctx context.Context, cotacao *model.Cotacao) {
	db, err := sql.Open("sqlite3", "cambio.db")
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
            valor Real,
			creatAt TEXT
        )
    `
	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()
	// Insertar datos
	insertSQL := "INSERT INTO cotacoes (moeda, valor, creatAt) VALUES (?, ?, ?)"
	_, err = db.ExecContext(ctx, insertSQL, cotacao.Usdbrl.Name, cotacao.Usdbrl.Bid, time.DateTime)
	if err != nil {
		panic(err)
	}
}
