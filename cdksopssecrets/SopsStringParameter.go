package cdksopssecrets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/dbsystel/cdk-sops-secrets-go/cdksopssecrets/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsssm"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/dbsystel/cdk-sops-secrets-go/cdksopssecrets/v2/internal"
)

// A drop in replacement for the normal String Parameter, that is populated with the encrypted content of the given sops file.
type SopsStringParameter interface {
	constructs.Construct
	awsssm.IStringParameter
	EncryptionKey() awskms.IKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed in a Stack (those created by
	// creating new class instances like `new Role()`, `new Bucket()`, etc.), this
	// is always the same as the environment of the stack they belong to.
	//
	// For referenced resources (those obtained from referencing methods like
	// `Role.fromRoleArn()`, `Bucket.fromBucketName()`, etc.), they might be
	// different than the stack they were imported into.
	Env() *interfaces.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// The ARN of the SSM Parameter resource.
	ParameterArn() *string
	// The name of the SSM Parameter resource.
	ParameterName() *string
	// A reference to a Parameter resource.
	ParameterRef() *interfacesawsssm.ParameterReference
	// The type of the SSM Parameter resource.
	ParameterType() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The parameter value.
	//
	// Value must not nest another parameter. Do not use {{}} in the value.
	StringValue() *string
	Sync() SopsSync
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Grants read (DescribeParameter, GetParameters, GetParameter, GetParameterHistory) permissions on the SSM Parameter.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grants write (PutParameter) permissions on the SSM Parameter.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for SopsStringParameter
type jsiiProxy_SopsStringParameter struct {
	internal.Type__constructsConstruct
	internal.Type__awsssmIStringParameter
}

func (j *jsiiProxy_SopsStringParameter) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsStringParameter) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsStringParameter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsStringParameter) ParameterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsStringParameter) ParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsStringParameter) ParameterRef() *interfacesawsssm.ParameterReference {
	var returns *interfacesawsssm.ParameterReference
	_jsii_.Get(
		j,
		"parameterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsStringParameter) ParameterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsStringParameter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsStringParameter) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsStringParameter) Sync() SopsSync {
	var returns SopsSync
	_jsii_.Get(
		j,
		"sync",
		&returns,
	)
	return returns
}


func NewSopsStringParameter(scope constructs.Construct, id *string, props *SopsStringParameterProps) SopsStringParameter {
	_init_.Initialize()

	if err := validateNewSopsStringParameterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SopsStringParameter{}

	_jsii_.Create(
		"cdk-sops-secrets.SopsStringParameter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSopsStringParameter_Override(s SopsStringParameter, scope constructs.Construct, id *string, props *SopsStringParameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-sops-secrets.SopsStringParameter",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SopsStringParameter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSopsStringParameter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-sops-secrets.SopsStringParameter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SopsStringParameter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SopsStringParameter) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SopsStringParameter) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SopsStringParameter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SopsStringParameter) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

