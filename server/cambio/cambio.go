package cambio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/danielzinhors/Client-Server-API-GO/server/banco"
	"github.com/danielzinhors/Client-Server-API-GO/server/model"
)

func Cotar(w http.ResponseWriter, r *http.Request) {
	moeda := "USD-BRL"
	req, err := http.Get("https://economia.awesomeapi.com.br/json/last/" + moeda)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v", err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v", err)
	}
	var cotacao model.Cotacao
	err = json.Unmarshal(res, &cotacao)
	if err != nil {
		panic(err)
	}
	banco.SalvaCotacao(&cotacao)
	json.NewEncoder(w).Encode(&cotacao.Usdbrl.Bid)
	fmt.Println(cotacao.Usdbrl.Bid)
}
