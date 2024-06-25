package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"reflect"
	"time"

	"log"

	"github.com/fatih/structs"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type TemplateState struct {
	TemplateArgs
	Template Template `pulumi:"template"`
}

type OutputDeployTemplate struct {
	Errors []struct {
		Message   string
		Locations []struct {
			Line   int
			Column int
		}
	}
	Data struct {
		SaveTemplate Template
	}
}

type Template struct {
	AdvancedStart           bool     `pulumi:"advancedStart"`
	ContainerDiskInGb       int      `pulumi:"containerDiskInGb"`
	ContainerRegistryAuthId string   `pulumi:"containerRegistryAuthId"`
	DockerArgs              string   `pulumi:"dockerArgs"`
	Earned                  float64  `pulumi:"earned"`
	Env                     []PodEnv `pulumi:"env,optional"`
	Id                      string   `pulumi:"id"`
	ImageName               string   `pulumi:"imageName"`
	IsPublic                bool     `pulumi:"isPublic"`
	IsRunpod                bool     `pulumi:"isRunpod"`
	IsServerless            bool     `pulumi:"isServerless"`
	BoundEndpointId         string   `pulumi:"boundEndpointId"`
	Name                    string   `pulumi:"name"`
	Ports                   string   `pulumi:"ports"`
	Readme                  string   `pulumi:"readme"`
	RuntimeInMin            int      `pulumi:"runtimeInMin"`
	StartJupyter            bool     `pulumi:"startJupyter"`
	StartScript             string   `pulumi:"startScript"`
	StartSsh                bool     `pulumi:"startSsh"`
	VolumeInGb              int      `pulumi:"volumeInGb"`
	VolumeMountPath         string   `pulumi:"volumeMountPath"`
	// Config                  interface{} `pulumi:"config"`
	Category string `pulumi:"category"`
}

type TemplateArgs struct {
	ContainerDiskInGb       int      `pulumi:"containerDiskInGb" structs:"containerDiskInGb,omitempty"`
	ContainerRegistryAuthId string   `pulumi:"containerRegistryAuthId,optional" structs:"containerRegistryAuthId,omitempty"`
	DockerArgs              string   `pulumi:"dockerArgs" structs:"dockerArgs,omitempty"`
	Env                     []PodEnv `pulumi:"env" structs:"env,omitempty"`
	ImageName               string   `pulumi:"imageName" structs:"imageName,omitempty"`
	IsPublic                bool     `pulumi:"isPublic,optional" structs:"isPublic,omitempty"`
	IsServerless            bool     `pulumi:"isServerless,optional" structs:"isServerless,omitempty"`
	Name                    string   `pulumi:"name" structs:"name,omitempty"`
	Ports                   string   `pulumi:"ports,optional" structs:"ports,omitempty"`
	Readme                  string   `pulumi:"readme,optional" structs:"readme,omitempty"`
	StartJupyter            bool     `pulumi:"startJupyter,optional" structs:"startJupyter,omitempty"`
	StartSsh                bool     `pulumi:"startSsh,optional" structs:"startSsh,omitempty"`
	VolumeInGb              int      `pulumi:"volumeInGb" structs:"volumeInGb"`
	VolumeMountPath         string   `pulumi:"volumeMountPath,optional" structs:"volumeMountPath,omitempty"`
}

