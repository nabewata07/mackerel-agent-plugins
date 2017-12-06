package mppassenger

import (
	"flag"
	"fmt"
	"strings"
)

type PassengerPlugin struct {
	Prefix string
}

func (p PassengerPlugin) MetricKeyPrefix() string {
	if p.Prefix == "" {
		p.Prefix = "passenger"
	}
	return p.Prefix
}

func (p PassengerPlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(p.Prefix)
	return map[string]mp.Graphs{
		"": mp.Graphs{
			Label: labelPrefix,
			Unit:  mp.UnitFloat, // refer go-mackerel-plugin
			Metrics: []mp.Metrics{
				mp.Metrics{Name: "seconds", Label: "Seconds"}, // use Scale for passenger
			},
		},
	}
}

func (p PassengerPlugin) FetchMetrics() (map[string]float64, error) {
	ut, err := uptime.Get()
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch passenger-status: %s", err)
	}
	return map[string]float64{"seconds": ut}, nil
}

func Do() {
	optPrefix := flag.String("metric-key-prefix", "uptime", "Metric key prefix")
	optTempfile := flag.String("tempfile", "", "Tempfile name")
	flag.Parse()

	p := PassengerPlugin{
		Prefix: *optPrefix,
	}
	helper := mp.NewMackerelPlugin(p)
	helper.Tempfile = *optTempfile
	helper.Run()
}
