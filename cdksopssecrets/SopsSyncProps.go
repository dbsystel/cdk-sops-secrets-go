package cdksopssecrets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// The configuration options extended by the target Secret / Parameter.
type SopsSyncProps struct {
	// The encryption key used by the CDK default Asset S3 Bucket.
	// Default: - Trying to get the key using the CDK Bootstrap context.
	//
	AssetEncryptionKey awskms.IKey `field:"optional" json:"assetEncryptionKey" yaml:"assetEncryptionKey"`
	// Should this construct automatically create IAM permissions?
	// Default: true.
	//
	AutoGenerateIamPermissions *bool `field:"optional" json:"autoGenerateIamPermissions" yaml:"autoGenerateIamPermissions"`
	// The age key that should be used for encryption.
	SopsAgeKey awscdk.SecretValue `field:"optional" json:"sopsAgeKey" yaml:"sopsAgeKey"`
	// The format of the sops file.
	// Default: - The fileformat will be derived from the file ending.
	//
	SopsFileFormat *string `field:"optional" json:"sopsFileFormat" yaml:"sopsFileFormat"`
	// The filepath to the sops file.
	SopsFilePath *string `field:"optional" json:"sopsFilePath" yaml:"sopsFilePath"`
	// The kmsKey used to encrypt the sops file.
	//
	// Encrypt permissions
	// will be granted to the custom resource provider.
	// Default: - The key will be derived from the sops file.
	//
	SopsKmsKey *[]awskms.IKey `field:"optional" json:"sopsKmsKey" yaml:"sopsKmsKey"`
	// The custom resource provider to use.
	//
	// If you don't specify any, a new
	// provider will be created - or if already exists within this stack - reused.
	// Default: - A new singleton provider will be created.
	//
	SopsProvider SopsSyncProvider `field:"optional" json:"sopsProvider" yaml:"sopsProvider"`
	// If you want to pass the sops file via s3, you can specify the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified.
	SopsS3Bucket *string `field:"optional" json:"sopsS3Bucket" yaml:"sopsS3Bucket"`
	// If you want to pass the sops file via s3, you can specify the key inside the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified.
	SopsS3Key *string `field:"optional" json:"sopsS3Key" yaml:"sopsS3Key"`
	// How should the secret be passed to the CustomResource?
	// Default: INLINE.
	//
	UploadType UploadType `field:"optional" json:"uploadType" yaml:"uploadType"`
	// Will this Sync deploy a Secret or Parameter(s).
	ResourceType ResourceType `field:"required" json:"resourceType" yaml:"resourceType"`
	// The target to populate with the sops file content.
	//
	// - for secret, it's the name or arn of the secret
	// - for parameter, it's the name of the parameter
	// - for parameter multi, it's the prefix of the parameters.
	Target *string `field:"required" json:"target" yaml:"target"`
	// The encryption key used for encrypting the ssm parameter if `parameterName` is set.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// If the structure should be flattened use the provided separator between keys.
	// Default: - undefined.
	//
	FlattenSeparator *string `field:"optional" json:"flattenSeparator" yaml:"flattenSeparator"`
	ParameterNames *[]*string `field:"optional" json:"parameterNames" yaml:"parameterNames"`
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
	// An inert value that triggers an update when related resource properties change.
	// Default: - undefined.
	//
	SyncTrigger *string `field:"optional" json:"syncTrigger" yaml:"syncTrigger"`
}

