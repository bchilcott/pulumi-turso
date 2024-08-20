// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package turso

import (
	"context"
	"reflect"

	"errors"
	"github.com/bchilcott/pulumi-turso/sdk/go/turso/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Database Token resource
type DatabaseToken struct {
	pulumi.CustomResourceState

	// Authorization level for the token (full-access or read-only).
	Authorization pulumi.StringPtrOutput `pulumi:"authorization"`
	// The name of the database.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// Expiration time for the token (e.g., 2w1d30m).
	Expiration pulumi.StringPtrOutput `pulumi:"expiration"`
	// The generated authorization token (JWT).
	Jwt pulumi.StringOutput `pulumi:"jwt"`
	// The name of the organization or user.
	OrganizationName pulumi.StringOutput `pulumi:"organizationName"`
}

// NewDatabaseToken registers a new resource with the given unique name, arguments, and options.
func NewDatabaseToken(ctx *pulumi.Context,
	name string, args *DatabaseTokenArgs, opts ...pulumi.ResourceOption) (*DatabaseToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.OrganizationName == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"jwt",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabaseToken
	err := ctx.RegisterResource("turso:index/databaseToken:DatabaseToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseToken gets an existing DatabaseToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseTokenState, opts ...pulumi.ResourceOption) (*DatabaseToken, error) {
	var resource DatabaseToken
	err := ctx.ReadResource("turso:index/databaseToken:DatabaseToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseToken resources.
type databaseTokenState struct {
	// Authorization level for the token (full-access or read-only).
	Authorization *string `pulumi:"authorization"`
	// The name of the database.
	DatabaseName *string `pulumi:"databaseName"`
	// Expiration time for the token (e.g., 2w1d30m).
	Expiration *string `pulumi:"expiration"`
	// The generated authorization token (JWT).
	Jwt *string `pulumi:"jwt"`
	// The name of the organization or user.
	OrganizationName *string `pulumi:"organizationName"`
}

type DatabaseTokenState struct {
	// Authorization level for the token (full-access or read-only).
	Authorization pulumi.StringPtrInput
	// The name of the database.
	DatabaseName pulumi.StringPtrInput
	// Expiration time for the token (e.g., 2w1d30m).
	Expiration pulumi.StringPtrInput
	// The generated authorization token (JWT).
	Jwt pulumi.StringPtrInput
	// The name of the organization or user.
	OrganizationName pulumi.StringPtrInput
}

func (DatabaseTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseTokenState)(nil)).Elem()
}

type databaseTokenArgs struct {
	// Authorization level for the token (full-access or read-only).
	Authorization *string `pulumi:"authorization"`
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// Expiration time for the token (e.g., 2w1d30m).
	Expiration *string `pulumi:"expiration"`
	// The name of the organization or user.
	OrganizationName string `pulumi:"organizationName"`
}

// The set of arguments for constructing a DatabaseToken resource.
type DatabaseTokenArgs struct {
	// Authorization level for the token (full-access or read-only).
	Authorization pulumi.StringPtrInput
	// The name of the database.
	DatabaseName pulumi.StringInput
	// Expiration time for the token (e.g., 2w1d30m).
	Expiration pulumi.StringPtrInput
	// The name of the organization or user.
	OrganizationName pulumi.StringInput
}

func (DatabaseTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseTokenArgs)(nil)).Elem()
}

type DatabaseTokenInput interface {
	pulumi.Input

	ToDatabaseTokenOutput() DatabaseTokenOutput
	ToDatabaseTokenOutputWithContext(ctx context.Context) DatabaseTokenOutput
}

func (*DatabaseToken) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseToken)(nil)).Elem()
}

func (i *DatabaseToken) ToDatabaseTokenOutput() DatabaseTokenOutput {
	return i.ToDatabaseTokenOutputWithContext(context.Background())
}

func (i *DatabaseToken) ToDatabaseTokenOutputWithContext(ctx context.Context) DatabaseTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseTokenOutput)
}

// DatabaseTokenArrayInput is an input type that accepts DatabaseTokenArray and DatabaseTokenArrayOutput values.
// You can construct a concrete instance of `DatabaseTokenArrayInput` via:
//
//	DatabaseTokenArray{ DatabaseTokenArgs{...} }
type DatabaseTokenArrayInput interface {
	pulumi.Input

	ToDatabaseTokenArrayOutput() DatabaseTokenArrayOutput
	ToDatabaseTokenArrayOutputWithContext(context.Context) DatabaseTokenArrayOutput
}

