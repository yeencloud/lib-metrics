package MetricsInflux

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	log "github.com/sirupsen/logrus"

	MetricsDomain "github.com/yeencloud/lib-metrics/domain"
)

func (i *Influx) newPoint(metricPoint MetricsDomain.Point) *write.Point {
	p := influxdb2.NewPointWithMeasurement(metricPoint.Name)

	for k, v := range metricPoint.Tags {
		p = p.AddTag(k, v)
	}

	return p
}

func (i *Influx) LogPoint(point MetricsDomain.Point, values MetricsDomain.Values) {
	writer := i.writeAPI

	p := i.newPoint(point)

	for k, v := range values {
		p = p.AddField(k, v)
	}
	err := writer.WritePoint(context.Background(), p)

	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"point":  point,
			"values": values,
		}).Error("failed to write point")
	}
}
