package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// KafkaConnectClient is a client for interacting with the Kafka Connect REST API.
type KafkaConnectClient struct {
	baseURL    string
	timeout    time.Duration
	httpClient *http.Client
}

// NewKafkaConnectClient creates a new KafkaConnectClient.
func NewKafkaConnectClient(baseURL string, timeout time.Duration) *KafkaConnectClient {
	return &KafkaConnectClient{
		baseURL:    baseURL,
		timeout:    timeout,
		httpClient: &http.Client{Timeout: timeout},
	}
}

// CreateConnnector gets a connector by name.
func (c *KafkaConnectClient) CreateConnnector(ctx context.Context, name string, config map[string]string) error {
	endpoint := fmt.Sprintf("%s/connectors", c.baseURL)

	payload, err := json.Marshal(map[string]interface{}{
		"name":   name,
		"config": config,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create connector: %s - %s", resp.Status, string(body))
	}

	return nil
}
