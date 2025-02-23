package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"todo-list/pkg/notifyService"
)

type HTTPNotifyClient struct {
	URLAddress string
	client     http.Client
}

func NewNotifyClient(url string) *HTTPNotifyClient {
	return &HTTPNotifyClient{URLAddress: url, client: http.Client{}}
}

func (n *HTTPNotifyClient) PushToNotifyService(users []notifyService.TaskDeadlineInfo) error {
	data, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("failed marshal users data: %v", err)
	}

	if err := n.doRequest(data); err != nil {
		return fmt.Errorf("failed do request: %v", err)
	}

	return nil
}

func (n *HTTPNotifyClient) doRequest(data []byte) error {
	req, err := http.NewRequest(http.MethodPost, n.URLAddress, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := n.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return fmt.Errorf("http status: %v failed to read response body: %v", resp.StatusCode, readErr)
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("notification service error: %s", string(bodyBytes))
		return fmt.Errorf("http status: %v, response: %s", resp.StatusCode, string(bodyBytes))
	}

	logrus.Println("notify service push successfully")
	return nil
}
