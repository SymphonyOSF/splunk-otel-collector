// Code generated by monitor-code-gen. DO NOT EDIT.

package dns

import (
	"github.com/signalfx/golib/v3/datapoint"

	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "telegraf/dns"

var groupSet = map[string]bool{}

const (
	dnsQueryTimeMs = "dns.query_time_ms"
	dnsRcodeValue  = "dns.rcode_value"
	dnsResultCode  = "dns.result_code"
)

var metricSet = map[string]monitors.MetricInfo{
	dnsQueryTimeMs: {Type: datapoint.Gauge},
	dnsRcodeValue:  {Type: datapoint.Gauge},
	dnsResultCode:  {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	dnsQueryTimeMs: true,
	dnsRcodeValue:  true,
	dnsResultCode:  true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "telegraf/dns",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
