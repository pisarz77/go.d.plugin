package snmp

import (
	"fmt"

	"github.com/netdata/go.d.plugin/agent/module"
)

var defaultSNMPchart = module.Chart{
	ID:       "snmp_%d",
	Title:    "%s",
	Units:    "kilobits/s",
	Type:     "area",
	Priority: 1,
	Fam:      "ports",
}

var defaultDims = module.Dims{
	{
		Name: "in",
		ID:   "1.3.6.1.2.1.2.2.1.10.2",
		Algo: module.Incremental,
		Mul:  8,
		Div:  1024,
	},
	{
		Name: "out",
		ID:   "1.3.6.1.2.1.2.2.1.16.2",
		Algo: module.Incremental,
		Mul:  -8,
		Div:  1024,
	},
}

func newChart(chartIn []ChartsConfig) *module.Charts {
	charts := &module.Charts{}
	for i, s := range chartIn {
		c := defaultSNMPchart.Copy()
		c.ID = fmt.Sprintf(c.ID, i)
		c.Title = fmt.Sprintf(c.Title, s.Title)
		if s.Family != nil {
			c.Fam = *s.Family
		}

		if s.Units != nil {
			c.Units = *s.Units
		}

		if s.Type != nil {
			c.Type = module.ChartType(*s.Type)
		}

		c.Priority = s.Priority
		for _, d := range s.Dimensions {
			dim := &module.Dim{
				Name: d.Name,
				ID:   d.OID,
			}
			if d.Algorithm != nil {
				dim.Algo = module.DimAlgo(*d.Algorithm)
			}
			if d.Multiplier != nil {
				dim.Mul = *d.Multiplier
			}
			if d.Divisor != nil {
				dim.Div = *d.Divisor
			}
			_ = c.AddDim(dim)
		}

		//Add default ones if no dimensions defined
		if len(c.Dims) == 0 {
			_ = c.AddDim(defaultDims[0])
			_ = c.AddDim(defaultDims[1])
		}
		_ = charts.Add(c)
	}
	return charts
}
