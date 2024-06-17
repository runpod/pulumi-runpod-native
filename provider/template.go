package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fatih/structs"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type TemplateState struct {
	TemplateArgs
	Template Template `pulumi:"template"`
}

type SaveTemplateOutput struct {
	Errors []struct {
		Message string
	}
	Data struct {
		SaveTemplate Template
	}
}

type Template struct {
	AdvancedStart           string   `pulumi:"advancedStart"`
	ContainerDiskInGb       int      `pulumi:"containerDiskInGb"`
	ContainerRegistryAuthId string   `pulumi:"containerRegistryAuthId"`
	DockerArgs              string   `pulumi:"dockerArgs"`
	Env                     []string `pulumi:"env,optional"`
	Id                      string   `pulumi:"id"`
	ImageName               string   `pulumi:"imageName"`
	Name                    string   `pulumi:"name"`
	Readme                  string   `pulumi:"readme"`
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
	VolumeInGb              int      `pulumi:"volumeInGb" structs:"volumeInGb,omitempty"`
	VolumeMountPath         string   `pulumi:"volumeMountPath,optional" structs:"volumeMountPath,omitempty"`
}

func (*Template) Create(ctx p.Context, name string, input TemplateArgs, preview bool) (string, TemplateState, error) {
	state := TemplateState{TemplateArgs: input}
	if preview {
		return name, state, nil
	}
	config := infer.GetConfig[Config](ctx)

	gqlVariable := structs.Map(input)

	gqlInput := GqlInput{
		Query: `
		mutation (			
			$containerDiskInGb: Int!
			$dockerArgs: String!
			$env: [EnvironmentVariableInput!]!
			$imageName: String!
			$name: String!
			$volumeInGb: Int
			$readme: String
			$isPublic: Boolean
			$isServerless: Boolean
			$ports: String
			$startJupyter: Boolean
			$startSsh: Boolean
			$volumeMountPath: String
			$containerRegistryAuthId: String
		) {
			saveTemplate(input: {				
				name: $name
				containerDiskInGb: $containerDiskInGb
				volumeInGb: $volumeInGb
				imageName: $imageName
				env: $env
				dockerArgs: $dockerArgs
				ports: $ports
				readme: $readme
				isPublic: $isPublic
				isServerless: $isServerless
				volumeMountPath: $volumeMountPath
				startJupyter: $startJupyter					
			}) {
        advancedStart
        containerDiskInGb
      	dockerArgs
				env {
					key
					value
				}
				id
				imageName
				name
				ports
				readme
    }
  }`,
		Variables: gqlVariable,
	}

	jsonValue, err := json.Marshal(gqlInput)
	if err != nil {
		return name, state, err
	}

	url := "https://api.runtemplate.io/graphql?api_key=" + config.Token

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return name, state, err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 20}

	resp, err := client.Do(req)
	if err != nil {
		return name, state, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return name, state, err
	}

	output := &SaveTemplateOutput{}
	err = json.Unmarshal(data, output)
	if err != nil {
		return name, state, err
	}

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

// func (*Template) Update(ctx p.Context, id string, olds TemplateState, news TemplateArgs, preview bool) (TemplateState, error) {
// 	state := TemplateState{TemplateArgs: news}
// 	if preview {
// 		return state, nil
// 	}
// 	config := infer.GetConfig[Config](ctx)

// 	gqlVariable := structs.Map(news)
// 	gqlVariable["templateId"] = olds.Template.Id
// 	gqlInput := GqlInput{
// 		Query: `
// 		mutation UpdatePodMutation (
// 			$templateId: String!
// 			$dockerArgs: String
// 			$imageName: String!
// 			$env: [EnvironmentVariableInput]
// 			$port: Port
// 			$ports: String
// 			$containerDiskInGb: Int!
// 			$volumeInGb: Int
// 			$volumeMountPath: String
// 			$containerRegistryAuthId: String
// 		) {
// 			templateEditJob(input: {
// 				templateId: $templateId,
// 				dockerArgs: $dockerArgs,
// 				imageName: $imageName,
// 				env: $env,
// 				port: $port,
// 				ports: $ports,
// 				containerDiskInGb: $containerDiskInGb,
// 				volumeInGb: $volumeInGb,
// 				volumeMountPath: $volumeMountPath,
// 				containerRegistryAuthId: $containerRegistryAuthId,
// 			}) {
// 				id
//     			imageName
//     			machineId
// 				containerDiskInGb
// 			}
// 		}`,
// 		Variables: gqlVariable,
// 	}

// 	jsonValue, err := json.Marshal(gqlInput)
// 	if err != nil {
// 		return state, err
// 	}

// 	url := "https://api.runtemplate.io/graphql?api_key=" + config.Token

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

// 	output := &OutputUpdatePod{}
// 	err = json.Unmarshal(data, output)
// 	if err != nil {
// 		return state, err
// 	}

// 	if len(output.Errors) > 0 {
// 		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
// 		return state, err
// 	}

// 	template := output.Data.PodEditJob
// 	if template.Id == "" {
// 		err = fmt.Errorf("graphql template is nil: %s", string(data))
// 		return state, err
// 	}

// 	state.Template = template

// 	return state, nil
// }

// func (*Template) Diff(ctx p.Context, id string, olds TemplateState, news PodArgs) (p.DiffResponse, error) {

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
// 	// if news.ContainerRegistryAuthId != olds.ContainerRegistryAuthId {
// 	// 	diff["containerRegistryAuthId"] = p.PropertyDiff{Kind: p.Update}
// 	// }
// 	if news.DockerArgs != olds.DockerArgs {
// 		diff["dockerArgs"] = p.PropertyDiff{Kind: p.Update}
// 	}

// 	return p.DiffResponse{
// 		DeleteBeforeReplace: true,
// 		HasChanges:          len(diff) > 0,
// 		DetailedDiff:        diff,
// 	}, nil
// }

// func (*Template) Delete(ctx p.Context, id string, props TemplateState) error {
// 	config := infer.GetConfig[Config](ctx)
// 	gqlVariable := map[string]interface{}{"templateId": props.Template.Id}

// 	gqlInput := GqlInput{
// 		Query: `
// 		mutation templateTerminateMutation (
// 			$templateId: String!
// 		) {
// 			templateTerminate(input: {
// 				templateId: $templateId
// 			})
// 		}`,
// 		Variables: gqlVariable,
// 	}

// 	jsonValue, err := json.Marshal(gqlInput)
// 	if err != nil {
// 		return err
// 	}

// 	url := "https://api.runtemplate.io/graphql?api_key=" + config.Token

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
