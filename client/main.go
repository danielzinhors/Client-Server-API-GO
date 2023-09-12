package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	err := godotenv.Load()
	endpoint := os.Getenv("ENDPOINT")
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	cotacao, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Erro ao criar o arquivo de resposta: %v\n", err)
	}
	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Printf("Erro ao criar o arquivo de resposta: %v\n", err)
	}
	defer file.Close()
	dolar := strings.Replace(string(cotacao), `"`, "", -1)
	_, err = file.WriteString(fmt.Sprintf("Dólar: {%s}", strings.TrimSpace(dolar)))
}
