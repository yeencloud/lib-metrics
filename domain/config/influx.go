package MetricsConfig

import (
	"fmt"

	"github.com/yeencloud/lib-shared/config"
)

type InfluxConfig struct {
	Host string `config:"INFLUXDB_HOST" default:"localhost"`
	Port int    `config:"INFLUXDB_PORT" default:"8086"`

	Token        config.Secret `config:"INFLUXDB_TOKEN"`
	Organization string        `config:"INFLUXDB_ORG"`
	Bucket       string        `config:"INFLUXDB_BUCKET"`
}

func (c *InfluxConfig) GetAddress() string {
	return fmt.Sprintf("http://%s:%d", c.Host, c.Port)
}
