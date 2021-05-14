module github.com/layer5io/meshery-kuma

go 1.13

replace (
	github.com/kudobuilder/kuttl => github.com/layer5io/kuttl v0.4.1-0.20200723152044-916f10574334
	gopkg.in/ini.v1 => github.com/go-ini/ini v1.62.0
)

require (
	github.com/layer5io/meshery-adapter-library v0.1.15
	github.com/layer5io/meshkit v0.2.10
	github.com/layer5io/service-mesh-performance v0.3.2
	gopkg.in/yaml.v2 v2.3.0
	helm.sh/helm/v3 v3.3.4 // indirect
)
