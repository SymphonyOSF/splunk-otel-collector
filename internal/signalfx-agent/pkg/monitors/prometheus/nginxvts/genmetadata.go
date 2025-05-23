// Code generated by monitor-code-gen. DO NOT EDIT.

package nginxvts

import (
	"github.com/signalfx/golib/v3/datapoint"

	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "prometheus/nginx-vts"

var groupSet = map[string]bool{}

const (
	nginxVtsInfo                            = "nginx_vts_info"
	nginxVtsMainConnections                 = "nginx_vts_main_connections"
	nginxVtsMainShmUsageBytes               = "nginx_vts_main_shm_usage_bytes"
	nginxVtsServerBytesTotal                = "nginx_vts_server_bytes_total"
	nginxVtsServerCacheTotal                = "nginx_vts_server_cache_total"
	nginxVtsServerRequestDurationSeconds    = "nginx_vts_server_request_duration_seconds"
	nginxVtsServerRequestSeconds            = "nginx_vts_server_request_seconds"
	nginxVtsServerRequestSecondsTotal       = "nginx_vts_server_request_seconds_total"
	nginxVtsServerRequestsTotal             = "nginx_vts_server_requests_total"
	nginxVtsStartTimeSeconds                = "nginx_vts_start_time_seconds"
	nginxVtsUpstreamBytesTotal              = "nginx_vts_upstream_bytes_total"
	nginxVtsUpstreamRequestDurationSeconds  = "nginx_vts_upstream_request_duration_seconds"
	nginxVtsUpstreamRequestSeconds          = "nginx_vts_upstream_request_seconds"
	nginxVtsUpstreamRequestSecondsTotal     = "nginx_vts_upstream_request_seconds_total"
	nginxVtsUpstreamRequestsTotal           = "nginx_vts_upstream_requests_total"
	nginxVtsUpstreamResponseDurationSeconds = "nginx_vts_upstream_response_duration_seconds"
	nginxVtsUpstreamResponseSeconds         = "nginx_vts_upstream_response_seconds"
	nginxVtsUpstreamResponseSecondsTotal    = "nginx_vts_upstream_response_seconds_total"
)

var metricSet = map[string]monitors.MetricInfo{
	nginxVtsInfo:                            {Type: datapoint.Gauge},
	nginxVtsMainConnections:                 {Type: datapoint.Gauge},
	nginxVtsMainShmUsageBytes:               {Type: datapoint.Gauge},
	nginxVtsServerBytesTotal:                {Type: datapoint.Counter},
	nginxVtsServerCacheTotal:                {Type: datapoint.Counter},
	nginxVtsServerRequestDurationSeconds:    {Type: datapoint.Counter},
	nginxVtsServerRequestSeconds:            {Type: datapoint.Gauge},
	nginxVtsServerRequestSecondsTotal:       {Type: datapoint.Counter},
	nginxVtsServerRequestsTotal:             {Type: datapoint.Counter},
	nginxVtsStartTimeSeconds:                {Type: datapoint.Gauge},
	nginxVtsUpstreamBytesTotal:              {Type: datapoint.Counter},
	nginxVtsUpstreamRequestDurationSeconds:  {Type: datapoint.Counter},
	nginxVtsUpstreamRequestSeconds:          {Type: datapoint.Gauge},
	nginxVtsUpstreamRequestSecondsTotal:     {Type: datapoint.Counter},
	nginxVtsUpstreamRequestsTotal:           {Type: datapoint.Counter},
	nginxVtsUpstreamResponseDurationSeconds: {Type: datapoint.Counter},
	nginxVtsUpstreamResponseSeconds:         {Type: datapoint.Gauge},
	nginxVtsUpstreamResponseSecondsTotal:    {Type: datapoint.Counter},
}

var defaultMetrics = map[string]bool{
	nginxVtsMainConnections:        true,
	nginxVtsServerRequestSeconds:   true,
	nginxVtsServerRequestsTotal:    true,
	nginxVtsUpstreamRequestSeconds: true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "prometheus/nginx-vts",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
