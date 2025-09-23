package provider

import (
	"context"
	"fmt"
	"reflect"

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
	CountryCode     string            `pulumi:"countryCode,optional" structs:"countryCode,omitempty"`
	CudaVersion     string            `pulumi:"cudaVersion,optional" structs:"cudaVersion,omitempty"`
	DataCenterId    string            `pulumi:"dataCenterId,optional" structs:"dataCenterId,omitempty"`
	DeployCost      float64           `pulumi:"deployCost,optional" structs:"deployCost,omitempty"`
	DockerArgs      string            `pulumi:"dockerArgs,optional" structs:"dockerArgs,omitempty"`
	Env             []PodEnv          `pulumi:"env,optional" structs:"env,omitempty"`
	GpuCount        int               `pulumi:"gpuCount" structs:"gpuCount,omitempty"`
	GpuTypeId       string            `pulumi:"gpuTypeId" structs:"gpuTypeId,omitempty"`
	GpuTypeIdList   []string          `pulumi:"gpuTypeIdList,optional" structs:"gpuTypeIdList,omitempty"`
	ImageName       string            `pulumi:"imageName" structs:"imageName,omitempty"`
	MinDisk         int               `pulumi:"minDisk,optional" structs:"minDisk,omitempty"`
	MinDownload     int               `pulumi:"minDownload,optional" structs:"minDownload,omitempty"`
	MinMemoryInGb   int               `pulumi:"minMemoryInGb,optional" structs:"minMemoryInGb,omitempty"`
	MinVcpuCount    int               `pulumi:"minVcpuCount,optional" structs:"minVcpuCount,omitempty"`
	MinUpload       int               `pulumi:"minUpload,optional" structs:"minUpload,omitempty"`
	Name            string            `pulumi:"name,optional" structs:"name,omitempty"`
	NetworkVolumeId string            `pulumi:"networkVolumeId,optional" structs:"networkVolumeId,omitempty"`
	Port            int               `pulumi:"port,optional" structs:"port,omitempty"`
	Ports           string            `pulumi:"ports,optional" structs:"ports,omitempty"`
	SavingsPlan     *SavingsPlanInput `pulumi:"savingsPlan,optional" structs:"savingsPlan,omitempty"`
	StartJupyter    bool              `pulumi:"startJupyter,optional" structs:"startJupyter,omitempty"`
	StartSsh        bool              `pulumi:"startSsh,optional" structs:"startSsh,omitempty"`
	StopAfter       string            `pulumi:"stopAfter,optional" structs:"stopAfter,omitempty"`
	SupportPublicIp bool              `pulumi:"supportPublicIp,optional" structs:"supportPublicIp,omitempty"`
	TemplateId      string            `pulumi:"templateId,optional" structs:"templateId,omitempty"`
	TerminateAfter  string            `pulumi:"terminateAfter,optional" structs:"terminateAfter,omitempty"`
	VolumeInGb      int               `pulumi:"volumeInGb,optional" structs:"volumeInGb,omitempty"`
	VolumeKey       string            `pulumi:"volumeKey,optional" structs:"volumeKey,omitempty"`
	VolumeMountPath string            `pulumi:"volumeMountPath,optional" structs:"volumeMountPath,omitempty"`
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

	// Create GraphQL client using generated types
	gqlClient := NewGraphQLClient(config.Token)

	// Convert string environment variables to GraphQL input format
	var envVars []*EnvironmentVariableInput
	for _, env := range input.Env {
		envVars = append(envVars, &EnvironmentVariableInput{
			Key:   env.Key,
			Value: env.Value,
		})
	}

	// Convert inputs to generated types
	var cloudType *CloudTypeEnum
	if input.CloudType != "" {
		ct := CloudTypeEnum(input.CloudType)
		cloudType = &ct
	}

	// Call the generated DeployMutation function
	response, err := gqlClient.GetClient().DeployMutation(
		context.Background(),
		stringPtr(input.AiApiId),              // aiAPIID
		cloudType,                             // cloudType
		intPtr(input.ContainerDiskInGb),       // containerDiskInGb
		stringPtr(input.CountryCode),          // countryCode
		float64Ptr(input.DeployCost),          // deployCost
		stringPtr(input.DockerArgs),           // dockerArgs
		envVars,                               // env
		intPtr(input.GpuCount),                // gpuCount
		stringPtr(input.GpuTypeId),            // gpuTypeID
		stringSlicePtr(input.GpuTypeIdList),   // gpuTypeIDList
		stringPtr(input.ImageName),            // imageName
		intPtr(input.MinDisk),                 // minDisk
		intPtr(input.MinDownload),             // minDownload
		intPtr(input.MinMemoryInGb),           // minMemoryInGb
		intPtr(input.MinUpload),               // minUpload
		intPtr(input.MinVcpuCount),            // minVcpuCount
		stringPtr(name),                       // name
		stringPtr(input.NetworkVolumeId),      // networkVolumeID
		intPtrToStringPtr(input.Port),         // port
		stringPtr(input.Ports),                // ports
		boolPtr(input.StartJupyter),           // startJupyter
		boolPtr(input.StartSsh),               // startSSH
		stringPtr(input.StopAfter),            // stopAfter
		boolPtr(input.SupportPublicIp),        // supportPublicIP
		stringPtr(input.TemplateId),           // templateID
		stringPtr(input.TerminateAfter),       // terminateAfter
		intPtr(input.VolumeInGb),              // volumeInGb
		stringPtr(input.VolumeKey),            // volumeKey
		stringPtr(input.VolumeMountPath),      // volumeMountPath
		stringPtr(input.DataCenterId),         // dataCenterID
		nil,                                   // savingsPlan - TODO: convert this type
		stringPtr(input.CudaVersion),          // cudaVersion
	)

	if err != nil {
		return name, state, fmt.Errorf("deploy mutation failed: %v", err)
	}

	// Extract the pod from the response
	if response.GetPodFindAndDeployOnDemand() == nil {
		return name, state, fmt.Errorf("no pod returned from deploy mutation")
	}

	podData := response.GetPodFindAndDeployOnDemand()

	// Convert generated type back to our Pod struct
	pod := Pod{
		Id:                podData.GetID(),
		ImageName:         *podData.GetImageName(),
		MachineId:         podData.GetMachineID(),
		ContainerDiskInGb: *podData.GetContainerDiskInGb(),
	}

	state.Pod = pod
	return name, state, nil
}

