package sfxclient

import "github.com/signalfx/golib/datapoint"

// MultiCollector acts like a datapoint collector over multiple collectors
type MultiCollector []Collector

// NewMultiCollector returns a collector that is the aggregate of every given collector
func NewMultiCollector(collectors ...Collector) Collector {
	if len(collectors) == 1 {
		return collectors[0]
	}
	return MultiCollector(collectors)
}

// Datapoints returns the datapoints from every collector
func (m MultiCollector) Datapoints() []*datapoint.Datapoint {
	var ret []*datapoint.Datapoint
	for _, col := range m {
		ret = append(ret, col.Datapoints()...)
	}
	return ret
}

var _ Collector = MultiCollector(nil)

// WithDimensions adds dimensions on top of the datapoints of a collector
type WithDimensions struct {
	Dimensions map[string]string
	Collector  Collector
}

var _ Collector = &WithDimensions{}

// Datapoints calls datapoints and adds on Dimensions
func (w *WithDimensions) Datapoints() []*datapoint.Datapoint {
	dps := w.Collector.Datapoints()
	for _, dp := range dps {
		dp.Dimensions = AddMaps(dp.Dimensions, w.Dimensions)
	}
	return dps
}