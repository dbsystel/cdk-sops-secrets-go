//go:build no_runtime_type_checking

package cdksopssecrets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SopsSyncProvider) validateAddAgeKeyParameters(key awscdk.SecretValue) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateAddAgeKeyFromSsmParameterParameters(param interface{}, encryptionKey awskms.IKey) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateAddEnvironmentParameters(key *string, value *string, options *awslambda.EnvironmentOptions) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateAddEventSourceParameters(source awslambda.IEventSource) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateAddEventSourceMappingParameters(id *string, options *awslambda.EventSourceMappingOptions) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateAddFunctionUrlParameters(options *awslambda.FunctionUrlOptions) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateAddMetadataParameters(type_ *string, data interface{}, options *constructs.MetadataOptions) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateAddPermissionParameters(id *string, permission *awslambda.Permission) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateConfigureAsyncInvokeParameters(options *awslambda.EventInvokeConfigOptions) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope constructs.Construct, action *string) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateDependOnParameters(down constructs.IConstruct) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateGrantInvokeParameters(identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateGrantInvokeCompositePrincipalParameters(compositePrincipal awsiam.CompositePrincipal) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateGrantInvokeLatestVersionParameters(identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateGrantInvokeUrlParameters(identity awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateGrantInvokeVersionParameters(identity awsiam.IGrantable, version awslambda.IVersion) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SopsSyncProvider) validateWarnInvokeFunctionPermissionsParameters(scope constructs.Construct) error {
	return nil
}

func validateSopsSyncProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSopsSyncProvider_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSopsSyncProvider_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSopsSyncProviderParameters(scope constructs.Construct, props *SopsSyncProviderProps) error {
	return nil
}

