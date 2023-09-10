package cambio

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/danielzinhors/Client-Server-API-GO/server/banco"
	"github.com/danielzinhors/Client-Server-API-GO/server/model"
)

func Cotar(ctx context.Context) (*model.Cotacao, error) {
	moeda := "USD-BRL"
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/"+moeda, nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseJson, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var cotacao model.Cotacao
	err = json.Unmarshal(responseJson, &cotacao)
	if err != nil {
		panic(err)
	}
	banco.SalvaCotacao(ctx, &cotacao)
	return &cotacao, nil
}
