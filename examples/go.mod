module example

go 1.15

replace github.com/marlonfan/watchat-exporter => ../

require (
	github.com/marlonfan-exporter/watchat v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.12.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)
