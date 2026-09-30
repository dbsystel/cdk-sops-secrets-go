package cdksopssecrets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/dbsystel/cdk-sops-secrets-go/cdksopssecrets/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/dbsystel/cdk-sops-secrets-go/cdksopssecrets/v2/internal"
)

type MultiStringParameter interface {
	constructs.Construct
	EncryptionKey() awskms.IKey
	Env() *interfaces.ResourceEnvironment
	KeyPrefix() *string
	KeySeparator() *string
	// The tree node.
	Node() constructs.Node
	Stack() awscdk.Stack
	Sync() SopsSync
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

// The jsii proxy struct for MultiStringParameter
type jsiiProxy_MultiStringParameter struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_MultiStringParameter) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultiStringParameter) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultiStringParameter) KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultiStringParameter) KeySeparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySeparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultiStringParameter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultiStringParameter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MultiStringParameter) Sync() SopsSync {
	var returns SopsSync
	_jsii_.Get(
		j,
		"sync",
		&returns,
	)
	return returns
}


func NewMultiStringParameter(scope constructs.Construct, id *string, props *MultiStringParameterProps) MultiStringParameter {
	_init_.Initialize()

	if err := validateNewMultiStringParameterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MultiStringParameter{}

	_jsii_.Create(
		"cdk-sops-secrets.MultiStringParameter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewMultiStringParameter_Override(m MultiStringParameter, scope constructs.Construct, id *string, props *MultiStringParameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-sops-secrets.MultiStringParameter",
		[]interface{}{scope, id, props},
		m,
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
func MultiStringParameter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMultiStringParameter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-sops-secrets.MultiStringParameter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MultiStringParameter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MultiStringParameter) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		m,
		"with",
		args,
		&returns,
	)

	return returns
}