func (*Template) Create(ctx p.Context, name string, input TemplateArgs, preview bool) (string, TemplateState, error) {
	state := TemplateState{TemplateArgs: input}
	if preview {
		return name, state, nil
	}
	config := infer.GetConfig[Config](ctx)

	if input.ImageName == "" || input.Readme == "" || input.Name == "" {
		return name, state, fmt.Errorf("imageName, readme and name are required")
	}

	gqlVariable := structs.Map(input)

	gqlInput := GqlInput{
		Query: `
		mutation SaveTemplate (
			$containerDiskInGb: Int!
			$containerRegistryAuthId: String
			$dockerArgs: String!
			$env: [EnvironmentVariableInput!]!
			$imageName: String!
			$isPublic: Boolean
			$isServerless: Boolean
			$name: String!
			$ports: String
			$readme: String
			$startJupyter: Boolean
			$startSsh: Boolean
			$volumeInGb: Int!
			$volumeMountPath: String
		) {
		saveTemplate(input: {
			containerDiskInGb: $containerDiskInGb,
			containerRegistryAuthId: $containerRegistryAuthId,
			dockerArgs: $dockerArgs,
			env: $env,
			imageName: $imageName,
			isPublic: $isPublic,
			isServerless: $isServerless,
			name: $name,
			ports: $ports,
			readme: $readme,
			startJupyter: $startJupyter,
			startSsh: $startSsh,
			volumeInGb: $volumeInGb,
			volumeMountPath: $volumeMountPath
		}) {
			id
			imageName
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

	output := &OutputDeployTemplate{}
	err = json.Unmarshal(data, output)
	if err != nil {
		return name, state, err
	}

	log.Print("output: ", output, gqlInput, string(data))

	if len(output.Errors) > 0 {
		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
		return name, state, err
	}

	template := output.Data.SaveTemplate
	if template.Id == "" {
		err = fmt.Errorf("graphql template is nil: %s", string(data))
		return name, state, err
	}

	state.Template = template

	return name, state, nil
}

func (*Template) Update(ctx p.Context, id string, olds TemplateState, news TemplateArgs, preview bool) (TemplateState, error) {
	state := TemplateState{TemplateArgs: news}
	if preview {
		return state, nil
	}
	config := infer.GetConfig[Config](ctx)

	if news.ImageName == "" || news.Readme == "" || news.Name == "" {
		return state, fmt.Errorf("imageName, readme and name are required")
	}

	gqlVariable := structs.Map(news)
	gqlVariable["id"] = olds.Template.Id
	gqlInput := GqlInput{
		Query: `
		mutation SaveTemplate (
			$id: String!
			$containerDiskInGb: Int!
			$containerRegistryAuthId: String
			$dockerArgs: String!
			$env: [EnvironmentVariableInput!]!
			$imageName: String!
			$isPublic: Boolean
			$isServerless: Boolean
			$name: String!
			$ports: String
			$readme: String
			$startJupyter: Boolean
			$startSsh: Boolean
			$volumeInGb: Int!
			$volumeMountPath: String
		) {
		saveTemplate(input: {
			id: $id,
			containerDiskInGb: $containerDiskInGb,
			containerRegistryAuthId: $containerRegistryAuthId,
			dockerArgs: $dockerArgs,
			env: $env,
			imageName: $imageName,
			isPublic: $isPublic,
			isServerless: $isServerless,
			name: $name,
			ports: $ports,
			readme: $readme,
			startJupyter: $startJupyter,
			startSsh: $startSsh,
			volumeInGb: $volumeInGb,
			volumeMountPath: $volumeMountPath
		}) {
			id
			imageName
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

	output := &OutputDeployTemplate{}
	err = json.Unmarshal(data, output)
	if err != nil {
		return state, err
	}

	if len(output.Errors) > 0 {
		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
		return state, err
	}

	template := output.Data.SaveTemplate
	if template.Id == "" {
		err = fmt.Errorf("graphql template is nil: %s", string(data))
		return state, err
	}

	state.Template = template

	return state, nil
}

func (*Template) Diff(ctx p.Context, id string, olds TemplateState, news TemplateArgs) (p.DiffResponse, error) {

	diff := map[string]p.PropertyDiff{}

	if !reflect.DeepEqual(news.Env, olds.Env) {
		diff["env"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.ImageName != olds.ImageName {
		diff["imageName"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.Ports != olds.Ports {
		diff["ports"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.ContainerDiskInGb != olds.ContainerDiskInGb {
		diff["containerDiskInGb"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.VolumeInGb != olds.VolumeInGb {
		diff["volumeInGb"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.VolumeMountPath != olds.VolumeMountPath {
		diff["volumeMountPath"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.ContainerRegistryAuthId != olds.ContainerRegistryAuthId {
		diff["containerRegistryAuthId"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.DockerArgs != olds.DockerArgs {
		diff["dockerArgs"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.IsPublic != olds.IsPublic {
		diff["isPublic"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.IsServerless != olds.IsServerless {
		diff["isServerless"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.Name != olds.Name {
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.Readme != olds.Readme {
		diff["readme"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.StartJupyter != olds.StartJupyter {
		diff["startJupyter"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.StartSsh != olds.StartSsh {
		diff["startSsh"] = p.PropertyDiff{Kind: p.Update}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*Template) Delete(ctx p.Context, id string, props TemplateState) error {
	config := infer.GetConfig[Config](ctx)
	gqlVariable := map[string]interface{}{"templateName": props.Template.Name}

	gqlInput := GqlInput{
		Query: `
		mutation DeleteTemplate ($templateName: String!) {
			deleteTemplate(templateName: $templateName)
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
