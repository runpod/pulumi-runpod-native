package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

// Secret represents a RunPod secret resource
type Secret struct{}

// SecretArgs represents the inputs for creating a secret
type SecretArgs struct {
	Name        string  `pulumi:"name" json:"name"`
	Value       string  `pulumi:"value" json:"value"`
	Description *string `pulumi:"description,optional" json:"description,omitempty"`
}

// SecretState represents the state of a secret resource
type SecretState struct {
	SecretArgs
	SecretId  string `pulumi:"secretId" json:"id"`
	CreatedAt string `pulumi:"createdAt" json:"createdAt"`
	UpdatedAt string `pulumi:"updatedAt" json:"updatedAt"`
}

// Create implements the creation of a secret resource
func (s *Secret) Create(ctx p.Context, name string, input SecretArgs, preview bool) (string, SecretState, error) {
	state := SecretState{SecretArgs: input}

	if preview {
		return name, state, nil
	}

	// Get API key from config
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Token == "" {
		return "", state, fmt.Errorf("RunPod API key is required")
	}

	// Create secret via GraphQL mutation
	mutation := `
		mutation SecretCreate($input: SecretCreateInput!) {
			secretCreate(input: $input) {
				id
				name
				description
				createdAt
				updatedAt
			}
		}
	`

	inputMap := map[string]interface{}{
		"name":  input.Name,
		"value": input.Value,
	}
	if input.Description != nil {
		inputMap["description"] = *input.Description
	}

	variables := map[string]interface{}{
		"input": inputMap,
	}

	requestBody := map[string]interface{}{
		"query":     mutation,
		"variables": variables,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", state, fmt.Errorf("failed to marshal GraphQL request: %w", err)
	}

	// Make HTTP request to RunPod GraphQL API
	url := "https://zackmckenna-api.runpod.dev/graphql?api_key=" + cfg.Token
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", state, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", state, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", state, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", state, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse GraphQL response
	var gqlResponse struct {
		Data struct {
			SecretCreate struct {
				Id          string  `json:"id"`
				Name        string  `json:"name"`
				Description *string `json:"description"`
				CreatedAt   string  `json:"createdAt"`
				UpdatedAt   string  `json:"updatedAt"`
			} `json:"secretCreate"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.Unmarshal(body, &gqlResponse); err != nil {
		return "", state, fmt.Errorf("failed to parse response: %w", err)
	}

	if len(gqlResponse.Errors) > 0 {
		return "", state, fmt.Errorf("GraphQL error: %s", gqlResponse.Errors[0].Message)
	}

	// Update state with response data
	state.SecretId = gqlResponse.Data.SecretCreate.Id
	state.Description = gqlResponse.Data.SecretCreate.Description
	state.CreatedAt = gqlResponse.Data.SecretCreate.CreatedAt
	state.UpdatedAt = gqlResponse.Data.SecretCreate.UpdatedAt

	return name, state, nil
}

// Read implements reading a secret resource
func (s *Secret) Read(ctx p.Context, id string, inputs SecretArgs, state SecretState) (string, SecretArgs, SecretState, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Token == "" {
		return "", inputs, state, fmt.Errorf("RunPod API key is required")
	}

	// Query secret metadata (values are never returned for security)
	query := `
		query GetSecret($id: ID!) {
			secret(id: $id) {
				id
				name
				description
				createdAt
				updatedAt
			}
		}
	`

	variables := map[string]interface{}{
		"id": state.SecretId,
	}

	requestBody := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", inputs, state, fmt.Errorf("failed to marshal GraphQL request: %w", err)
	}

	url := "https://zackmckenna-api.runpod.dev/graphql?api_key=" + cfg.Token
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", inputs, state, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", inputs, state, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", inputs, state, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", inputs, state, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var gqlResponse struct {
		Data struct {
			Secret *struct {
				Id          string  `json:"id"`
				Name        string  `json:"name"`
				Description *string `json:"description"`
				CreatedAt   string  `json:"createdAt"`
				UpdatedAt   string  `json:"updatedAt"`
			} `json:"secret"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.Unmarshal(body, &gqlResponse); err != nil {
		return "", inputs, state, fmt.Errorf("failed to parse response: %w", err)
	}

	if len(gqlResponse.Errors) > 0 {
		return "", inputs, state, fmt.Errorf("GraphQL error: %s", gqlResponse.Errors[0].Message)
	}

	if gqlResponse.Data.Secret == nil {
		return "", inputs, state, fmt.Errorf("secret not found")
	}

	// Update state with fresh data
	state.SecretId = gqlResponse.Data.Secret.Id
	state.Name = gqlResponse.Data.Secret.Name
	state.Description = gqlResponse.Data.Secret.Description
	state.CreatedAt = gqlResponse.Data.Secret.CreatedAt
	state.UpdatedAt = gqlResponse.Data.Secret.UpdatedAt

	return id, inputs, state, nil
}

// Update implements updating a secret resource
func (s *Secret) Update(ctx p.Context, id string, olds SecretState, news SecretArgs, preview bool) (SecretState, error) {
	state := SecretState{SecretArgs: news, SecretId: olds.SecretId, CreatedAt: olds.CreatedAt}

	if preview {
		return state, nil
	}

	cfg := infer.GetConfig[Config](ctx)
	if cfg.Token == "" {
		return state, fmt.Errorf("RunPod API key is required")
	}

	// Update secret value via GraphQL mutation
	mutation := `
		mutation SecretValueUpdate($input: SecretValueUpdateInput!) {
			secretValueUpdate(input: $input) {
				id
				name
				description
				createdAt
				updatedAt
			}
		}
	`

	variables := map[string]interface{}{
		"input": map[string]interface{}{
			"id":    state.SecretId,
			"value": news.Value,
		},
	}

	requestBody := map[string]interface{}{
		"query":     mutation,
		"variables": variables,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return state, fmt.Errorf("failed to marshal GraphQL request: %w", err)
	}

	url := "https://zackmckenna-api.runpod.dev/graphql?api_key=" + cfg.Token
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return state, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return state, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return state, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return state, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var gqlResponse struct {
		Data struct {
			SecretValueUpdate struct {
				Id          string  `json:"id"`
				Name        string  `json:"name"`
				Description *string `json:"description"`
				CreatedAt   string  `json:"createdAt"`
				UpdatedAt   string  `json:"updatedAt"`
			} `json:"secretValueUpdate"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.Unmarshal(body, &gqlResponse); err != nil {
		return state, fmt.Errorf("failed to parse response: %w", err)
	}

	if len(gqlResponse.Errors) > 0 {
		return state, fmt.Errorf("GraphQL error: %s", gqlResponse.Errors[0].Message)
	}

	// Update state with response data
	state.SecretId = gqlResponse.Data.SecretValueUpdate.Id
	state.Description = gqlResponse.Data.SecretValueUpdate.Description
	state.CreatedAt = gqlResponse.Data.SecretValueUpdate.CreatedAt
	state.UpdatedAt = gqlResponse.Data.SecretValueUpdate.UpdatedAt

	return state, nil
}

// Delete implements deletion of a secret resource
func (s *Secret) Delete(ctx p.Context, id string, state SecretState) error {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Token == "" {
		return fmt.Errorf("RunPod API key is required")
	}

	// Delete secret via GraphQL mutation
	mutation := `
		mutation DeleteSecret($id: ID!) {
			secretDelete(id: $id)
		}
	`

	variables := map[string]interface{}{
		"id": state.SecretId,
	}

	requestBody := map[string]interface{}{
		"query":     mutation,
		"variables": variables,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal GraphQL request: %w", err)
	}

	url := "https://zackmckenna-api.runpod.dev/graphql?api_key=" + cfg.Token
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var gqlResponse struct {
		Data struct {
			SecretDelete interface{} `json:"secretDelete"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.Unmarshal(body, &gqlResponse); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	if len(gqlResponse.Errors) > 0 {
		return fmt.Errorf("GraphQL error: %s", gqlResponse.Errors[0].Message)
	}

	// secretDelete returns Void (null) on success, so we just check for absence of errors
	// If we get here without GraphQL errors, the delete was successful

	return nil
}