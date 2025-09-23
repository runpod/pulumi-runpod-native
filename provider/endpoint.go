package provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type OutputDeployEndpoint struct {
	Errors []struct {
		Message   string
		Locations []struct {
			Line   int
			Column int
		}
	}
	Data struct {
		SaveEndpoint Endpoint
	}
}

type Endpoint struct {
	Id              string `pulumi:"id"`
	Name            string `pulumi:"name"`
	GpuIds          string `pulumi:"gpuIds"`
	IdleTimeout     int    `pulumi:"idleTimeout"`
	Locations       string `pulumi:"locations"`
	NetworkVolumeId string `pulumi:"networkVolumeId"`
	ScalerType      string `pulumi:"scalerType"`
	ScalerValue     int    `pulumi:"scalerValue"`
	WorkersMax      int    `pulumi:"workersMax"`
	WorkersMin      int    `pulumi:"workersMin"`
	TemplateId      string `pulumi:"templateId"`
}

type EndpointArgs struct {
	Name            string  `pulumi:"name" structs:"name,omitempty"`
	TemplateId      *string `pulumi:"templateId,optional" structs:"templateId,omitempty"`
	GpuIds          string  `pulumi:"gpuIds,optional" structs:"gpuIds,omitempty"`
	IdleTimeout     int     `pulumi:"idleTimeout,optional" structs:"idleTimeout,omitempty"`
	Locations       string  `pulumi:"locations,optional" structs:"locations,omitempty"`
	NetworkVolumeId string  `pulumi:"networkVolumeId,optional" structs:"networkVolumeId,omitempty"`
	ScalerType      string  `pulumi:"scalerType,optional" structs:"scalerType,omitempty"`
	ScalerValue     int     `pulumi:"scalerValue,optional" structs:"scalerValue,omitempty"`
	WorkersMax      int     `pulumi:"workersMax,optional" structs:"workersMax"`
	WorkersMin      int     `pulumi:"workersMin,optional" structs:"workersMin"`
}

type EndpointState struct {
	EndpointArgs
	Endpoint Endpoint `pulumi:"endpoint"`
}

func (*Endpoint) Create(ctx p.Context, name string, input EndpointArgs, preview bool) (string, EndpointState, error) {
	state := EndpointState{EndpointArgs: input}
	if preview {
		return name, state, nil
	}
	config := infer.GetConfig[Config](ctx)

	// Name is always required
	if input.Name == "" {
		return name, state, fmt.Errorf("name is required")
	}

	// For GPU endpoints, either gpuIds or templateId+instanceIds is required
	// For now, we require gpuIds to maintain compatibility
	if input.GpuIds == "" {
		return name, state, fmt.Errorf("gpuIds is required for GPU endpoints")
	}

	// Create GraphQL client using generated types
	gqlClient := NewGraphQLClient(config.Token)

	// Call the generated SaveEndpoint function
	response, err := gqlClient.GetClient().SaveEndpoint(
		context.Background(),
		input.GpuIds,                    // gpuIds (required)
		*input.TemplateId,               // templateID (required)
		input.Name,                      // name (required)
		intPtr(input.IdleTimeout),       // idleTimeout
		stringPtr(input.Locations),      // locations
		stringPtr(input.NetworkVolumeId), // networkVolumeID
		stringPtr(input.ScalerType),     // scalerType
		intPtr(input.ScalerValue),       // scalerValue
		intPtr(input.WorkersMax),        // workersMax
		intPtr(input.WorkersMin),        // workersMin
	)

	if err != nil {
		return name, state, fmt.Errorf("save endpoint mutation failed: %v", err)
	}

	// Extract the endpoint from the response
	if response.GetSaveEndpoint() == nil {
		return name, state, fmt.Errorf("no endpoint returned from save mutation")
	}

	endpointData := response.GetSaveEndpoint()

	// Convert generated type back to our Endpoint struct
	var endpointId, endpointName string
	if endpointData.GetID() != nil {
		endpointId = *endpointData.GetID()
	}
	if endpointData.GetName() != nil {
		endpointName = *endpointData.GetName()
	}

	endpoint := Endpoint{
		Id:   endpointId,
		Name: endpointName,
	}

	state.Endpoint = endpoint

	return name, state, nil
}

// Helper functions for pointer conversions - moved to pod.go to avoid duplication