type DatabaseTokenArray []DatabaseTokenInput

func (DatabaseTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseToken)(nil)).Elem()
}

func (i DatabaseTokenArray) ToDatabaseTokenArrayOutput() DatabaseTokenArrayOutput {
	return i.ToDatabaseTokenArrayOutputWithContext(context.Background())
}

func (i DatabaseTokenArray) ToDatabaseTokenArrayOutputWithContext(ctx context.Context) DatabaseTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseTokenArrayOutput)
}

// DatabaseTokenMapInput is an input type that accepts DatabaseTokenMap and DatabaseTokenMapOutput values.
// You can construct a concrete instance of `DatabaseTokenMapInput` via:
//
//	DatabaseTokenMap{ "key": DatabaseTokenArgs{...} }
type DatabaseTokenMapInput interface {
	pulumi.Input

	ToDatabaseTokenMapOutput() DatabaseTokenMapOutput
	ToDatabaseTokenMapOutputWithContext(context.Context) DatabaseTokenMapOutput
}

type DatabaseTokenMap map[string]DatabaseTokenInput

func (DatabaseTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseToken)(nil)).Elem()
}

func (i DatabaseTokenMap) ToDatabaseTokenMapOutput() DatabaseTokenMapOutput {
	return i.ToDatabaseTokenMapOutputWithContext(context.Background())
}

func (i DatabaseTokenMap) ToDatabaseTokenMapOutputWithContext(ctx context.Context) DatabaseTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseTokenMapOutput)
}

type DatabaseTokenOutput struct{ *pulumi.OutputState }

func (DatabaseTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseToken)(nil)).Elem()
}

func (o DatabaseTokenOutput) ToDatabaseTokenOutput() DatabaseTokenOutput {
	return o
}

func (o DatabaseTokenOutput) ToDatabaseTokenOutputWithContext(ctx context.Context) DatabaseTokenOutput {
	return o
}

// Authorization level for the token (full-access or read-only).
func (o DatabaseTokenOutput) Authorization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseToken) pulumi.StringPtrOutput { return v.Authorization }).(pulumi.StringPtrOutput)
}

// The name of the database.
func (o DatabaseTokenOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseToken) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// Expiration time for the token (e.g., 2w1d30m).
func (o DatabaseTokenOutput) Expiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseToken) pulumi.StringPtrOutput { return v.Expiration }).(pulumi.StringPtrOutput)
}

// The generated authorization token (JWT).
func (o DatabaseTokenOutput) Jwt() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseToken) pulumi.StringOutput { return v.Jwt }).(pulumi.StringOutput)
}

// The name of the organization or user.
func (o DatabaseTokenOutput) OrganizationName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseToken) pulumi.StringOutput { return v.OrganizationName }).(pulumi.StringOutput)
}

type DatabaseTokenArrayOutput struct{ *pulumi.OutputState }

func (DatabaseTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseToken)(nil)).Elem()
}

func (o DatabaseTokenArrayOutput) ToDatabaseTokenArrayOutput() DatabaseTokenArrayOutput {
	return o
}

func (o DatabaseTokenArrayOutput) ToDatabaseTokenArrayOutputWithContext(ctx context.Context) DatabaseTokenArrayOutput {
	return o
}

func (o DatabaseTokenArrayOutput) Index(i pulumi.IntInput) DatabaseTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseToken {
		return vs[0].([]*DatabaseToken)[vs[1].(int)]
	}).(DatabaseTokenOutput)
}

type DatabaseTokenMapOutput struct{ *pulumi.OutputState }

func (DatabaseTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseToken)(nil)).Elem()
}

func (o DatabaseTokenMapOutput) ToDatabaseTokenMapOutput() DatabaseTokenMapOutput {
	return o
}

func (o DatabaseTokenMapOutput) ToDatabaseTokenMapOutputWithContext(ctx context.Context) DatabaseTokenMapOutput {
	return o
}

func (o DatabaseTokenMapOutput) MapIndex(k pulumi.StringInput) DatabaseTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseToken {
		return vs[0].(map[string]*DatabaseToken)[vs[1].(string)]
	}).(DatabaseTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseTokenInput)(nil)).Elem(), &DatabaseToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseTokenArrayInput)(nil)).Elem(), DatabaseTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseTokenMapInput)(nil)).Elem(), DatabaseTokenMap{})
	pulumi.RegisterOutputType(DatabaseTokenOutput{})
	pulumi.RegisterOutputType(DatabaseTokenArrayOutput{})
	pulumi.RegisterOutputType(DatabaseTokenMapOutput{})
}