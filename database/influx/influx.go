package MetricsInflux

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	MetricsDomain "github.com/yeencloud/lib-metrics/domain"
	"github.com/yeencloud/lib-metrics/domain/config"
	"github.com/yeencloud/lib-metrics/ports"
	"github.com/yeencloud/lib-shared/config"
)

type Influx struct {
	cfg      *MetricsConfig.InfluxConfig
	client   influxdb2.Client
	writeAPI api.WriteAPIBlocking
}

func (i *Influx) WritePoint(ctx context.Context, metricHeader MetricsDomain.Point, metricValues MetricsDomain.Values) {
	writer := i.writeAPI

	p := influxdb2.NewPointWithMeasurement(metricHeader.Name)

	for k, v := range metricHeader.Tags {
		p = p.AddTag(k, v)
	}

	for k, v := range metricValues {
		p = p.AddField(k, v)
	}

	err := writer.WritePoint(ctx, p)
	if err != nil {
		println(err.Error()) //nolint:forbidigo
	}
}

func (i *Influx) Connect() error {
	i.client = influxdb2.NewClient(i.cfg.GetAddress(), i.cfg.Token.Value)
	i.writeAPI = i.client.WriteAPIBlocking(i.cfg.Organization, i.cfg.Bucket)

	_, err := i.client.Ping(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func NewInflux() (ports.MetricsInterface, error) {
	cfg, err := config.FetchConfig[MetricsConfig.InfluxConfig]()
	if err != nil {
		return nil, err
	}

	return &Influx{
		cfg: cfg,
	}, nil
}
