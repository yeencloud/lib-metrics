package metrics

import (
	"github.com/yeencloud/lib-metrics/database/influx"
	"github.com/yeencloud/lib-metrics/domain"
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

func (m *Metrics) Connect() error {
	return m.provider.Connect()
}

func Connect() error {
	if localMetrics == nil {
		return &errors.MetricsNotInitializedError{}
	}

	return localMetrics.Connect()
}

func (m *Metrics) LogPoint(point MetricsDomain.Point, values MetricsDomain.Values) {
	if point.Tags == nil {
		point.Tags = make(map[string]string)
	}

	point.Tags["service"] = m.serviceName
	point.Tags["hostname"] = m.hostname

	m.provider.LogPoint(point, values)
}

func LogPoint(point MetricsDomain.Point, values MetricsDomain.Values) error {
	if localMetrics == nil {
		return &errors.MetricsNotInitializedError{}
	}

	localMetrics.LogPoint(point, values)
	return nil
}
