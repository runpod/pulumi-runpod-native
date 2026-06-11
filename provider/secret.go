package provider

import (
	"encoding/json"
	"fmt"

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

	request := GraphQLRequest{
		Query:     mutation,
		Variables: variables,
	}

	resp, err := MakeGraphQLRequest(ctx, cfg.Token, request)
	if err != nil {
		return "", state, err
	}

	// Parse GraphQL response
	var secretCreateData struct {
		SecretCreate struct {
			Id          string  `json:"id"`
			Name        string  `json:"name"`
			Description *string `json:"description"`
			CreatedAt   string  `json:"createdAt"`
			UpdatedAt   string  `json:"updatedAt"`
		} `json:"secretCreate"`
	}

	if err := json.Unmarshal(resp.Data, &secretCreateData); err != nil {
		return "", state, fmt.Errorf("failed to parse response: %w", err)
	}

	// Update state with response data
	state.SecretId = secretCreateData.SecretCreate.Id
	state.Description = secretCreateData.SecretCreate.Description
	state.CreatedAt = secretCreateData.SecretCreate.CreatedAt
	state.UpdatedAt = secretCreateData.SecretCreate.UpdatedAt

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

	request := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	resp, err := MakeGraphQLRequest(ctx, cfg.Token, request)
	if err != nil {
		return "", inputs, state, err
	}

	var secretData struct {
		Secret *struct {
			Id          string  `json:"id"`
			Name        string  `json:"name"`
			Description *string `json:"description"`
			CreatedAt   string  `json:"createdAt"`
			UpdatedAt   string  `json:"updatedAt"`
		} `json:"secret"`
	}

	if err := json.Unmarshal(resp.Data, &secretData); err != nil {
		return "", inputs, state, fmt.Errorf("failed to parse response: %w", err)
	}

	if secretData.Secret == nil {
		// Secret not found - return empty ID to signal deletion to Pulumi
		return id, inputs, SecretState{}, nil
	}

	// Update state with fresh data
	state.SecretId = secretData.Secret.Id
	state.Name = secretData.Secret.Name
	state.Description = secretData.Secret.Description
	state.CreatedAt = secretData.Secret.CreatedAt
	state.UpdatedAt = secretData.Secret.UpdatedAt

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

	// Check if name changed - this requires recreating the resource
	if news.Name != olds.Name {
		return state, fmt.Errorf("secret name cannot be changed - this requires recreating the resource")
	}

	// Update secret value if changed
	if news.Value != olds.Value {
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

		request := GraphQLRequest{
			Query:     mutation,
			Variables: variables,
		}

		resp, err := MakeGraphQLRequest(ctx, cfg.Token, request)
		if err != nil {
			return state, err
		}

		var updateData struct {
			SecretValueUpdate struct {
				Id          string  `json:"id"`
				Name        string  `json:"name"`
				Description *string `json:"description"`
				CreatedAt   string  `json:"createdAt"`
				UpdatedAt   string  `json:"updatedAt"`
			} `json:"secretValueUpdate"`
		}

		if err := json.Unmarshal(resp.Data, &updateData); err != nil {
			return state, fmt.Errorf("failed to parse response: %w", err)
		}

		// Update state with response data
		state.SecretId = updateData.SecretValueUpdate.Id
		state.Description = updateData.SecretValueUpdate.Description
		state.CreatedAt = updateData.SecretValueUpdate.CreatedAt
		state.UpdatedAt = updateData.SecretValueUpdate.UpdatedAt
	}

	// Update description if changed
	if (news.Description == nil && olds.Description != nil) ||
	   (news.Description != nil && olds.Description == nil) ||
	   (news.Description != nil && olds.Description != nil && *news.Description != *olds.Description) {

		if news.Description != nil {
			// Update description
			mutation := `
				mutation SecretDescriptionUpdate($input: SecretDescriptionUpdateInput!) {
					secretDescriptionUpdate(input: $input) {
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
					"id":          state.SecretId,
					"description": *news.Description,
				},
			}

			request := GraphQLRequest{
				Query:     mutation,
				Variables: variables,
			}

			resp, err := MakeGraphQLRequest(ctx, cfg.Token, request)
			if err != nil {
				return state, err
			}

			var updateData struct {
				SecretDescriptionUpdate struct {
					Id          string  `json:"id"`
					Name        string  `json:"name"`
					Description *string `json:"description"`
					CreatedAt   string  `json:"createdAt"`
					UpdatedAt   string  `json:"updatedAt"`
				} `json:"secretDescriptionUpdate"`
			}

			if err := json.Unmarshal(resp.Data, &updateData); err != nil {
				return state, fmt.Errorf("failed to parse response: %w", err)
			}

			// Update state with response data
			state.SecretId = updateData.SecretDescriptionUpdate.Id
			state.Description = updateData.SecretDescriptionUpdate.Description
			state.CreatedAt = updateData.SecretDescriptionUpdate.CreatedAt
			state.UpdatedAt = updateData.SecretDescriptionUpdate.UpdatedAt
		} else {
			// Note: RunPod API doesn't appear to support removing descriptions
			// For now, we'll keep the existing description
			state.Description = olds.Description
		}
	}

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

	request := GraphQLRequest{
		Query:     mutation,
		Variables: variables,
	}

	_, err := MakeGraphQLRequest(ctx, cfg.Token, request)
	if err != nil {
		return err
	}

	// secretDelete returns Void (null) on success, so we just check for absence of errors
	// If we get here without GraphQL errors, the delete was successful

	return nil
}