//go:build no_runtime_type_checking

package cdksopssecrets

// Building without runtime type checking enabled, so all the below just return nil

func validateSopsSync_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSopsSyncParameters(scope constructs.Construct, id *string, props *SopsSyncProps) error {
	return nil
}

