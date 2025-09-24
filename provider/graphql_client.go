package provider

import (
	"context"
	"net/http"
	"time"

	"github.com/Yamashou/gqlgenc/clientv2"
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

	// Create GraphQL client with Authorization header using RequestInterceptor
	endpoint := "https://api.runpod.io/graphql"

	// Create interceptor to add Authorization header
	authInterceptor := func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res any, next clientv2.RequestInterceptorFunc) error {
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		return next(ctx, req, gqlInfo, res)
	}

	client := NewClient(httpClient, endpoint, nil, authInterceptor)

	return &GraphQLClient{
		client: client,
	}
}

// GetClient returns the underlying GraphQL client
func (gc *GraphQLClient) GetClient() *Client {
	return gc.client
}