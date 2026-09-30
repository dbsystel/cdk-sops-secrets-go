// CDK Constructs that syncs your sops secrets into AWS SecretsManager secrets.
package cdksopssecrets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk-sops-secrets.ExpirationOptions",
		reflect.TypeOf((*ExpirationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-sops-secrets.MultiStringParameter",
		reflect.TypeOf((*MultiStringParameter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "keyPrefix", GoGetter: "KeyPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "keySeparator", GoGetter: "KeySeparator"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "sync", GoGetter: "Sync"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_MultiStringParameter{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-sops-secrets.MultiStringParameterProps",
		reflect.TypeOf((*MultiStringParameterProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-sops-secrets.RawOutput",
		reflect.TypeOf((*RawOutput)(nil)).Elem(),
		map[string]interface{}{
			"STRING": RawOutput_STRING,
			"BINARY": RawOutput_BINARY,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-sops-secrets.ResourceType",
		reflect.TypeOf((*ResourceType)(nil)).Elem(),
		map[string]interface{}{
			"SECRET": ResourceType_SECRET,
			"SECRET_RAW": ResourceType_SECRET_RAW,
			"SECRET_BINARY": ResourceType_SECRET_BINARY,
			"PARAMETER": ResourceType_PARAMETER,
			"PARAMETER_MULTI": ResourceType_PARAMETER_MULTI,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-sops-secrets.SopsCommonParameterProps",
		reflect.TypeOf((*SopsCommonParameterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-sops-secrets.SopsSecret",
		reflect.TypeOf((*SopsSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRotationSchedule", GoMethod: "AddRotationSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "attach", GoMethod: "Attach"},
			_jsii_.MemberMethod{JsiiMethod: "cfnDynamicReferenceKey", GoMethod: "CfnDynamicReferenceKey"},
			_jsii_.MemberMethod{JsiiMethod: "currentVersionId", GoMethod: "CurrentVersionId"},
			_jsii_.MemberMethod{JsiiMethod: "denyAccountRootDelete", GoMethod: "DenyAccountRootDelete"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "expirationNotificationTopic", GoGetter: "ExpirationNotificationTopic"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "secretArn", GoGetter: "SecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "secretFullArn", GoGetter: "SecretFullArn"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretRef", GoGetter: "SecretRef"},
			_jsii_.MemberProperty{JsiiProperty: "secretValue", GoGetter: "SecretValue"},
			_jsii_.MemberMethod{JsiiMethod: "secretValueFromJson", GoMethod: "SecretValueFromJson"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "sync", GoGetter: "Sync"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_SopsSecret{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awssecretsmanagerISecret)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-sops-secrets.SopsSecretProps",
		reflect.TypeOf((*SopsSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-sops-secrets.SopsStringParameter",
		reflect.TypeOf((*SopsStringParameter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "parameterArn", GoGetter: "ParameterArn"},
			_jsii_.MemberProperty{JsiiProperty: "parameterName", GoGetter: "ParameterName"},
			_jsii_.MemberProperty{JsiiProperty: "parameterRef", GoGetter: "ParameterRef"},
			_jsii_.MemberProperty{JsiiProperty: "parameterType", GoGetter: "ParameterType"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stringValue", GoGetter: "StringValue"},
			_jsii_.MemberProperty{JsiiProperty: "sync", GoGetter: "Sync"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_SopsStringParameter{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awsssmIStringParameter)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-sops-secrets.SopsStringParameterProps",
		reflect.TypeOf((*SopsStringParameterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-sops-secrets.SopsSync",
		reflect.TypeOf((*SopsSync)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "versionId", GoGetter: "VersionId"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_SopsSync{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-sops-secrets.SopsSyncOptions",
		reflect.TypeOf((*SopsSyncOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-sops-secrets.SopsSyncProps",
		reflect.TypeOf((*SopsSyncProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-sops-secrets.SopsSyncProvider",
		reflect.TypeOf((*SopsSyncProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAgeKey", GoMethod: "AddAgeKey"},
			_jsii_.MemberMethod{JsiiMethod: "addAgeKeyFromSsmParameter", GoMethod: "AddAgeKeyFromSsmParameter"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addEnvironment", GoMethod: "AddEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "addEventSource", GoMethod: "AddEventSource"},
			_jsii_.MemberMethod{JsiiMethod: "addEventSourceMapping", GoMethod: "AddEventSourceMapping"},
			_jsii_.MemberMethod{JsiiMethod: "addFunctionUrl", GoMethod: "AddFunctionUrl"},
			_jsii_.MemberMethod{JsiiMethod: "addLayers", GoMethod: "AddLayers"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addPermission", GoMethod: "AddPermission"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "architecture", GoGetter: "Architecture"},
			_jsii_.MemberProperty{JsiiProperty: "canCreatePermissions", GoGetter: "CanCreatePermissions"},
			_jsii_.MemberMethod{JsiiMethod: "configureAsyncInvoke", GoMethod: "ConfigureAsyncInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "considerWarningOnInvokeFunctionPermissions", GoMethod: "ConsiderWarningOnInvokeFunctionPermissions"},
			_jsii_.MemberProperty{JsiiProperty: "constructName", GoGetter: "ConstructName"},
			_jsii_.MemberProperty{JsiiProperty: "currentVersion", GoGetter: "CurrentVersion"},
			_jsii_.MemberMethod{JsiiMethod: "dependOn", GoMethod: "DependOn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberProperty{JsiiProperty: "functionRef", GoGetter: "FunctionRef"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeCompositePrincipal", GoMethod: "GrantInvokeCompositePrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeLatestVersion", GoMethod: "GrantInvokeLatestVersion"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeUrl", GoMethod: "GrantInvokeUrl"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeVersion", GoMethod: "GrantInvokeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "isBoundToVpc", GoGetter: "IsBoundToVpc"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersion", GoGetter: "LatestVersion"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricDuration", GoMethod: "MetricDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrors", GoMethod: "MetricErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottles", GoMethod: "MetricThrottles"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissionsNode", GoGetter: "PermissionsNode"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceArnsForGrantInvoke", GoGetter: "ResourceArnsForGrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "runtime", GoGetter: "Runtime"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tenancyConfig", GoGetter: "TenancyConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "warnInvokeFunctionPermissions", GoMethod: "WarnInvokeFunctionPermissions"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_SopsSyncProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaSingletonFunction)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-sops-secrets.SopsSyncProviderProps",
		reflect.TypeOf((*SopsSyncProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-sops-secrets.UploadType",
		reflect.TypeOf((*UploadType)(nil)).Elem(),
		map[string]interface{}{
			"INLINE": UploadType_INLINE,
			"ASSET": UploadType_ASSET,
		},
	)
}
