package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iamcaye/aitor-cli/utils"
)

func NewClient() *http.Client {
	return &http.Client{}
}

type AuditDto struct {
	Content string `json:"content"`
}

func SendAuditRequest(data []byte) (*http.Response, error) {
	client := NewClient()
	reqBody := AuditDto{
		Content: string(data),
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	// Compress the request body using Brotli
	compressedBody, err := utils.CompressContent(reqBodyBytes)
	if err != nil {
		return nil, err
	}
	// print the difference between the compressed and uncompressed body
	fmt.Printf("Uncompressed size: %d bytes\n", len(reqBodyBytes))
	fmt.Printf("Compressed size: %d bytes\n", len(compressedBody))
	fmt.Printf("Compression ratio: %.2f%%\n", float64(len(compressedBody))/float64(len(reqBodyBytes))*100)

	req, err := http.NewRequest("POST", "http://localhost:3000/api/audit", bytes.NewBuffer(compressedBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Encoding", "br")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println("Response Status:", resp.Status)
	return resp, nil
}
