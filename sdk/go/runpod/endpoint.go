// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package runpod

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/runpod/pulumi-runpod-native/sdk/go/runpod/internal"
)

type Endpoint struct {
	pulumi.CustomResourceState

	Endpoint        EndpointTypeOutput     `pulumi:"endpoint"`
	GpuIds          pulumi.StringOutput    `pulumi:"gpuIds"`
	IdleTimeout     pulumi.IntPtrOutput    `pulumi:"idleTimeout"`
	Locations       pulumi.StringPtrOutput `pulumi:"locations"`
	Name            pulumi.StringOutput    `pulumi:"name"`
	NetworkVolumeId pulumi.StringPtrOutput `pulumi:"networkVolumeId"`
	ScalerType      pulumi.StringPtrOutput `pulumi:"scalerType"`
	ScalerValue     pulumi.IntPtrOutput    `pulumi:"scalerValue"`
	TemplateId      pulumi.StringOutput    `pulumi:"templateId"`
	WorkersMax      pulumi.IntPtrOutput    `pulumi:"workersMax"`
	WorkersMin      pulumi.IntPtrOutput    `pulumi:"workersMin"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GpuIds == nil {
		return nil, errors.New("invalid value for required argument 'GpuIds'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Endpoint
	err := ctx.RegisterResource("runpod:index:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("runpod:index:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
}

type EndpointState struct {
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	GpuIds          string  `pulumi:"gpuIds"`
	IdleTimeout     *int    `pulumi:"idleTimeout"`
	Locations       *string `pulumi:"locations"`
	Name            string  `pulumi:"name"`
	NetworkVolumeId *string `pulumi:"networkVolumeId"`
	ScalerType      *string `pulumi:"scalerType"`
	ScalerValue     *int    `pulumi:"scalerValue"`
	TemplateId      string  `pulumi:"templateId"`
	WorkersMax      *int    `pulumi:"workersMax"`
	WorkersMin      *int    `pulumi:"workersMin"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	GpuIds          pulumi.StringInput
	IdleTimeout     pulumi.IntPtrInput
	Locations       pulumi.StringPtrInput
	Name            pulumi.StringInput
	NetworkVolumeId pulumi.StringPtrInput
	ScalerType      pulumi.StringPtrInput
	ScalerValue     pulumi.IntPtrInput
	TemplateId      pulumi.StringInput
	WorkersMax      pulumi.IntPtrInput
	WorkersMin      pulumi.IntPtrInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

// EndpointArrayInput is an input type that accepts EndpointArray and EndpointArrayOutput values.
// You can construct a concrete instance of `EndpointArrayInput` via:
//
//	EndpointArray{ EndpointArgs{...} }
type EndpointArrayInput interface {
	pulumi.Input

	ToEndpointArrayOutput() EndpointArrayOutput
	ToEndpointArrayOutputWithContext(context.Context) EndpointArrayOutput
}

type EndpointArray []EndpointInput

func (EndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (i EndpointArray) ToEndpointArrayOutput() EndpointArrayOutput {
	return i.ToEndpointArrayOutputWithContext(context.Background())
}

func (i EndpointArray) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointArrayOutput)
}

// EndpointMapInput is an input type that accepts EndpointMap and EndpointMapOutput values.
// You can construct a concrete instance of `EndpointMapInput` via:
//
//	EndpointMap{ "key": EndpointArgs{...} }
type EndpointMapInput interface {
	pulumi.Input

	ToEndpointMapOutput() EndpointMapOutput
	ToEndpointMapOutputWithContext(context.Context) EndpointMapOutput
}

type EndpointMap map[string]EndpointInput

func (EndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (i EndpointMap) ToEndpointMapOutput() EndpointMapOutput {
	return i.ToEndpointMapOutputWithContext(context.Background())
}

func (i EndpointMap) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMapOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

func (o EndpointOutput) Endpoint() EndpointTypeOutput {
	return o.ApplyT(func(v *Endpoint) EndpointTypeOutput { return v.Endpoint }).(EndpointTypeOutput)
}

func (o EndpointOutput) GpuIds() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.GpuIds }).(pulumi.StringOutput)
}

func (o EndpointOutput) IdleTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.IntPtrOutput { return v.IdleTimeout }).(pulumi.IntPtrOutput)
}

func (o EndpointOutput) Locations() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.Locations }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EndpointOutput) NetworkVolumeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.NetworkVolumeId }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) ScalerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.ScalerType }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) ScalerValue() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.IntPtrOutput { return v.ScalerValue }).(pulumi.IntPtrOutput)
}

func (o EndpointOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

func (o EndpointOutput) WorkersMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.IntPtrOutput { return v.WorkersMax }).(pulumi.IntPtrOutput)
}

func (o EndpointOutput) WorkersMin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.IntPtrOutput { return v.WorkersMin }).(pulumi.IntPtrOutput)
}

type EndpointArrayOutput struct{ *pulumi.OutputState }

func (EndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (o EndpointArrayOutput) ToEndpointArrayOutput() EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) Index(i pulumi.IntInput) EndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].([]*Endpoint)[vs[1].(int)]
	}).(EndpointOutput)
}

type EndpointMapOutput struct{ *pulumi.OutputState }

func (EndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (o EndpointMapOutput) ToEndpointMapOutput() EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) MapIndex(k pulumi.StringInput) EndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].(map[string]*Endpoint)[vs[1].(string)]
	}).(EndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointInput)(nil)).Elem(), &Endpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointArrayInput)(nil)).Elem(), EndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointMapInput)(nil)).Elem(), EndpointMap{})
	pulumi.RegisterOutputType(EndpointOutput{})
	pulumi.RegisterOutputType(EndpointArrayOutput{})
	pulumi.RegisterOutputType(EndpointMapOutput{})
}
