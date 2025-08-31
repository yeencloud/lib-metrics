package disabled

import (
	"context"

	MetricsDomain "github.com/yeencloud/lib-metrics/domain"
)

type DisabledMetrics struct {
}

func (i *DisabledMetrics) WritePoint(ctx context.Context, metricHeader MetricsDomain.Point, metricValues MetricsDomain.Values) {

}

func (i *DisabledMetrics) Connect() error {
	return nil
}
