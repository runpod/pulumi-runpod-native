package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/fatih/structs"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Pod struct {
	AiApiId                 string      `pulumi:"aiApiId"`
	ApiKey                  string      `pulumi:"apiKey"`
	ConsumerUserId          string      `pulumi:"consumerUserId"`
	ContainerDiskInGb       int         `pulumi:"containerDiskInGb"`
	ContainerRegistryAuthId string      `pulumi:"containerRegistryAuthId"`
	CostMultiplier          float64     `pulumi:"costMultiplier"`
	CostPerHr               float64     `pulumi:"costPerHr"`
	CreatedAt               string      `pulumi:"createdAt"`
	AdjustedCostPerHr       float64     `pulumi:"adjustedCostPerHr"`
	DesiredStatus           string      `pulumi:"desiredStatus"`
	DockerArgs              string      `pulumi:"dockerArgs"`
	DockerId                string      `pulumi:"dockerId"`
	Env                     []string    `pulumi:"env,optional"`
	GpuCount                int         `pulumi:"gpuCount"`
	GpuPowerLimitPercent    int         `pulumi:"gpuPowerLimitPercent"`
	Gpus                    []Gpu       `pulumi:"gpus,optional"`
	Id                      string      `pulumi:"id"`
	ImageName               string      `pulumi:"imageName"`
	LastStatusChange        string      `pulumi:"lastStatusChange"`
	Locked                  bool        `pulumi:"locked"`
	MachineId               string      `pulumi:"machineId"`
	MemoryInGb              float64     `pulumi:"memoryInGb"`
	Name                    string      `pulumi:"name"`
	PodType                 string      `pulumi:"podType"`
	Port                    int         `pulumi:"port"`
	Ports                   string      `pulumi:"ports"`
	Registry                PodRegistry `pulumi:"registry"`
	TemplateId              string      `pulumi:"templateId"`
	UptimeSeconds           int         `pulumi:"uptimeSeconds"`
	VcpuCount               float64     `pulumi:"vcpuCount"`
	Version                 int         `pulumi:"version"`
	VolumeEncrypted         bool        `pulumi:"volumeEncrypted"`
	VolumeInGb              float64     `pulumi:"volumeInGb"`
	VolumeKey               string      `pulumi:"volumeKey"`
	VolumeMountPath         string      `pulumi:"volumeMountPath"`
	LastStartedAt           string      `pulumi:"lastStartedAt"`
}
type PodRegistry struct {
	Auth     string `pulumi:"auth"`
	Pass     string `pulumi:"pass"`
	Url      string `pulumi:"url"`
	User     string `pulumi:"user"`
	Username string `pulumi:"username"`
}
type Gpu struct {
	Id    string `pulumi:"id"`
	PodId string `pulumi:"podId"`
}
type PodArgs struct {
	AiApiId           string       `pulumi:"aiApiId,optional" structs:"aiApiId,omitempty"`
	CloudType         PodCloudType `pulumi:"cloudType,optional" structs:"cloudType,omitempty"`
	ContainerDiskInGb int          `pulumi:"containerDiskInGb,optional" structs:"containerDiskInGb,omitempty"`
	// ContainerRegistryAuthId string       `pulumi:"containerRegistryAuthId" structs:"containerRegistryAuthId"`
	CountryCode     string           `pulumi:"countryCode,optional" structs:"countryCode,omitempty"`
	CudaVersion     string           `pulumi:"cudaVersion,optional" structs:"cudaVersion,omitempty"`
	DataCenterId    string           `pulumi:"dataCenterId,optional" structs:"dataCenterId,omitempty"`
	DeployCost      float64          `pulumi:"deployCost,optional" structs:"deployCost,omitempty"`
	DockerArgs      string           `pulumi:"dockerArgs,optional" structs:"dockerArgs,omitempty"`
	Env             []PodEnv         `pulumi:"env,optional" structs:"env,omitempty"`
	GpuCount        int              `pulumi:"gpuCount" structs:"gpuCount,omitempty"`
	GpuTypeId       string           `pulumi:"gpuTypeId" structs:"gpuTypeId,omitempty"`
	GpuTypeIdList   []string         `pulumi:"gpuTypeIdList,optional" structs:"gpuTypeIdList,omitempty"`
	ImageName       string           `pulumi:"imageName" structs:"imageName,omitempty"`
	MinDisk         int              `pulumi:"minDisk,optional" structs:"minDisk,omitempty"`
	MinDownload     int              `pulumi:"minDownload,optional" structs:"minDownload,omitempty"`
	MinMemoryInGb   int              `pulumi:"minMemoryInGb,optional" structs:"minMemoryInGb,omitempty"`
	MinVcpuCount    int              `pulumi:"minVcpuCount,optional" structs:"minVcpuCount,omitempty"`
	MinUpload       int              `pulumi:"minUpload,optional" structs:"minUpload,omitempty"`
	Name            string           `pulumi:"name,optional" structs:"name,omitempty"`
	NetworkVolumeId string           `pulumi:"networkVolumeId,optional" structs:"networkVolumeId,omitempty"`
	Port            int              `pulumi:"port,optional" structs:"port,omitempty"`
	Ports           string           `pulumi:"ports,optional" structs:"ports,omitempty"`
	SavingsPlan     SavingsPlanInput `pulumi:"savingsPlan,optional" structs:"savingsPlan,omitempty"`
	StartJupyter    bool             `pulumi:"startJupyter,optional" structs:"startJupyter,omitempty"`
	StartSsh        bool             `pulumi:"startSsh,optional" structs:"startSsh,omitempty"`
	StopAfter       string           `pulumi:"stopAfter,optional" structs:"stopAfter,omitempty"`
	SupportPublicIp bool             `pulumi:"supportPublicIp,optional" structs:"supportPublicIp,omitempty"`
	TemplateId      string           `pulumi:"templateId,optional" structs:"templateId,omitempty"`
	TerminateAfter  string           `pulumi:"terminateAfter,optional" structs:"terminateAfter,omitempty"`
	VolumeInGb      int              `pulumi:"volumeInGb,optional" structs:"volumeInGb,omitempty"`
	VolumeKey       string           `pulumi:"volumeKey,optional" structs:"volumeKey,omitempty"`
	VolumeMountPath string           `pulumi:"volumeMountPath,optional" structs:"volumeMountPath,omitempty"`
}

