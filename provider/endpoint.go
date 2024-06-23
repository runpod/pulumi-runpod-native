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

type EndpointState struct {
	EndpointArgs
	Endpoint Endpoint `pulumi:"endpoint"`
}

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
	Name            string `pulumi:"name"`
	Id              string `pulumi:"id"`
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
	Name            string `pulumi:"name" structs:"name,omitempty"`
	TemplateId      string `pulumi:"templateId" structs:"templateId,omitempty"`
	GpuIds          string `pulumi:"gpuIds" structs:"gpuIds,omitempty"`
	IdleTimeout     int    `pulumi:"idleTimeout,optional" structs:"idleTimeout,omitempty"`
	Locations       string `pulumi:"locations,optional" structs:"locations,omitempty"`
	NetworkVolumeId string `pulumi:"networkVolumeId,optional" structs:"networkVolumeId,omitempty"`
	ScalerType      string `pulumi:"scalerType,optional" structs:"scalerType,omitempty"`
	ScalerValue     int    `pulumi:"scalerValue,optional" structs:"scalerValue,omitempty"`
	WorkersMax      int    `pulumi:"workersMax,optional" structs:"workersMax,omitempty"`
	WorkersMin      int    `pulumi:"workersMin,optional" structs:"workersMin,omitempty"`
}