// Helper functions for pointer conversions
func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func intPtr(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func intPtrToStringPtr(i int) *string {
	if i == 0 {
		return nil
	}
	s := fmt.Sprintf("%d", i)
	return &s
}

func float64Ptr(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}

func boolPtr(b bool) *bool {
	return &b
}

func stringSlicePtr(s []string) []*string {
	if len(s) == 0 {
		return nil
	}
	result := make([]*string, len(s))
	for i, str := range s {
		result[i] = &str
	}
	return result
}

func (*Pod) Update(ctx p.Context, id string, olds PodState, news PodArgs, preview bool) (PodState, error) {
	state := PodState{PodArgs: news}
	if preview {
		return state, nil
	}
	config := infer.GetConfig[Config](ctx)

	// Create GraphQL client using generated types
	gqlClient := NewGraphQLClient(config.Token)

	// Convert string environment variables to GraphQL input format
	var envVars []*EnvironmentVariableInput
	for _, env := range news.Env {
		envVars = append(envVars, &EnvironmentVariableInput{
			Key:   env.Key,
			Value: env.Value,
		})
	}

	// Call the generated UpdatePodMutation function
	response, err := gqlClient.GetClient().UpdatePodMutation(
		context.Background(),
		olds.Pod.Id,                                    // podID (required)
		stringPtr(news.DockerArgs),                     // dockerArgs
		news.ImageName,                                 // imageName (required)
		envVars,                                        // env
		intPtrToStringPtr(news.Port),                   // port
		stringPtr(news.Ports),                          // ports
		news.ContainerDiskInGb,                         // containerDiskInGb (required)
		intPtr(news.VolumeInGb),                        // volumeInGb
		stringPtr(news.VolumeMountPath),                // volumeMountPath
		nil,                                            // containerRegistryAuthId
	)

	if err != nil {
		return state, fmt.Errorf("update pod mutation failed: %v", err)
	}

	// Extract the pod from the response
	if response.GetPodEditJob() == nil {
		return state, fmt.Errorf("no pod returned from update mutation")
	}

	podData := response.GetPodEditJob()

	// Convert generated type back to our Pod struct
	pod := Pod{
		Id:                podData.GetID(),
		ImageName:         *podData.GetImageName(),
		MachineId:         podData.GetMachineID(),
		ContainerDiskInGb: *podData.GetContainerDiskInGb(),
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

	// Create GraphQL client using generated types
	gqlClient := NewGraphQLClient(config.Token)

	// Call the generated PodTerminateMutation function
	_, err := gqlClient.GetClient().PodTerminateMutation(
		context.Background(),
		props.Pod.Id, // podID (required)
	)

	if err != nil {
		return fmt.Errorf("pod terminate mutation failed: %v", err)
	}

	return nil
}