type PodEnv struct {
	Key   string `pulumi:"key" structs:"key"`
	Value string `pulumi:"value" structs:"value"`
}

type SavingsPlanInput struct {
	PlanLength  string  `pulumi:"planLength" structs:"planLength"`
	UpfrontCost float64 `pulumi:"upfrontCost" structs:"upfrontCost"`
}

type PodCloudType string

const (
	ALL       PodCloudType = "ALL"
	SECURE    PodCloudType = "SECURE"
	COMMUNITY PodCloudType = "COMMUNITY"
)

type PodState struct {
	PodArgs
	Pod Pod `pulumi:"pod"`
}

type GqlInput struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}
type OutputDeployPod struct {
	Errors []struct {
		Message string
	}
	Data struct {
		PodFindAndDeployOnDemand Pod
	}
}
type OutputUpdatePod struct {
	Errors []struct {
		Message string
	}
	Data struct {
		PodEditJob Pod
	}
}

func (*Pod) Create(ctx p.Context, name string, input PodArgs, preview bool) (string, PodState, error) {
	state := PodState{PodArgs: input}
	if preview {
		return name, state, nil
	}
	config := infer.GetConfig[Config](ctx)

	gqlVariable := structs.Map(input)
	gqlInput := GqlInput{
		Query: `
		mutation DeployMutation (			
			$aiApiId: String
			$cloudType: CloudTypeEnum
			$containerDiskInGb: Int
			$countryCode: String
			$deployCost: Float
			$dockerArgs: String
			$env: [EnvironmentVariableInput]
			$gpuCount: Int
			$gpuTypeId: String
			$gpuTypeIdList: [String]
			$imageName: String
			$minDisk: Int
			$minDownload: Int
			$minMemoryInGb: Int
			$minUpload: Int
			$minVcpuCount: Int
			$name: String
			$networkVolumeId: String
			$port: Port
			$ports: String
			$startJupyter: Boolean
			$startSsh: Boolean
			$stopAfter: DateTime
			$supportPublicIp: Boolean
			$templateId: String
			$terminateAfter: DateTime
			$volumeInGb: Int
			$volumeKey: String
			$volumeMountPath: String
			$dataCenterId: String
			$savingsPlan: SavingsPlanInput
			$cudaVersion: String
		) {
			podFindAndDeployOnDemand(input: {				
				aiApiId: $aiApiId,
				cloudType: $cloudType,
				containerDiskInGb: $containerDiskInGb,
				countryCode: $countryCode,
				deployCost: $deployCost,
				dockerArgs: $dockerArgs,
				env: $env,
				gpuCount: $gpuCount,
				gpuTypeId: $gpuTypeId,
				gpuTypeIdList: $gpuTypeIdList,
				imageName: $imageName,
				minDisk: $minDisk,
				minDownload: $minDownload,
				minMemoryInGb: $minMemoryInGb,
				minUpload: $minUpload,
				minVcpuCount: $minVcpuCount,
				name: $name,
				networkVolumeId: $networkVolumeId,
				port: $port,
				ports: $ports,
				startJupyter: $startJupyter,
				startSsh: $startSsh,
				stopAfter: $stopAfter,
				supportPublicIp: $supportPublicIp,
				templateId: $templateId,
				terminateAfter: $terminateAfter,
				volumeInGb: $volumeInGb,
				volumeKey: $volumeKey,
				volumeMountPath: $volumeMountPath,
				dataCenterId: $dataCenterId,
				savingsPlan: $savingsPlan,
				cudaVersion: $cudaVersion,
			}) {
				id
    			imageName    			
    			machineId
				containerDiskInGb				
			}
		}`,
		Variables: gqlVariable,
	}

	jsonValue, err := json.Marshal(gqlInput)
	if err != nil {
		return name, state, err
	}

	url := "https://api.runpod.dev/graphql?api_key=" + config.Token

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

	output := &OutputDeployPod{}
	err = json.Unmarshal(data, output)
	if err != nil {
		return name, state, err
	}

	if len(output.Errors) > 0 {
		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
		return name, state, err
	}

	pod := output.Data.PodFindAndDeployOnDemand
	if pod.Id == "" {
		err = fmt.Errorf("graphql pod is nil: %s", string(data))
		return name, state, err
	}

	state.Pod = pod

	return name, state, nil
}

