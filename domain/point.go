package MetricsDomain

import (
	log "github.com/sirupsen/logrus"

	"github.com/yeencloud/lib-shared/namespace"
)

type Tags map[string]string

type Point struct {
	Name string
	Tags Tags
}

func (p *Point) SetTag(value namespace.NamespaceValue) {
	if p.Tags == nil {
		p.Tags = make(Tags)
	}

	if !value.MetricTag() {
		log.WithFields(log.Fields{
			"key": value.Namespace.MetricKey(),
			"val": value.String(),
		}).Warn("Point.SetTag: called with non-metric tag")
	}

	p.Tags[value.Namespace.MetricKey()] = value.String()
}

type Values map[string]any
