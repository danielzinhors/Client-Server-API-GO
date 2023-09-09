package cambio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Cotacao struct {
	Usdbrl Usdbrl `json:"USDBRL"`
}

type Usdbrl struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

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
	var cotacao Cotacao
	err = json.Unmarshal(res, &cotacao)
	if err != nil {
		panic(err)
	}
	//	SalvaCotacao(&cotacao)
	json.NewEncoder(w).Encode(&cotacao.Usdbrl.Bid)
	fmt.Println(cotacao.Usdbrl.Bid)
}
