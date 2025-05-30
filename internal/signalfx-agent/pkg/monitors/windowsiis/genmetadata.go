// Code generated by monitor-code-gen. DO NOT EDIT.

package windowsiis

import (
	"github.com/signalfx/golib/v3/datapoint"

	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "windows-iis"

var groupSet = map[string]bool{}

const (
	processHandleCount                  = "process.handle_count"
	processIDProcess                    = "process.id_process"
	processPctProcessorTime             = "process.pct_processor_time"
	processPrivateBytes                 = "process.private_bytes"
	processThreadCount                  = "process.thread_count"
	processVirtualBytes                 = "process.virtual_bytes"
	processWorkingSet                   = "process.working_set"
	webServiceAnonymousUsersSec         = "web_service.anonymous_users_sec"
	webServiceBytesReceivedSec          = "web_service.bytes_received_sec"
	webServiceBytesSentSec              = "web_service.bytes_sent_sec"
	webServiceConnectionAttemptsSec     = "web_service.connection_attempts_sec"
	webServiceCurrentConnections        = "web_service.current_connections"
	webServiceFilesReceivedSec          = "web_service.files_received_sec"
	webServiceFilesSentSec              = "web_service.files_sent_sec"
	webServiceGetRequestsSec            = "web_service.get_requests_sec"
	webServiceIsapiExtensionRequestsSec = "web_service.isapi_extension_requests_sec"
	webServiceNonanonymousUsersSec      = "web_service.nonanonymous_users_sec"
	webServiceNotFoundErrorsSec         = "web_service.not_found_errors_sec"
	webServicePostRequestsSec           = "web_service.post_requests_sec"
	webServiceServiceUptime             = "web_service.service_uptime"
	webServiceTotalMethodRequestsSec    = "web_service.total_method_requests_sec"
)

var metricSet = map[string]monitors.MetricInfo{
	processHandleCount:                  {Type: datapoint.Gauge},
	processIDProcess:                    {Type: datapoint.Gauge},
	processPctProcessorTime:             {Type: datapoint.Gauge},
	processPrivateBytes:                 {Type: datapoint.Gauge},
	processThreadCount:                  {Type: datapoint.Gauge},
	processVirtualBytes:                 {Type: datapoint.Gauge},
	processWorkingSet:                   {Type: datapoint.Gauge},
	webServiceAnonymousUsersSec:         {Type: datapoint.Gauge},
	webServiceBytesReceivedSec:          {Type: datapoint.Gauge},
	webServiceBytesSentSec:              {Type: datapoint.Gauge},
	webServiceConnectionAttemptsSec:     {Type: datapoint.Gauge},
	webServiceCurrentConnections:        {Type: datapoint.Gauge},
	webServiceFilesReceivedSec:          {Type: datapoint.Gauge},
	webServiceFilesSentSec:              {Type: datapoint.Gauge},
	webServiceGetRequestsSec:            {Type: datapoint.Gauge},
	webServiceIsapiExtensionRequestsSec: {Type: datapoint.Gauge},
	webServiceNonanonymousUsersSec:      {Type: datapoint.Gauge},
	webServiceNotFoundErrorsSec:         {Type: datapoint.Gauge},
	webServicePostRequestsSec:           {Type: datapoint.Gauge},
	webServiceServiceUptime:             {Type: datapoint.Gauge},
	webServiceTotalMethodRequestsSec:    {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	webServiceAnonymousUsersSec:      true,
	webServiceBytesReceivedSec:       true,
	webServiceBytesSentSec:           true,
	webServiceConnectionAttemptsSec:  true,
	webServiceCurrentConnections:     true,
	webServiceFilesReceivedSec:       true,
	webServiceFilesSentSec:           true,
	webServiceGetRequestsSec:         true,
	webServiceNonanonymousUsersSec:   true,
	webServiceNotFoundErrorsSec:      true,
	webServicePostRequestsSec:        true,
	webServiceTotalMethodRequestsSec: true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "windows-iis",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
