//go:build no_runtime_type_checking

package cdksopssecrets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SopsSecret) validateAddRotationScheduleParameters(id *string, options *awssecretsmanager.RotationScheduleOptions) error {
	return nil
}

func (s *jsiiProxy_SopsSecret) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SopsSecret) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SopsSecret) validateAttachParameters(target awssecretsmanager.ISecretAttachmentTarget) error {
	return nil
}

func (s *jsiiProxy_SopsSecret) validateCfnDynamicReferenceKeyParameters(options *awscdk.SecretsManagerSecretOptions) error {
	return nil
}

func (s *jsiiProxy_SopsSecret) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SopsSecret) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SopsSecret) validateSecretValueFromJsonParameters(key *string) error {
	return nil
}

func validateSopsSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSopsSecretParameters(scope constructs.Construct, id *string, props *SopsSecretProps) error {
	return nil
}

