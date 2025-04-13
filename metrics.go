package metrics

import (
	"github.com/yeencloud/lib-metrics/database/influx"
	"github.com/yeencloud/lib-metrics/domain/config"
	"github.com/yeencloud/lib-metrics/errors"
	"github.com/yeencloud/lib-metrics/ports"
	"github.com/yeencloud/lib-shared/config"
)

var localMetrics *Metrics

type Metrics struct {
	serviceName string
	hostname    string

	provider ports.MetricsInterface
}

func (m *Metrics) Connect() error {
	return m.provider.Connect()
}

func NewMetrics(serviceName string, hostname string) (*Metrics, error) {
	cfg, err := config.FetchConfig[MetricsConfig.Config]()
	if err != nil {
		return nil, err
	}

	var provider ports.MetricsInterface
	switch cfg.Provider {
	case "influxdb":
		influx, err := MetricsInflux.NewInflux()
		if err != nil {
			return nil, err
		}
		provider = influx
	default:
		return nil, &errors.UnknownProviderError{Provider: cfg.Provider}
	}

	metrics := &Metrics{
		serviceName: serviceName,
		hostname:    hostname,
		provider:    provider,
	}

	localMetrics = metrics

	return metrics, nil
}
