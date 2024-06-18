package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/fatih/structs"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Template struct {
	AdvancedStart           *bool                      `pulumi:"advancedStart" json:"advancedStart"`
	ContainerDiskInGb       int                        `pulumi:"containerDiskInGb" json:"containerDiskInGb"`
	ContainerRegistryAuthId *string                    `pulumi:"containerRegistryAuthId" json:"containerRegistryAuthId"`
	DockerArgs              *string                    `pulumi:"dockerArgs" json:"dockerArgs,omitempty"`
	Earned                  *float64                   `pulumi:"earned" json:"earned"`
	Env                     []EnvironmentVariableInput `pulumi:"env" json:"env"`
	Id                      *string                    `pulumi:"id" json:"id"`
	ImageName               string                     `pulumi:"imageName" json:"imageName"`
	IsPublic                *bool                      `pulumi:"isPublic" json:"isPublic"`
	IsServerless            *bool                      `pulumi:"isServerless" json:"isServerless"`
	BoundEndpointID         *string                    `pulumi:"boundEndpointID" json:"boundEndpointID"`
	Name                    string                     `pulumi:"name" json:"name"`
	Ports                   *string                    `pulumi:"ports" json:"ports"`
	Readme                  *string                    `pulumi:"readme" json:"readme"`
	RuntimeInMinutes        *int                       `pulumi:"runtimeInMinutes" json:"runtimeInMinutes"`
	StartJupyter            *bool                      `pulumi:"startJupyter" json:"startJupyter"`
	StartScript             *string                    `pulumi:"startScript" json:"startScript"`
	StartSsh                *bool                      `pulumi:"startSsh" json:"startSsh"`
	VolumeInGb              int                        `pulumi:"volumeInGb" json:"volumeInGb"`
	VolumeMountPath         string                     `pulumi:"volumeMountPath" json:"volumeMountPath"`
	Config                  *map[string]interface{}    `pulumi:"config" json:"config"`
	Category                string                     `pulumi:"category" json:"category`
}

type EnvironmentVariableInput struct {
	Key   string `pulumi:"key" json:"key"`
	Value string `pulumi:"value" json:"value"`
}

type TemplateArgs struct {
	AdvancedStart           *bool                      `pulumi:"advancedStart,optional" json:"advancedStart,omitempty"`
	ContainerDiskInGb       int                        `pulumi:"containerDiskInGb" json:"containerDiskInGb"`
	ContainerRegistryAuthId *string                    `pulumi:"containerRegistryAuthId,optional" json:"containerRegistryAuthId,omitempty"`
	DockerArgs              string                     `pulumi:"dockerArgs,optional" json:"dockerArgs"`
	Env                     []EnvironmentVariableInput `pulumi:"env" json:"env"`
	ImageName               string                     `pulumi:"imageName" json:"imageName"`
	IsPublic                *bool                      `pulumi:"isPublic,optional" json:"isPublic,omitempty"`
	IsServerless            *bool                      `pulumi:"isServerless,optional" json:"isServerless,omitempty"`
	Name                    string                     `pulumi:"name" json:"name"`
	Ports                   *string                    `pulumi:"ports,optional"	json:"ports,omitempty"`
	Readme                  string                     `pulumi:"readme,optional" json:"readme"`
	StartJupyter            *bool                      `pulumi:"startJupyter,optional" json:"startJupyter,omitempty"`
	StartScript             *string                    `pulumi:"startScript,optional" json:"startScript,omitempty"`
	StartSsh                *bool                      `pulumi:"startSsh,optional" json:"startSsh,omitempty"`
	VolumeInGb              int                        `pulumi:"volumeInGb" json:"volumeInGb"`
	VolumeMountPath         string                     `pulumi:"volumeMountPath" json:"volumeMountPath"`
	Config                  *map[string]interface{}    `pulumi:"config,optional" json:"config,omitempty"`
	Category                string                     `pulumi:"category" json:"category"`
}

type TemplateUpdateArgs struct {
	TemplateArgs
	Id string `pulumi:"id"`
}

type TemplateState struct {
	TemplateUpdateArgs
	TemplateArgs
	Template Template `pulumi:"template"`
}

type Message struct {
	Message string
}

type TemplateData struct {
	SaveTemplate Template `json:"saveTemplate"`
}

type CreateTemplateOutput struct {
	Errors []Message    `json:"errors"`
	Data   TemplateData `json:"data"`
}

