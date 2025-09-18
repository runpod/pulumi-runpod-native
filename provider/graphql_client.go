package provider

import (
	"context"
	"net/http"

	gqlclient "git.sr.ht/~emersion/gqlclient"
)

// RunpodGraphQLClient wraps the generated GraphQL client with authentication
type RunpodGraphQLClient struct {
	client *gqlclient.Client
	token  string
}

// NewRunpodGraphQLClient creates a new authenticated GraphQL client
func NewRunpodGraphQLClient(token string) *RunpodGraphQLClient {
	// Create HTTP client with authentication
	httpClient := &http.Client{
		Transport: &authenticatedTransport{
			token: token,
			base:  http.DefaultTransport,
		},
	}

	// Create gqlclient with the authenticated HTTP client
	client := gqlclient.New("https://api.runpod.io/graphql", httpClient)

	return &RunpodGraphQLClient{
		client: client,
		token:  token,
	}
}

// GetClient returns the underlying gqlclient.Client for use with generated functions
func (c *RunpodGraphQLClient) GetClient() *gqlclient.Client {
	return c.client
}

// authenticatedTransport adds the API key to requests
type authenticatedTransport struct {
	token string
	base  http.RoundTripper
}

func (t *authenticatedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Add API key as query parameter (Runpod's authentication method)
	q := req.URL.Query()
	q.Add("api_key", t.token)
	req.URL.RawQuery = q.Encode()

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	return t.base.RoundTrip(req)
}

// DeployPodWrapper calls the generated DeployPod function with better ergonomics
func (c *RunpodGraphQLClient) DeployPod(ctx context.Context, req *DeployPodRequest) (*GraphQLPod, error) {
	return DeployPod(
		c.client, ctx,
		req.AiApiId, req.CloudType, req.ContainerDiskInGb, req.CountryCode,
		req.DeployCost, req.DockerArgs, req.Env, req.GpuCount, req.GpuTypeId,
		req.GpuTypeIdList, req.ImageName, req.MinDisk, req.MinDownload,
		req.MinMemoryInGb, req.MinUpload, req.MinVcpuCount, req.Name,
		req.NetworkVolumeId, req.Port, req.Ports, req.StartJupyter, req.StartSsh,
		req.StopAfter, req.SupportPublicIp, req.TemplateId, req.TerminateAfter,
		req.VolumeInGb, req.VolumeKey, req.VolumeMountPath, req.DataCenterId,
		req.SavingsPlan, req.CudaVersion,
	)
}