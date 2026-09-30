//go:build no_runtime_type_checking

package cdksopssecrets

// Building without runtime type checking enabled, so all the below just return nil

func validateMultiStringParameter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewMultiStringParameterParameters(scope constructs.Construct, id *string, props *MultiStringParameterProps) error {
	return nil
}

