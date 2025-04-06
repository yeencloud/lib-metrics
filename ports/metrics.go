package ports

import (
	"context"

	MetricsDomain "github.com/yeencloud/lib-metrics/domain"
)

type MetricsInterface interface {
	WritePoint(ctx context.Context, metricHeader MetricsDomain.Point, metricValues MetricsDomain.Values)
	Connect() error
}
