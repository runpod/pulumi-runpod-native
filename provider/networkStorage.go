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

type NetworkStorage struct {
	Id           string     `pulumi:"id"`
	Name         string     `pulumi:"name"`
	Size         int        `pulumi:"size"`
	DataCenterId string     `pulumi:"dataCenterId"`
	DataCenter   DataCenter `pulumi:"dataCenter"`
}

type DataCenter struct {
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	Location       string `pulumi:"location"`
	StorageSupport bool   `pulumi:"storageSupport"`
}

type NetworkStorageArgs struct {
	Name         string `pulumi:"name" structs:"name"`
	Size         int    `pulumi:"size" structs:"size"`
	DataCenterId string `pulumi:"dataCenterId" structs:"dataCenterId"`
}

type NetworkStorageState struct {
	NetworkStorageArgs
	NetworkStorage NetworkStorage `pulumi:"networkStorage"`
}

type CreateNetworkStorageOutput struct {
	Errors []struct {
		Message string
	}
	Data struct {
		CreateNetworkVolume NetworkStorage
	}
}

type UpdateNetworkStorageOutput struct {
	Errors []struct {
		Message string
	}
	Data struct {
		UpdateNetworkVolume NetworkStorage
	}
}

func (*NetworkStorage) Create(ctx p.Context, name string, input NetworkStorageArgs, preview bool) (string, NetworkStorageState, error) {
	state := NetworkStorageState{NetworkStorageArgs: input}
	if preview {
		return name, state, nil
	}

	gqlVariable := structs.Map(input)
	gqlInput := GqlInput{
		Query: `
		mutation NetworkVolumeCreate (			
			$name: String
			$size: Int
			$dataCenterId: String
		) {
			createNetworkVolume(input: {				
				name: $name,
				size: $size,
				dataCenterId: $dataCenterId,
			}) {
				id
				name
				size
				dataCenterId
				dataCenter {
					id
					name
					location
					storageSupport
				}	
			}
		}`,
		Variables: gqlVariable,
	}

	jsonValue, err := json.Marshal(gqlInput)
	if err != nil {
		return name, state, err
	}

	config := infer.GetConfig[Config](ctx)
	url := "https://api.runpod.dev/graphql?api_key=" + config.Token

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

	output := &CreateNetworkStorageOutput{}
	err = json.Unmarshal(data, output)
	if err != nil {
		return name, state, err
	}

	if len(output.Errors) > 0 {
		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
		return name, state, err
	}

	ns := output.Data.CreateNetworkVolume
	if ns.Id == "" {
		err = fmt.Errorf("graphql nw is nil: %s", string(data))
		return name, state, err
	}

	state.NetworkStorage = ns

	return name, state, nil
}

func (*NetworkStorage) Update(ctx p.Context, id string, olds NetworkStorageState, news NetworkStorageArgs, preview bool) (NetworkStorageState, error) {
	state := NetworkStorageState{NetworkStorageArgs: news}
	if preview {
		return state, nil
	}
	config := infer.GetConfig[Config](ctx)

	gqlVariable := structs.Map(news)
	gqlVariable["id"] = olds.NetworkStorage.Id
	gqlInput := GqlInput{
		Query: `
		mutation updateNetworkVolume (			
			$id: String!
			$name: String
			$size: Int
		) {
			updateNetworkVolume(input: {				
				id: $id,
				name: $name,
				size: $size,
			}) {
				id
				name
				size
				dataCenterId
				dataCenter {
					id
					name
					location
					storageSupport
				}	
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

	output := &UpdateNetworkStorageOutput{}
	err = json.Unmarshal(data, output)
	if err != nil {
		return state, err
	}

	if len(output.Errors) > 0 {
		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
		return state, err
	}

	ns := output.Data.UpdateNetworkVolume
	if ns.Id == "" {
		err = fmt.Errorf("graphql ns is nil: %s", string(data))
		return state, err
	}

	state.NetworkStorage = ns

	return state, nil
}

func (*NetworkStorage) Diff(ctx p.Context, id string, olds NetworkStorageState, news NetworkStorageArgs) (p.DiffResponse, error) {

	diff := map[string]p.PropertyDiff{}

	if news.Name != olds.Name {
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.Size != olds.Size {
		diff["size"] = p.PropertyDiff{Kind: p.Update}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*NetworkStorage) Delete(ctx p.Context, id string, props NetworkStorageState) error {
	config := infer.GetConfig[Config](ctx)
	gqlVariable := map[string]interface{}{"id": props.NetworkStorage.Id}

	gqlInput := GqlInput{
		Query: `
		mutation deleteNetworkVolume (
			$id: String!	
		) {
			deleteNetworkVolume(input: {				
				id: $id				
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
		err = fmt.Errorf("graphql err: %s", output.Errors[0].Message)
		return err
	}

	return nil

}
