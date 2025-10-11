package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ollamaRequest struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
}

type ollamaResponse struct {
	Response string `json:"response"`
}

func CallOllama(ctx context.Context, model, prompt string) (string,error) {
	reqBody,err := json.Marshal(ollamaRequest{ Model: model , Prompt: prompt})

	if err != nil {
		return "", fmt.Errorf("marshal: %w", err)
	}

	ctx,cancel := context.WithTimeout(ctx,30*time.Second)
	defer cancel()

	req,err := http.NewRequestWithContext(ctx,"POST", "http://localhost:11434/api/generate", bytes.NewBuffer(reqBody))
	if err != nil{
		return "", fmt.Errorf("marshal: %w", err)
	}
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("http do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var or ollamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&or); err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	return or.Response, nil
}


// テスト用関数
func TestOllama() {
	ctx := context.Background()
	out, err := CallOllama(ctx, "llama3", "こんにちは！自己紹介してください。")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Ollama Response:", out)
}