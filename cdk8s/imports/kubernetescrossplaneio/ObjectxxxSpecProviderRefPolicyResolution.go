package kubernetescrossplaneio

// Resolution specifies whether resolution of this reference is required.
//
// The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.
type ObjectSpecProviderRefPolicyResolution string

const (
	// Required.
	ObjectSpecProviderRefPolicyResolution_REQUIRED ObjectSpecProviderRefPolicyResolution = "REQUIRED"
	// Optional.
	ObjectSpecProviderRefPolicyResolution_OPTIONAL ObjectSpecProviderRefPolicyResolution = "OPTIONAL"
)
