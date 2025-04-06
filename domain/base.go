package MetricsDomain

import "time"

type BaseMetric struct {
	Time time.Time `metric:"timestamp"`

	AdditionalTags   map[string]string
	AdditionalFields map[string]any
}
