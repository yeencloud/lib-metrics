package ports

import (
	MetricsDomain "github.com/yeencloud/lib-metrics/domain"
)

type MetricsInterface interface {
	LogPoint(point MetricsDomain.Point, value MetricsDomain.Values)

	Connect() error
}
