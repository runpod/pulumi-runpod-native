package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/fatih/structs"
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
	GpuIds          string  `pulumi:"gpuIds" structs:"gpuIds"`
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

	if input.Name == "" || input.GpuIds == "" || input.TemplateId == nil {
		return name, state, fmt.Errorf("TemplateId, gpuIds and name are required")
	}

	gqlVariable := structs.Map(input)

	gqlInput := GqlInput{
		Query: `
		mutation SaveEndpoint (
	    $gpuIds: String!
	    $templateId: String!
	    $name: String!
	    $idleTimeout: Int 
	    $locations: String
	    $networkVolumeId: String
	    $scalerType: String
	    $scalerValue: Int
	    $workersMax: Int
	    $workersMin: Int 
		) {
		saveEndpoint(input: {
	    gpuIds: $gpuIds,
	    templateId: $templateId,
	    name: $name,
	    idleTimeout: $idleTimeout,
	    locations: $locations,
	    networkVolumeId: $networkVolumeId,
	    scalerType: $scalerType,
	    scalerValue: $scalerValue,
	    workersMax: $workersMax,
	    workersMin: $workersMin 
		}) {
			id
			name
		}
	}`,
		Variables: gqlVariable,
	}

	jsonValue, err := json.Marshal(gqlInput)

	if err != nil {
		return name, state, err
	}

	url := URL + config.Token

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	b, _ := httputil.DumpRequest(req, true)

	fmt.Println(string(b))

	if err != nil {
		return name, state, err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 20}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return name, state, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return name, state, err
	}

	output := &OutputDeployEndpoint{}
	err = json.Unmarshal(data, output)
	if err != nil {
		return name, state, err
	}

	if len(output.Errors) > 0 {
		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
		return name, state, err
	}

	endpoint := output.Data.SaveEndpoint
	if endpoint.Id == "" {
		err = fmt.Errorf("graphql endpoint is nil: %s", string(data))
		return name, state, err
	}

	state.Endpoint = endpoint

	return name, state, nil
}

func (*Endpoint) Update(ctx p.Context, id string, olds EndpointState, news EndpointArgs, preview bool) (EndpointState, error) {
	state := EndpointState{EndpointArgs: news}
	if preview {
		return state, nil
	}
	config := infer.GetConfig[Config](ctx)

	if news.Name == "" || news.GpuIds == "" || news.TemplateId == nil {
		return state, fmt.Errorf("templateId, gpuIds and name are required")
	}

	gqlVariable := structs.Map(news)
	gqlVariable["id"] = olds.Endpoint.Id

	gqlInput := GqlInput{
		Query: `
	mutation SaveEndpoint (
	    $id: String!
	    $gpuIds: String!
	    $templateId: String!
	    $name: String!
	    $idleTimeout: Int 
	    $locations: String
	    $networkVolumeId: String
	    $scalerType: String
	    $scalerValue: Int
	    $workersMax: Int
	    $workersMin: Int 
		) {
		saveEndpoint(input: {
	    id: $id,
	    gpuIds: $gpuIds,
	    templateId: $templateId,
	    name: $name,
	    idleTimeout: $idleTimeout,
	    locations: $locations,
	    networkVolumeId: $networkVolumeId,
	    scalerType: $scalerType,
	    scalerValue: $scalerValue,
	    workersMax: $workersMax,
	    workersMin: $workersMin 
		}) {
			id
			name
		}
	}`,
		Variables: gqlVariable,
	}

	jsonValue, err := json.Marshal(gqlInput)
	if err != nil {
		return state, err
	}

	url := URL + config.Token

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return state, err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 20}

	resp, err := client.Do(req)
	if err != nil {
		return state, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return state, err
	}

	output := &OutputDeployEndpoint{}
	err = json.Unmarshal(data, output)
	if err != nil {
		return state, err
	}

	if len(output.Errors) > 0 {
		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
		return state, err
	}

	endpoint := output.Data.SaveEndpoint
	if endpoint.Id == "" {
		err = fmt.Errorf("graphql endpoint is nil: %s", string(data))
		return state, err
	}

	state.Endpoint = endpoint

	return state, nil
}

func (*Endpoint) Diff(ctx p.Context, id string, olds EndpointState, news EndpointArgs) (p.DiffResponse, error) {

	diff := map[string]p.PropertyDiff{}

	if news.TemplateId != olds.TemplateId {
		// Template ID is not updatable
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
	for attempts < 10 {
		err = deleteEndpoint(ctx, props.Endpoint.Id)

		if err == nil {
			break
		}

		time.Sleep(2 * time.Second)

		attempts++
	}

	if err != nil {
		println("delete endpoint failed", err)
		return err
	}

	return nil
}

func deleteEndpoint(ctx p.Context, id string) error {
	config := infer.GetConfig[Config](ctx)
	gqlVariable := map[string]interface{}{"id": id}

	gqlInput := GqlInput{
		Query: `
		mutation DeleteEndpoint ($id: String!) {
			deleteEndpoint(id: $id)
		}`,
		Variables: gqlVariable,
	}

	jsonValue, err := json.Marshal(gqlInput)
	if err != nil {
		return err
	}

	url := URL + config.Token

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 20}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var output struct {
		Errors []struct {
			Message string
		}
	}

	err = json.Unmarshal(data, &output)
	if err != nil {
		return err
	}

	if len(output.Errors) > 0 {
		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
		return err
	}

	return nil
}

// func (*Endpoint) WireDependencies(f infer.FieldSelector, args *EndpointArgs, state *EndpointState) {
// 	f.OutputField(&state.TemplateId).DependsOn(f.InputField(&args.TemplateId))
// }
