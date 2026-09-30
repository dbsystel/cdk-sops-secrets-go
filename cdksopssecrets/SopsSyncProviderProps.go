package cdksopssecrets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Configuration options for a custom SopsSyncProvider.
type SopsSyncProviderProps struct {
	// The log group the function sends logs to.
	//
	// By default, Lambda functions send logs to an automatically created default log group named /aws/lambda/{function-name}.
	// However you cannot change the properties of this auto-created log group using the AWS CDK, e.g. you cannot set a different log retention.
	//
	// Use the `logGroup` property to create a fully customizable LogGroup ahead of time, and instruct the Lambda function to send logs to it.
	//
	// Providing a user-controlled log group was rolled out to commercial regions on 2023-11-16.
	// If you are deploying to another type of region, please check regional availability first.
	// Default: `/aws/lambda/${this.functionName}` - default log group created by Lambda
	//
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	//
	// This is a legacy API and we strongly recommend you move away from it if you can.
	// Instead create a fully customizable log group with `logs.LogGroup` and use the `logGroup` property
	// to instruct the Lambda function to send logs to it.
	// Migrating from `logRetention` to `logGroup` will cause the name of the log group to change.
	// Users and code and referencing the name verbatim will have to adjust.
	//
	// In AWS CDK code, you can access the log group name directly from the LogGroup construct:
	// ```ts
	// import * as logs from 'aws-cdk-lib/aws-logs';
	//
	// declare const myLogGroup: logs.LogGroup;
	// myLogGroup.logGroupName;
	// ```.
	// Default: logs.RetentionDays.INFINITE
	//
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// The role that should be used for the custom resource provider.
	//
	// If you don't specify any, a new role will be created with all required permissions.
	// Default: - a new role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Only if `vpc` is supplied: The list of security groups to associate with the Lambda's network interfaces.
	// Default: - A dedicated security group will be created for the lambda function.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// A unique identifier to identify this provider.
	//
	// Overwrite the default, if you need a dedicated provider.
	// Default: SopsSyncProvider.
	//
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
	// VPC network to place Lambda network interfaces.
	// Default: - Lambda function is not placed within a VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place the network interfaces within the VPC.
	// Default: - Subnets will be chosen automatically.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