type UpdateTemplateOutput struct {
	Errors []struct {
		Message string
	}
	Data struct {
		UpdateTemplate Template
	}
}

func (*Template) Create(ctx p.Context, name string, input TemplateArgs, preview bool) (string, TemplateState, error) {
	state := TemplateState{TemplateArgs: input}
	if preview {
		return name, state, nil
	}

	if input.Category != "CPU" && input.Category != "NVIDIA" && input.Category != "AMD" {
		return name, state, fmt.Errorf("category must be one of CPU, NVIDIA, or AMD")
	}

	gqlVariable, err := json.Marshal(input)
	if err != nil {
		return name, state, err
	}

	var a map[string]interface{}
	err = json.Unmarshal(gqlVariable, &a)
	if err != nil {
		return name, state, err
	}

	environmentVariables := []EnvironmentVariableInput{}
	envVarsBytes, err := json.Marshal(input.Env)
	if err != nil {
		return name, state, err
	}
	err = json.Unmarshal(envVarsBytes, &environmentVariables)
	if err != nil {
		return name, state, err
	}

	a["env"] = environmentVariables

	gqlInput := GqlInput{
		Query: `
		mutation TemplateCreate (
			$advancedStart: Boolean
			$containerDiskInGb: Int!
			$containerRegistryAuthId: String
			$dockerArgs: String!
			$env: [EnvironmentVariableInput]!
			$imageName: String!
			$isPublic: Boolean
			$isServerless: Boolean
			$name: String!
			$ports: String
			$readme: String
			$startJupyter: Boolean
			$startScript: String
			$startSsh: Boolean
			$volumeInGb: Int!
			$volumeMountPath: String
			$config: JSON
			$category: TemplateCategory
		) {
			saveTemplate(input: {				
				advancedStart: $advancedStart
				containerDiskInGb: $containerDiskInGb
				containerRegistryAuthId: $containerRegistryAuthId
				dockerArgs: $dockerArgs
				env: $env
				imageName: $imageName
				isPublic: $isPublic
				isServerless: $isServerless
				name: $name
				ports: $ports
				readme: $readme
				startJupyter: $startJupyter
				startScript: $startScript
				startSsh: $startSsh
				volumeInGb: $volumeInGb
				volumeMountPath: $volumeMountPath
				config: $config
				category: $category
			}) {
				advancedStart
				containerDiskInGb
				containerRegistryAuthId
				dockerArgs
				earned
				env {
					key
					value
				}
				id
				imageName
				isPublic
				isRunpod
				isServerless
				boundEndpointId
				name
				ports
				readme
				runtimeInMin
				startJupyter
				startScript
				startSsh
				volumeInGb
				volumeMountPath
				config
				category
			}
		}`,
		Variables: a,
	}

	jsonValue, err := json.Marshal(gqlInput)
	if err != nil {
		return name, state, err
	}

	config := infer.GetConfig[Config](ctx)
	url := URL + config.Token

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return name, state, err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return name, state, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return name, state, err
	}

	output := &CreateTemplateOutput{}
	err = json.Unmarshal(data, output)
	if err != nil {
		return name, state, err
	}

	if len(output.Errors) > 0 {
		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
		return name, state, err
	}

	newTemplate := output.Data.SaveTemplate
	log.Print("new template", newTemplate)
	if newTemplate.Id == nil {
		err = fmt.Errorf("graphql template is nil: %s", string(data))
		return name, state, err
	}

	state.Template = newTemplate

	return name, state, nil
}

