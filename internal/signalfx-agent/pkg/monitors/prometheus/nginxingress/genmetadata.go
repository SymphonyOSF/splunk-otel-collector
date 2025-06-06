// Code generated by monitor-code-gen. DO NOT EDIT.

package nginxingress

import (
	"github.com/signalfx/golib/v3/datapoint"

	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "prometheus/nginx-ingress"

var groupSet = map[string]bool{}

const (
	nginxIngressControllerIngressUpstreamLatencySeconds   = "nginx_ingress_controller_ingress_upstream_latency_seconds"
	nginxIngressControllerNginxProcessConnections         = "nginx_ingress_controller_nginx_process_connections"
	nginxIngressControllerNginxProcessConnectionsTotal    = "nginx_ingress_controller_nginx_process_connections_total"
	nginxIngressControllerNginxProcessCPUSecondsTotal     = "nginx_ingress_controller_nginx_process_cpu_seconds_total"
	nginxIngressControllerNginxProcessResidentMemoryBytes = "nginx_ingress_controller_nginx_process_resident_memory_bytes"
	nginxIngressControllerNginxProcessVirtualMemoryBytes  = "nginx_ingress_controller_nginx_process_virtual_memory_bytes"
	nginxIngressControllerRequestDurationSeconds          = "nginx_ingress_controller_request_duration_seconds"
	nginxIngressControllerRequests                        = "nginx_ingress_controller_requests"
)

var metricSet = map[string]monitors.MetricInfo{
	nginxIngressControllerIngressUpstreamLatencySeconds:   {Type: datapoint.Counter},
	nginxIngressControllerNginxProcessConnections:         {Type: datapoint.Gauge},
	nginxIngressControllerNginxProcessConnectionsTotal:    {Type: datapoint.Counter},
	nginxIngressControllerNginxProcessCPUSecondsTotal:     {Type: datapoint.Counter},
	nginxIngressControllerNginxProcessResidentMemoryBytes: {Type: datapoint.Gauge},
	nginxIngressControllerNginxProcessVirtualMemoryBytes:  {Type: datapoint.Gauge},
	nginxIngressControllerRequestDurationSeconds:          {Type: datapoint.Counter},
	nginxIngressControllerRequests:                        {Type: datapoint.Counter},
}

var defaultMetrics = map[string]bool{
	nginxIngressControllerIngressUpstreamLatencySeconds:   true,
	nginxIngressControllerNginxProcessCPUSecondsTotal:     true,
	nginxIngressControllerNginxProcessResidentMemoryBytes: true,
	nginxIngressControllerRequestDurationSeconds:          true,
	nginxIngressControllerRequests:                        true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "prometheus/nginx-ingress",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