func (*Pod) Update(ctx p.Context, id string, olds PodState, news PodArgs, preview bool) (PodState, error) {
	state := PodState{PodArgs: news}
	if preview {
		return state, nil
	}
	config := infer.GetConfig[Config](ctx)

	gqlVariable := structs.Map(news)
	gqlVariable["podId"] = olds.Pod.Id
	gqlInput := GqlInput{
		Query: `
		mutation UpdatePodMutation (
			$podId: String!
			$dockerArgs: String
			$imageName: String!
			$env: [EnvironmentVariableInput]
			$port: Port
			$ports: String
			$containerDiskInGb: Int!
			$volumeInGb: Int
			$volumeMountPath: String
			$containerRegistryAuthId: String
		) {
			podEditJob(input: {				
				podId: $podId,
				dockerArgs: $dockerArgs,
				imageName: $imageName,
				env: $env,
				port: $port,
				ports: $ports,
				containerDiskInGb: $containerDiskInGb,
				volumeInGb: $volumeInGb,
				volumeMountPath: $volumeMountPath,
				containerRegistryAuthId: $containerRegistryAuthId,
			}) {
				id
    			imageName    			
    			machineId
				containerDiskInGb								
			}
		}`,
		Variables: gqlVariable,
	}

	jsonValue, err := json.Marshal(gqlInput)
	if err != nil {
		return state, err
	}

	url := "https://api.runpod.dev/graphql?api_key=" + config.Token

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

	output := &OutputUpdatePod{}
	err = json.Unmarshal(data, output)
	if err != nil {
		return state, err
	}

	if len(output.Errors) > 0 {
		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
		return state, err
	}

	pod := output.Data.PodEditJob
	if pod.Id == "" {
		err = fmt.Errorf("graphql pod is nil: %s", string(data))
		return state, err
	}

	state.Pod = pod

	return state, nil
}

func (*Pod) Diff(ctx p.Context, id string, olds PodState, news PodArgs) (p.DiffResponse, error) {

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
	// if news.ContainerRegistryAuthId != olds.ContainerRegistryAuthId {
	// 	diff["containerRegistryAuthId"] = p.PropertyDiff{Kind: p.Update}
	// }
	if news.DockerArgs != olds.DockerArgs {
		diff["dockerArgs"] = p.PropertyDiff{Kind: p.Update}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*Pod) Delete(ctx p.Context, id string, props PodState) error {
	config := infer.GetConfig[Config](ctx)
	gqlVariable := map[string]interface{}{"podId": props.Pod.Id}

	gqlInput := GqlInput{
		Query: `
		mutation podTerminateMutation (
			$podId: String!	
		) {
			podTerminate(input: {				
				podId: $podId				
			})
		}`,
		Variables: gqlVariable,
	}

	jsonValue, err := json.Marshal(gqlInput)
	if err != nil {
		return err
	}

	url := "https://api.runpod.dev/graphql?api_key=" + config.Token

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
