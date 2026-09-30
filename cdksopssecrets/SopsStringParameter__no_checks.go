//go:build no_runtime_type_checking

package cdksopssecrets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SopsStringParameter) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SopsStringParameter) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SopsStringParameter) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateSopsStringParameter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSopsStringParameterParameters(scope constructs.Construct, id *string, props *SopsStringParameterProps) error {
	return nil
}

