package MetricsInflux

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"

	"github.com/yeencloud/lib-metrics/domain/config"
	"github.com/yeencloud/lib-metrics/ports"
	"github.com/yeencloud/lib-shared/config"
)

type Influx struct {
	cfg *MetricsConfig.InfluxConfig

	client influxdb2.Client

	writeAPI api.WriteAPIBlocking
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
