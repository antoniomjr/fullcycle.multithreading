package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type APIResponse struct {
	URL     string
	Content string
}

func fetchURL(ctx context.Context, url string, resultChan chan<- APIResponse) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		resultChan <- APIResponse{URL: url, Content: fmt.Sprintf("Erro: %v", err)}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		resultChan <- APIResponse{URL: url, Content: fmt.Sprintf("Erro ao ler corpo: %v", err)}
		return
	}

	resultChan <- APIResponse{URL: url, Content: string(body)}
}

func main() { //  go run main.go <cep>
	cep := os.Args[1] // Pega o CEP da linha de comando

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resultChan := make(chan APIResponse, 2)

	go fetchURL(ctx, fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep), resultChan)
	go fetchURL(ctx, fmt.Sprintf("https://viacep.com.br/ws/%s/json", cep), resultChan)

	select {
	case result := <-resultChan:
		fmt.Printf("Resposta mais rápida foi de %s: %s\n", result.URL, result.Content)
		return
	case <-ctx.Done():
		fmt.Println("Timeout")
		return
	}
}
