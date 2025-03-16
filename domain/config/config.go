package MetricsConfig

import (
	"fmt"
	"strings"

	"github.com/yeencloud/lib-shared/config"
)

type Config struct {
	Provider string `config:"METRICS_PROVIDER" default:"influxdb"`
}

func (c Config) IsDisabled() bool {
	return strings.ToLower(c.Provider) == "none"
}

// Influx Config
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