func (*Template) Update(ctx p.Context, id string, olds TemplateState, updateArgs TemplateUpdateArgs, preview bool) (TemplateState, error) {
	state := TemplateState{TemplateUpdateArgs: updateArgs}
	if preview {
		return state, nil
	}
	config := infer.GetConfig[Config](ctx)

	if updateArgs.Category != "CPU" && updateArgs.Category != "NVIDIA" && updateArgs.Category != "AMD" {
		return TemplateState{}, fmt.Errorf("category must be one of CPU, NVIDIA, or AMD")
	}

	gqlVariable := structs.Map(updateArgs)
	gqlVariable["id"] = olds.Template.Id
	gqlInput := GqlInput{
		Query: `
		mutation TemplateCreate (
			$advancedStart: Boolean
			$containerDiskInGb: Int!
			$containerRegistryAuthId: String
			$dockerArgs: String!
			$env: [EnvironmentVariableInput]!
			$id: String
			$imageName: String!
			$isPublic: Boolean
			$isServerless: Boolean
			$name: String!
			$ports: String
			$readme: String
			$startJupyter: Boolean
			$startScript: String
			$startSsh: Boolean
			$volumeInGb: Int!
			$volumeMountPath: String
			$config: JSON
			$category: TemplateCategory
		) {
			saveTemplate(input: {				
				advancedStart: $advancedStart
				containerDiskInGb: $containerDiskInGb
				containerRegistryAuthId: $containerRegistryAuthId
				dockerArgs: $dockerArgs
				env: $env
				imageName: $imageName
				isPublic: $isPublic
				isServerless: $isServerless
				name: $name
				ports: $ports
				readme: $readme
				startJupyter: $startJupyter
				startScript: $startScript
				startSsh: $startSsh
				volumeInGb: $volumeInGb
				volumeMountPath: $volumeMountPath
				config: $config
				category: $category
				id: $id
			}) {
				advancedStart
				containerDiskInGb
				containerRegistryAuthId
				dockerArgs
				earned
				env {
					key
					value
				}
				id
				imageName
				isPublic
				isRunpod
				isServerless
				boundEndpointId
				name
				ports
				readme
				runtimeInMin
				startJupyter
				startScript
				startSsh
				volumeInGb
				volumeMountPath
				config
				category
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

	output := &UpdateTemplateOutput{}
	err = json.Unmarshal(data, output)
	if err != nil {
		return state, err
	}

	if len(output.Errors) > 0 {
		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
		return state, err
	}

	updatedTemplate := output.Data.UpdateTemplate
	if updatedTemplate.Id == nil {
		err = fmt.Errorf("graphql updated template is nil: %s", string(data))
		return state, err
	}

	state.Template = updatedTemplate

	return state, nil
}

func CompareEnvChanges(new []EnvironmentVariableInput, old []EnvironmentVariableInput) bool {
	if len(new) != len(old) {
		return true
	}

	for i, v := range new {
		if v.Key != old[i].Key || v.Value != old[i].Value {
			return true
		}
	}

	return false
}

func (*Template) Diff(ctx p.Context, id string, old TemplateState, new TemplateArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	if new.AdvancedStart != old.AdvancedStart {
		diff["advancedStart"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.ContainerDiskInGb != old.ContainerDiskInGb {
		diff["containerDiskInGb"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.ContainerRegistryAuthId != old.ContainerRegistryAuthId {
		diff["containerRegistryAuthId"] = p.PropertyDiff{Kind: p.Update}
	}

	// TODO: bring it back
	if new.DockerArgs != old.DockerArgs {
		diff["dockerArgs"] = p.PropertyDiff{Kind: p.Update}
	}

	if CompareEnvChanges(new.Env, old.Env) {
		diff["env"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.ImageName != old.ImageName {
		diff["imageName"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.IsPublic != old.IsPublic {
		diff["isPublic"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.IsServerless != old.IsServerless {
		diff["isServerless"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.Name != old.Name {
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.Ports != old.Ports {
		diff["ports"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.Readme != old.Readme {
		diff["readme"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.StartJupyter != old.StartJupyter {
		diff["startJupyter"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.StartScript != old.StartScript {
		diff["startScript"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.StartSsh != old.StartSsh {
		diff["startSsh"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.VolumeInGb != old.VolumeInGb {
		diff["volumeInGb"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.VolumeMountPath != old.VolumeMountPath {
		diff["volumeMountPath"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.Config != old.Config {
		diff["config"] = p.PropertyDiff{Kind: p.Update}
	}

	if new.Category != old.Category {
		diff["category"] = p.PropertyDiff{Kind: p.Update}
	}

	log.Printf("diff", diff)

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*Template) Delete(ctx p.Context, id string, props TemplateState) error {
	config := infer.GetConfig[Config](ctx)
	gqlVariable := map[string]interface{}{"id": props.Template.Id}

	gqlInput := GqlInput{
		Query: `
		mutation deleteTemplate (
			$templateName: String
		) {
			deleteTemplate(input: {				
				templateName: $templateName
			})
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
	client := &http.Client{}

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
		err = fmt.Errorf("Could not delete template: %s", output.Errors[0].Message)
		return err
	}

	return nil
}
