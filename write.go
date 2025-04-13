package metrics

import (
	"context"
	"fmt"

	"github.com/fatih/structs"
	MetricsDomain "github.com/yeencloud/lib-metrics/domain"
	"github.com/yeencloud/lib-metrics/errors"
	"github.com/yeencloud/lib-shared/metrics"
)

func (m *Metrics) NewPoint() MetricsDomain.Point {
	return MetricsDomain.Point{
		Tags: map[string]string{
			"service":  m.serviceName,
			"hostname": m.hostname,
		},
	}
}

func NewPoint() MetricsDomain.Point {
	if localMetrics == nil {
		return MetricsDomain.Point{}
	}

	return localMetrics.NewPoint()
}

func (m *Metrics) SetTag(ctx context.Context, key string, value any) context.Context {
	if localMetrics == nil {
		return ctx
	}

	point := m.getMetricPointFromCtx(ctx)
	point.Tags[key] = fmt.Sprintf("%v", value)
	ctx = context.WithValue(ctx, metrics.MetricsPointKey, point)

	return ctx
}

func SetTag(ctx context.Context, key string, value any) context.Context {
	if localMetrics == nil {
		return ctx
	}

	return localMetrics.SetTag(ctx, key, value)
}

func (m *Metrics) CreateMetricPoint(ctx context.Context) context.Context {
	if localMetrics == nil {
		return ctx
	}

	point := m.NewPoint()

	ctx = context.WithValue(ctx, metrics.MetricsPointKey, point)

	return ctx
}

func (m *Metrics) getMetricPointFromCtx(ctx context.Context) MetricsDomain.Point {
	point, ok := ctx.Value(metrics.MetricsPointKey).(MetricsDomain.Point)
	if !ok {
		point = m.NewPoint()
	}

	return point
}

func (m *Metrics) WritePoint(ctx context.Context, metricType string, point any) error {
	if localMetrics == nil {
		return &errors.MetricsNotInitializedError{}
	}

	if ctx == nil {
		return nil
	}

	header := localMetrics.getMetricPointFromCtx(ctx)
	values := MetricsDomain.Values{}

	header.Name = metricType

	if structs.IsStruct(point) {
		fields := structs.Fields(point)

		for _, field := range fields {
			metricTag := field.Tag("metric")
			if field.IsExported() && metricTag != "" {
				values[metricTag] = field.Value()
			}
		}
	}

	localMetrics.provider.WritePoint(ctx, header, values)
	return nil
}

func WritePoint(ctx context.Context, metricType string, point any) error {
	if localMetrics == nil {
		return &errors.MetricsNotInitializedError{}
	}

	return localMetrics.WritePoint(ctx, metricType, point)
}