func (*Endpoint) Create(ctx p.Context, name string, input EndpointArgs, preview bool) (string, EndpointState, error) {
	state := EndpointState{EndpointArgs: input}
	if preview {
		return name, state, nil
	}
	config := infer.GetConfig[Config](ctx)

	if input.Name == "" || input.GpuIds == "" || input.TemplateId == "" {
		return name, state, fmt.Errorf("TemplateId, gpuIds and name are required")
	}

	gqlVariable := structs.Map(input)
	// "AMPERE_16"
	// # append -fb to your endpoint's name to enable FlashBoot

	println(gqlVariable)
	gqlInput := GqlInput{
		Query: `
		mutation SaveEndpoint (
	    gpuIds: String!
	    templateId: String!
	    name: String!
	    idleTimeout: Int 
	    locations: String
	    networkVolumeId: String
	    scalerType: String
	    scalerValue: Int
	    workersMax: Int
	    workersMin: Int 
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

// func (*Endpoint) Update(ctx p.Context, id string, olds EndpointState, news EndpointArgs, preview bool) (EndpointState, error) {
// 	state := EndpointState{EndpointArgs: news}
// 	if preview {
// 		return state, nil
// 	}
// 	config := infer.GetConfig[Config](ctx)

// 	if news.ImageName == "" || news.Readme == "" || news.Name == "" {
// 		return state, fmt.Errorf("imageName, readme and name are required")
// 	}

// 	gqlVariable := structs.Map(news)
// 	gqlVariable["id"] = olds.Endpoint.Id
// 	gqlInput := GqlInput{
// 		Query: `
// 		mutation SaveEndpoint (
// 			$id: String!
// 			$containerDiskInGb: Int!
// 			$containerRegistryAuthId: String
// 			$dockerArgs: String!
// 			$env: [EnvironmentVariableInput!]!
// 			$imageName: String!
// 			$isPublic: Boolean
// 			$isServerless: Boolean
// 			$name: String!
// 			$ports: String
// 			$readme: String
// 			$startJupyter: Boolean
// 			$startSsh: Boolean
// 			$volumeInGb: Int!
// 			$volumeMountPath: String
// 		) {
// 		saveEndpoint(input: {
// 			id: $id,
// 			containerDiskInGb: $containerDiskInGb,
// 			containerRegistryAuthId: $containerRegistryAuthId,
// 			dockerArgs: $dockerArgs,
// 			env: $env,
// 			imageName: $imageName,
// 			isPublic: $isPublic,
// 			isServerless: $isServerless,
// 			name: $name,
// 			ports: $ports,
// 			readme: $readme,
// 			startJupyter: $startJupyter,
// 			startSsh: $startSsh,
// 			volumeInGb: $volumeInGb,
// 			volumeMountPath: $volumeMountPath
// 		}) {
// 			id
// 			imageName
// 			name
// 		}
// 	}`,
// 		Variables: gqlVariable,
// 	}

// 	jsonValue, err := json.Marshal(gqlInput)
// 	if err != nil {
// 		return state, err
// 	}

// 	url := URL + config.Token

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
// 	if err != nil {
// 		return state, err
// 	}

// 	req.Header.Add("Content-Type", "application/json")
// 	client := &http.Client{Timeout: time.Second * 20}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return state, err
// 	}

// 	defer resp.Body.Close()

// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return state, err
// 	}

// 	output := &OutputDeployEndpoint{}
// 	err = json.Unmarshal(data, output)
// 	if err != nil {
// 		return state, err
// 	}

// 	if len(output.Errors) > 0 {
// 		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
// 		return state, err
// 	}

// 	endpoint := output.Data.SaveEndpoint
// 	if endpoint.Id == "" {
// 		err = fmt.Errorf("graphql endpoint is nil: %s", string(data))
// 		return state, err
// 	}

// 	state.Endpoint = endpoint

// 	return state, nil
// }

// func (*Endpoint) Diff(ctx p.Context, id string, olds EndpointState, news EndpointArgs) (p.DiffResponse, error) {

// 	diff := map[string]p.PropertyDiff{}

// 	if !reflect.DeepEqual(news.Env, olds.Env) {
// 		diff["env"] = p.PropertyDiff{Kind: p.Update}
// 	}
// 	if news.ImageName != olds.ImageName {
// 		diff["imageName"] = p.PropertyDiff{Kind: p.Update}
// 	}
// 	if news.Ports != olds.Ports {
// 		diff["ports"] = p.PropertyDiff{Kind: p.Update}
// 	}
// 	if news.ContainerDiskInGb != olds.ContainerDiskInGb {
// 		diff["containerDiskInGb"] = p.PropertyDiff{Kind: p.Update}
// 	}
// 	if news.VolumeInGb != olds.VolumeInGb {
// 		diff["volumeInGb"] = p.PropertyDiff{Kind: p.Update}
// 	}
// 	if news.VolumeMountPath != olds.VolumeMountPath {
// 		diff["volumeMountPath"] = p.PropertyDiff{Kind: p.Update}
// 	}
// 	if news.ContainerRegistryAuthId != olds.ContainerRegistryAuthId {
// 		diff["containerRegistryAuthId"] = p.PropertyDiff{Kind: p.Update}
// 	}
// 	if news.DockerArgs != olds.DockerArgs {
// 		diff["dockerArgs"] = p.PropertyDiff{Kind: p.Update}
// 	}
// 	if news.IsPublic != olds.IsPublic {
// 		diff["isPublic"] = p.PropertyDiff{Kind: p.Update}
// 	}
// 	if news.IsServerless != olds.IsServerless {
// 		diff["isServerless"] = p.PropertyDiff{Kind: p.Update}
// 	}
// 	if news.Name != olds.Name {
// 		diff["name"] = p.PropertyDiff{Kind: p.Update}
// 	}
// 	if news.Readme != olds.Readme {
// 		diff["readme"] = p.PropertyDiff{Kind: p.Update}
// 	}
// 	if news.StartJupyter != olds.StartJupyter {
// 		diff["startJupyter"] = p.PropertyDiff{Kind: p.Update}
// 	}
// 	if news.StartSsh != olds.StartSsh {
// 		diff["startSsh"] = p.PropertyDiff{Kind: p.Update}
// 	}

// 	return p.DiffResponse{
// 		DeleteBeforeReplace: true,
// 		HasChanges:          len(diff) > 0,
// 		DetailedDiff:        diff,
// 	}, nil
// }

// func (*Endpoint) Delete(ctx p.Context, id string, props EndpointState) error {
// 	config := infer.GetConfig[Config](ctx)
// 	gqlVariable := map[string]interface{}{"id": props.Endpoint.Id}

// 	gqlInput := GqlInput{
// 		Query: `
// 		mutation DeleteEndpoint ($id: String!) {
// 			deleteTemplate(id: $tid)
// 		}`,
// 		Variables: gqlVariable,
// 	}

// 	jsonValue, err := json.Marshal(gqlInput)
// 	if err != nil {
// 		return err
// 	}

// 	url := URL + config.Token

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Add("Content-Type", "application/json")
// 	client := &http.Client{Timeout: time.Second * 20}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}

// 	defer resp.Body.Close()

// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	var output struct {
// 		Errors []struct {
// 			Message string
// 		}
// 	}

// 	err = json.Unmarshal(data, &output)
// 	if err != nil {
// 		return err
// 	}

// 	if len(output.Errors) > 0 {
// 		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
// 		return err
// 	}

// 	return nil

// }