func (*Endpoint) Update(ctx p.Context, id string, olds EndpointState, news EndpointArgs, preview bool) (EndpointState, error) {
	state := EndpointState{EndpointArgs: news}
	if preview {
		return state, nil
	}
	config := infer.GetConfig[Config](ctx)

	// Name is always required
	if news.Name == "" {
		return state, fmt.Errorf("name is required")
	}

	// For GPU endpoints, either gpuIds or templateId+instanceIds is required
	// For now, we require gpuIds to maintain compatibility
	if news.GpuIds == "" {
		return state, fmt.Errorf("gpuIds is required for GPU endpoints")
	}

	// Create GraphQL client using generated types
	gqlClient := NewGraphQLClient(config.Token)

	// Call the generated UpdateEndpoint function
	response, err := gqlClient.GetClient().UpdateEndpoint(
		context.Background(),
		olds.Endpoint.Id,               // id (required)
		news.GpuIds,                    // gpuIds (required)
		*news.TemplateId,               // templateID (required)
		news.Name,                      // name (required)
		intPtr(news.IdleTimeout),       // idleTimeout
		stringPtr(news.Locations),      // locations
		stringPtr(news.NetworkVolumeId), // networkVolumeID
		stringPtr(news.ScalerType),     // scalerType
		intPtr(news.ScalerValue),       // scalerValue
		intPtr(news.WorkersMax),        // workersMax
		intPtr(news.WorkersMin),        // workersMin
	)

	if err != nil {
		return state, fmt.Errorf("update endpoint mutation failed: %v", err)
	}

	// Extract the endpoint from the response
	if response.GetSaveEndpoint() == nil {
		return state, fmt.Errorf("no endpoint returned from update mutation")
	}

	endpointData := response.GetSaveEndpoint()

	// Convert generated type back to our Endpoint struct
	var endpointId, endpointName string
	if endpointData.GetID() != nil {
		endpointId = *endpointData.GetID()
	}
	if endpointData.GetName() != nil {
		endpointName = *endpointData.GetName()
	}

	endpoint := Endpoint{
		Id:   endpointId,
		Name: endpointName,
	}

	state.Endpoint = endpoint

	return state, nil
}

func compareTemplateId(a, b string) bool {
	return strings.EqualFold(a, b)
}

func (*Endpoint) Diff(ctx p.Context, id string, olds EndpointState, news EndpointArgs) (p.DiffResponse, error) {

	diff := map[string]p.PropertyDiff{}

	if !compareTemplateId(*olds.TemplateId, *news.TemplateId) {
		diff["templateId"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	if news.Name != olds.Name {
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.GpuIds != olds.GpuIds {
		diff["gpuIds"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.IdleTimeout != olds.IdleTimeout {
		diff["idleTimeout"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.Locations != olds.Locations {
		diff["locations"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.NetworkVolumeId != olds.NetworkVolumeId {
		diff["networkVolumeId"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.ScalerType != olds.ScalerType {
		diff["scalerType"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.ScalerValue != olds.ScalerValue {
		diff["scalerValue"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.WorkersMax != olds.WorkersMax {
		diff["workersMax"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.WorkersMin != olds.WorkersMin {
		diff["workersMin"] = p.PropertyDiff{Kind: p.Update}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*Endpoint) Delete(ctx p.Context, id string, props EndpointState) error {
	endpointArgs := EndpointArgs{
		TemplateId: props.TemplateId,
		Name:       props.Name,
		GpuIds:     props.GpuIds,
		WorkersMax: 0,
		WorkersMin: 0,
	}

	_, err := props.Endpoint.Update(ctx, props.Endpoint.Id, props, endpointArgs, false)

	if err != nil {
		return err
	}

	attempts := 0
	for attempts < 3 {
		err = deleteEndpoint(ctx, props.Endpoint.Id)

		if err == nil {
			break
		}

		time.Sleep(2 * time.Second)

		attempts++
	}

	if err != nil {
		return err
	}

	return nil
}

func deleteEndpoint(ctx p.Context, id string) error {
	config := infer.GetConfig[Config](ctx)

	// Create GraphQL client using generated types
	gqlClient := NewGraphQLClient(config.Token)

	// Call the generated DeleteEndpoint function
	_, err := gqlClient.GetClient().DeleteEndpoint(
		context.Background(),
		id, // id (required)
	)

	if err != nil {
		return fmt.Errorf("delete endpoint mutation failed: %v", err)
	}

	return nil
}

// func (*Endpoint) WireDependencies(f infer.FieldSelector, args *EndpointArgs, state *EndpointState) {
// 	f.OutputField(&state.TemplateId).DependsOn(f.InputField(&args.TemplateId))
// }
