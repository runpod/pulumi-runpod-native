package provider

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type TemplateState struct {
	TemplateArgs
	Template Template `pulumi:"template"`
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
	ImageName               string   `pulumi:"imageName,optional" structs:"imageName,omitempty"`
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

	// Validate required fields according to actual GraphQL schema
	if input.Name == "" {
		return name, state, fmt.Errorf("name is required")
	}

	// ImageName is optional but if provided, cannot be empty
	if input.ImageName != "" && len(strings.TrimSpace(input.ImageName)) == 0 {
		return name, state, fmt.Errorf("imageName cannot be empty if provided")
	}

	// Create GraphQL client
	client := NewGraphQLClient(config.Token)

	// Convert PodEnv to EnvironmentVariableInput
	var envVars []*EnvironmentVariableInput
	for _, env := range input.Env {
		envVars = append(envVars, &EnvironmentVariableInput{
			Key:   env.Key,
			Value: env.Value,
		})
	}

	// Call the generated SaveTemplate mutation
	result, err := client.GetClient().SaveTemplate(
		context.Background(),
		nil, // id - nil for create
		input.ContainerDiskInGb,
		&input.ContainerRegistryAuthId,
		input.DockerArgs,
		envVars,
		input.ImageName,
		&input.IsPublic,
		&input.IsServerless,
		input.Name,
		&input.Ports,
		&input.Readme,
		&input.StartJupyter,
		&input.StartSsh,
		input.VolumeInGb,
		&input.VolumeMountPath,
	)
	if err != nil {
		return name, state, fmt.Errorf("failed to save template: %w", err)
	}

	if result.SaveTemplate == nil {
		return name, state, fmt.Errorf("template creation failed: no template returned")
	}

	// Convert result to Template struct
	template := Template{
		Id:        *result.SaveTemplate.ID,
		ImageName: *result.SaveTemplate.ImageName,
		Name:      *result.SaveTemplate.Name,
		// Set other fields from input since they're not returned by the mutation
		ContainerDiskInGb:       input.ContainerDiskInGb,
		ContainerRegistryAuthId: input.ContainerRegistryAuthId,
		DockerArgs:              input.DockerArgs,
		Env:                     input.Env,
		IsPublic:                input.IsPublic,
		IsServerless:            input.IsServerless,
		Ports:                   input.Ports,
		Readme:                  input.Readme,
		StartJupyter:            input.StartJupyter,
		StartSsh:                input.StartSsh,
		VolumeInGb:              input.VolumeInGb,
		VolumeMountPath:         input.VolumeMountPath,
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

	// Validate required fields according to actual GraphQL schema
	if news.Name == "" {
		return state, fmt.Errorf("name is required")
	}

	// ImageName is optional but if provided, cannot be empty
	if news.ImageName != "" && len(strings.TrimSpace(news.ImageName)) == 0 {
		return state, fmt.Errorf("imageName cannot be empty if provided")
	}

	// Create GraphQL client
	client := NewGraphQLClient(config.Token)

	// Convert PodEnv to EnvironmentVariableInput
	var envVars []*EnvironmentVariableInput
	for _, env := range news.Env {
		envVars = append(envVars, &EnvironmentVariableInput{
			Key:   env.Key,
			Value: env.Value,
		})
	}

	// Call the generated SaveTemplate mutation with the existing ID for update
	templateId := olds.Template.Id
	result, err := client.GetClient().SaveTemplate(
		context.Background(),
		&templateId, // id - provide existing ID for update
		news.ContainerDiskInGb,
		&news.ContainerRegistryAuthId,
		news.DockerArgs,
		envVars,
		news.ImageName,
		&news.IsPublic,
		&news.IsServerless,
		news.Name,
		&news.Ports,
		&news.Readme,
		&news.StartJupyter,
		&news.StartSsh,
		news.VolumeInGb,
		&news.VolumeMountPath,
	)
	if err != nil {
		return state, fmt.Errorf("failed to update template: %w", err)
	}

	if result.SaveTemplate == nil {
		return state, fmt.Errorf("template update failed: no template returned")
	}

	// Convert result to Template struct
	template := Template{
		Id:        *result.SaveTemplate.ID,
		ImageName: *result.SaveTemplate.ImageName,
		Name:      *result.SaveTemplate.Name,
		// Set other fields from input since they're not returned by the mutation
		ContainerDiskInGb:       news.ContainerDiskInGb,
		ContainerRegistryAuthId: news.ContainerRegistryAuthId,
		DockerArgs:              news.DockerArgs,
		Env:                     news.Env,
		IsPublic:                news.IsPublic,
		IsServerless:            news.IsServerless,
		Ports:                   news.Ports,
		Readme:                  news.Readme,
		StartJupyter:            news.StartJupyter,
		StartSsh:                news.StartSsh,
		VolumeInGb:              news.VolumeInGb,
		VolumeMountPath:         news.VolumeMountPath,
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

	// Create GraphQL client
	client := NewGraphQLClient(config.Token)

	// Call the generated DeleteTemplate mutation
	_, err := client.GetClient().DeleteTemplate(
		context.Background(),
		props.Template.Name,
	)
	if err != nil {
		return fmt.Errorf("failed to delete template: %w", err)
	}

	return nil
}
