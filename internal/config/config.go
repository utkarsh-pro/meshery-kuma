package config

import (
	"path"

	"github.com/layer5io/meshery-adapter-library/config"
	configprovider "github.com/layer5io/meshery-adapter-library/config/provider"
	"github.com/layer5io/meshkit/utils"
)

const (
	KumaOperation = "kuma"
	Development   = "development"
	Production    = "production"
)

var (
	configRootPath = path.Join(utils.GetHome(), ".meshery")
)

// New creates a new config instance
func New(provider string) (config.Handler, error) {
	opts := configprovider.Options{
		ServerConfig:   ServerDefaults,
		MeshSpec:       MeshSpecDefaults,
		ProviderConfig: ProviderConfigDefaults,
		Operations:     OperationsDefaults,
	}

	// Config provider
	switch provider {
	case configprovider.ViperKey:
		return configprovider.NewViper(opts)
	case configprovider.InMemKey:
		return configprovider.NewInMem(opts)
	}

	return nil, config.ErrEmptyConfig
}

func NewKubeconfigBuilder(provider string) (config.Handler, error) {

	opts := configprovider.Options{}

	// Config environment
	opts.ProviderConfig = KubeConfigDefaults

	// Config provider
	switch provider {
	case configprovider.ViperKey:
		return configprovider.NewViper(opts)
	case configprovider.InMemKey:
		return configprovider.NewInMem(opts)
	}
	return nil, config.ErrEmptyConfig
}

// RootPath returns the configRootPath
func RootPath() string {
	return configRootPath
}
