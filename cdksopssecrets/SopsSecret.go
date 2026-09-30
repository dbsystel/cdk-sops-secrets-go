package cdksopssecrets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/dbsystel/cdk-sops-secrets-go/cdksopssecrets/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/dbsystel/cdk-sops-secrets-go/cdksopssecrets/v2/internal"
)

// A drop in replacement for the normal Secret, that is populated with the encrypted content of the given sops file.
type SopsSecret interface {
	constructs.Construct
	awssecretsmanager.ISecret
	// The customer-managed encryption key that is used to encrypt this secret, if any.
	//
	// When not specified, the default
	// KMS key for the account and region is being used.
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
	// The SNS topic that receives expiration notifications.
	//
	// Only set when expiration notifications are enabled.
	ExpirationNotificationTopic() awssns.ITopic
	// The tree node.
	Node() constructs.Node
	// The ARN of the secret in AWS Secrets Manager.
	//
	// Will return the full ARN if available, otherwise a partial arn.
	// For secrets imported by the deprecated `fromSecretName`, it will return the `secretName`.
	SecretArn() *string
	// The full ARN of the secret in AWS Secrets Manager, which is the ARN including the Secrets Manager-supplied 6-character suffix.
	//
	// This is equal to `secretArn` in most cases, but is undefined when a full ARN is not available (e.g., secrets imported by name).
	SecretFullArn() *string
	// The name of the secret.
	//
	// For "owned" secrets, this will be the full resource name (secret name + suffix), unless the
	// '@aws-cdk/aws-secretsmanager:parseOwnedSecretName' feature flag is set.
	SecretName() *string
	// A reference to a Secret resource.
	SecretRef() *interfacesawssecretsmanager.SecretReference
	// Retrieve the value of the stored secret as a `SecretValue`.
	SecretValue() awscdk.SecretValue
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	Sync() SopsSync
	// Adds a rotation schedule to the secret.
	AddRotationSchedule(id *string, options *awssecretsmanager.RotationScheduleOptions) awssecretsmanager.RotationSchedule
	// Adds a statement to the IAM resource policy associated with this secret.
	//
	// If this secret was created in this stack, a resource policy will be
	// automatically created upon the first call to `addToResourcePolicy`. If
	// the secret is imported, then this is a no-op.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
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
	// Attach a target to this secret.
	Attach(target awssecretsmanager.ISecretAttachmentTarget) awssecretsmanager.ISecret
	// Returns a key which can be used within an AWS CloudFormation dynamic reference to dynamically load this secret from AWS Secrets Manager.
	CfnDynamicReferenceKey(options *awscdk.SecretsManagerSecretOptions) *string
	CurrentVersionId() *string
	// Denies the `DeleteSecret` action to all principals within the current account.
	DenyAccountRootDelete()
	// Grants reading the secret value to some role.
	GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant
	// Grants writing and updating the secret value to some role.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Interpret the secret as a JSON object and return a field's value from it as a `SecretValue`.
	SecretValueFromJson(key *string) awscdk.SecretValue
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

// The jsii proxy struct for SopsSecret
type jsiiProxy_SopsSecret struct {
	internal.Type__constructsConstruct
	internal.Type__awssecretsmanagerISecret
}

func (j *jsiiProxy_SopsSecret) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsSecret) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsSecret) ExpirationNotificationTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"expirationNotificationTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsSecret) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsSecret) SecretFullArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretFullArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsSecret) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsSecret) SecretRef() *interfacesawssecretsmanager.SecretReference {
	var returns *interfacesawssecretsmanager.SecretReference
	_jsii_.Get(
		j,
		"secretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsSecret) SecretValue() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"secretValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsSecret) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SopsSecret) Sync() SopsSync {
	var returns SopsSync
	_jsii_.Get(
		j,
		"sync",
		&returns,
	)
	return returns
}


func NewSopsSecret(scope constructs.Construct, id *string, props *SopsSecretProps) SopsSecret {
	_init_.Initialize()

	if err := validateNewSopsSecretParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SopsSecret{}

	_jsii_.Create(
		"cdk-sops-secrets.SopsSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSopsSecret_Override(s SopsSecret, scope constructs.Construct, id *string, props *SopsSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-sops-secrets.SopsSecret",
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
func SopsSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSopsSecret_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-sops-secrets.SopsSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SopsSecret) AddRotationSchedule(id *string, options *awssecretsmanager.RotationScheduleOptions) awssecretsmanager.RotationSchedule {
	if err := s.validateAddRotationScheduleParameters(id, options); err != nil {
		panic(err)
	}
	var returns awssecretsmanager.RotationSchedule

	_jsii_.Invoke(
		s,
		"addRotationSchedule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SopsSecret) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := s.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		s,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SopsSecret) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SopsSecret) Attach(target awssecretsmanager.ISecretAttachmentTarget) awssecretsmanager.ISecret {
	if err := s.validateAttachParameters(target); err != nil {
		panic(err)
	}
	var returns awssecretsmanager.ISecret

	_jsii_.Invoke(
		s,
		"attach",
		[]interface{}{target},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SopsSecret) CfnDynamicReferenceKey(options *awscdk.SecretsManagerSecretOptions) *string {
	if err := s.validateCfnDynamicReferenceKeyParameters(options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"cfnDynamicReferenceKey",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SopsSecret) CurrentVersionId() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"currentVersionId",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SopsSecret) DenyAccountRootDelete() {
	_jsii_.InvokeVoid(
		s,
		"denyAccountRootDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SopsSecret) GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant {
	if err := s.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantRead",
		[]interface{}{grantee, versionStages},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SopsSecret) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
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

func (s *jsiiProxy_SopsSecret) SecretValueFromJson(key *string) awscdk.SecretValue {
	if err := s.validateSecretValueFromJsonParameters(key); err != nil {
		panic(err)
	}
	var returns awscdk.SecretValue

	_jsii_.Invoke(
		s,
		"secretValueFromJson",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SopsSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SopsSecret) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

