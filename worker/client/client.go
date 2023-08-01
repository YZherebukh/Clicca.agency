package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Clicca.agency/worker/domain"
	"github.com/monzo/typhon"
)

type Client struct {
	URL string
}

func (c *Client) GetTask(ctx context.Context) (*domain.TaskResponse, error) {
	resp, err := makeRequest(ctx, c.URL+"/task", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var taskResponse domain.TaskResponse

	err = resp.Decode(&taskResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response %w", err)
	}

	return &taskResponse, nil
}

func (c *Client) VerifyTask(ctx context.Context, block domain.Block) (*domain.VerifyResponse, error) {
	resp, err := makeRequest(ctx, c.URL+"/verify", http.MethodPost, block)
	if err != nil {
		return nil, err
	}

	var verifyResponse domain.VerifyResponse

	err = resp.Decode(&verifyResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response %w", err)
	}

	return &verifyResponse, nil
}

func makeRequest(ctx context.Context, url, method string, body interface{}) (typhon.Response, error) {
	resp := typhon.NewRequest(ctx, method, url, body).Send().Response()
	if resp.Error != nil {
		return resp, fmt.Errorf("request %s failed, %w", url, resp.Error)
	}
	if resp.StatusCode == 404 {
		return resp, fmt.Errorf("resource %s not found", url)
	}

	return resp, nil
}
