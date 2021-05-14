package kuma

import (
	"fmt"

	"github.com/layer5io/meshkit/errors"
)

var (
	ErrInstallKumaCode   = "kuma_test_code"
	ErrMeshConfigCode    = "kuma_test_code"
	ErrFetchManifestCode = "kuma_test_code"
	ErrClientConfigCode  = "kuma_test_code"
	ErrClientSetCode     = "kuma_test_code"
	ErrStreamEventCode   = "kuma_test_code"
	ErrSampleAppCode     = "kuma_test_code"

	// ErrCreatingHelmIndexCode represents the errors which are generated
	// during creation of helm index
	ErrCreatingHelmIndexCode = "kuma_test_code"

	// ErrEntryWithAppVersionNotExistsCode represents the error which is generated
	// when no entry is found with specified name and app version
	ErrEntryWithAppVersionNotExistsCode = "kuma_test_code"

	// ErrHelmRepositoryNotFoundCode represents the error which is generated when
	// no valid helm repository is found
	ErrHelmRepositoryNotFoundCode = "kuma_test_code"

	// ErrDecodeYamlCode represents the error which is generated when yaml
	// decode process fails
	ErrDecodeYamlCode = "kuma_test_code"

	// ErrApplyHelmChartCode represents the error which are generated
	// during the process of applying helm chart
	ErrApplyHelmChartCode = "kuma_test_code"

	// ErrConvertingAppVersionToChartVersionCode represents the errors which are generated
	// during the process of converting app version to chart version
	ErrConvertingAppVersionToChartVersionCode = "kuma_test_code"

	ErrOpInvalid = errors.NewDefault(errors.ErrOpInvalid, "Invalid operation")
)

// ErrInstallKuma is the error for install mesh
func ErrInstallKuma(err error) error {
	return errors.NewDefault(ErrInstallKumaCode, fmt.Sprintf("Error with kuma operation: %s", err.Error()))
}

// ErrMeshConfig is the error for mesh config
func ErrMeshConfig(err error) error {
	return errors.NewDefault(ErrMeshConfigCode, fmt.Sprintf("Error configuration mesh: %s", err.Error()))
}

// ErrFetchManifest is the error occured during the process
// fetching manifest
func ErrFetchManifest(err error, des string) error {
	return errors.NewDefault(ErrFetchManifestCode, fmt.Sprintf("Error fetching mesh manifest: %s", des))
}

// ErrClientConfig is the error for setting client config
func ErrClientConfig(err error) error {
	return errors.NewDefault(ErrClientConfigCode, fmt.Sprintf("Error setting client config: %s", err.Error()))
}

// ErrClientSet is the error for setting clientset
func ErrClientSet(err error) error {
	return errors.NewDefault(ErrClientSetCode, fmt.Sprintf("Error setting clientset: %s", err.Error()))
}

// ErrStreamEvent is the error for streaming event
func ErrStreamEvent(err error) error {
	return errors.NewDefault(ErrStreamEventCode, fmt.Sprintf("Error streaming event: %s", err.Error()))
}

// ErrSampleApp is the error for streaming event
func ErrSampleApp(err error) error {
	return errors.NewDefault(ErrSampleAppCode, fmt.Sprintf("Error with sample app operation: %s", err.Error()))
}

// ErrCreatingHelmIndex is the error for creating helm index
func ErrCreatingHelmIndex(err error) error {
	return errors.NewDefault(ErrCreatingHelmIndexCode, fmt.Sprintf("Error with traefik operation: %s", err.Error()))
}

// ErrEntryWithAppVersionNotExists is the error when an entry with the given app version is not found
func ErrEntryWithAppVersionNotExists(entry, appVersion string) error {
	return errors.NewDefault(
		ErrEntryWithAppVersionNotExistsCode,
		fmt.Sprintf("entry %s with app version %s does not exists", entry, appVersion),
	)
}

// ErrHelmRepositoryNotFound is the error when no valid remote helm repository is found
func ErrHelmRepositoryNotFound(repo string, err error) error {
	return errors.NewDefault(
		ErrHelmRepositoryNotFoundCode,
		fmt.Sprintf("either the repo %s does not exists or is corrupt: %v", repo, err),
	)
}

// ErrDecodeYaml is the error when the yaml unmarshal fails
func ErrDecodeYaml(err error) error {
	return errors.NewDefault(
		ErrDecodeYamlCode,
		fmt.Sprintf("error decoding yaml: %v", err),
	)
}

// ErrApplyHelmChart is the error for applying helm chart
func ErrApplyHelmChart(err error) error {
	return errors.NewDefault(ErrApplyHelmChartCode, fmt.Sprintf("error applying helm chart: %s", err.Error()))
}

// ErrConvertingAppVersionToChartVersion is the error for converting app version to chart version
func ErrConvertingAppVersionToChartVersion(err error) error {
	return errors.NewDefault(
		ErrConvertingAppVersionToChartVersionCode,
		fmt.Sprintf("error converting app version to chart version: %s", err.Error()),
	)
}

// ErrCustomOperation is the error occured during the process of
// applying custom operation
func ErrCustomOperation(err error) error {
	return errors.NewDefault(ErrCustomOperationCode, fmt.Sprintf("Error applying custom operation: %s", err.Error()))
}
