package provider

import (
	"net/http"
	"time"
)

// GraphQLClient wraps the generated GraphQL client
type GraphQLClient struct {
	client *Client
}

// NewGraphQLClient creates a new GraphQL client for RunPod API
func NewGraphQLClient(token string) *GraphQLClient {
	// Create HTTP client with timeout
	httpClient := &http.Client{
		Timeout: time.Second * 30,
	}

	// Create GraphQL client
	endpoint := "https://api.runpod.dev/graphql?api_key=" + token

	client := NewClient(httpClient, endpoint, nil)

	return &GraphQLClient{
		client: client,
	}
}

// GetClient returns the underlying GraphQL client
func (gc *GraphQLClient) GetClient() *Client {
	return gc.client
}