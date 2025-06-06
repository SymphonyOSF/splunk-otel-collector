// Code generated by monitor-code-gen. DO NOT EDIT.

package ntp

import (
	"github.com/signalfx/golib/v3/datapoint"

	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "ntp"

var groupSet = map[string]bool{}

const (
	ntpOffsetSeconds = "ntp.offset_seconds"
)

var metricSet = map[string]monitors.MetricInfo{
	ntpOffsetSeconds: {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	ntpOffsetSeconds: true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "ntp",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
