// Code generated by monitor-code-gen. DO NOT EDIT.

package solr

import (
	"github.com/signalfx/golib/v3/datapoint"

	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "collectd/solr"

var groupSet = map[string]bool{}

const (
	counterSolrHTTP2xxResponses                  = "counter.solr.http_2xx_responses"
	counterSolrHTTP4xxResponses                  = "counter.solr.http_4xx_responses"
	counterSolrHTTP5xxResponses                  = "counter.solr.http_5xx_responses"
	counterSolrHTTPRequests                      = "counter.solr.http_requests"
	counterSolrJvmClassesLoaded                  = "counter.solr.jvm_classes_loaded"
	counterSolrNodeCollectionsRequests           = "counter.solr.node_collections_requests"
	counterSolrNodeCoresRequests                 = "counter.solr.node_cores_requests"
	counterSolrNodeMetricRequestCount            = "counter.solr.node_metric_request_count"
	counterSolrNodeMetricsRequests               = "counter.solr.node_metrics_requests"
	counterSolrNodeZookeeperRequests             = "counter.solr.node_zookeeper_requests"
	counterSolrOpenFileDescriptorCount           = "counter.solr.openFileDescriptorCount"
	counterSolrReplicationHandlerRequests        = "counter.solr.replication_handler_requests"
	counterSolrSearchQueryRequests               = "counter.solr.search_query_requests"
	counterSolrUpdateHandlerRequests             = "counter.solr.update_handler_requests"
	counterSolrZookeeperErrors                   = "counter.solr.zookeeper_errors"
	gaugeSolrCoreDeletedDocs                     = "gauge.solr.core_deleted_docs"
	gaugeSolrCoreIndexSize                       = "gauge.solr.core_index_size"
	gaugeSolrCoreMaxDocs                         = "gauge.solr.core_max_docs"
	gaugeSolrCoreNumDocs                         = "gauge.solr.core_num_docs"
	gaugeSolrCoreTotalspace                      = "gauge.solr.core_totalspace"
	gaugeSolrCoreUsablespace                     = "gauge.solr.core_usablespace"
	gaugeSolrDocumentCacheCumulativeHitratio     = "gauge.solr.document_cache_cumulative_hitratio"
	gaugeSolrFieldValueCacheCumulativeHitratio   = "gauge.solr.field_value_cache_cumulative_hitratio"
	gaugeSolrHTTPActiveRequests                  = "gauge.solr.http_active_requests"
	gaugeSolrJettyGetRequestLatency              = "gauge.solr.jetty_get_request_latency"
	gaugeSolrJettyPostRequestLatency             = "gauge.solr.jetty_post_request_latency"
	gaugeSolrJettyRequestLatency                 = "gauge.solr.jetty_request_latency"
	gaugeSolrJvmGcCmsCount                       = "gauge.solr.jvm_gc_cms_count"
	gaugeSolrJvmGcCmsTime                        = "gauge.solr.jvm_gc_cms_time"
	gaugeSolrJvmGcParnewCount                    = "gauge.solr.jvm_gc_parnew_count"
	gaugeSolrJvmGcParnewTime                     = "gauge.solr.jvm_gc_parnew_time"
	gaugeSolrJvmHeapUsage                        = "gauge.solr.jvm_heap_usage"
	gaugeSolrJvmMappedMemoryCapacity             = "gauge.solr.jvm_mapped_memory_capacity"
	gaugeSolrJvmMappedMemoryUsed                 = "gauge.solr.jvm_mapped_memory_used"
	gaugeSolrJvmMemoryPoolsCodeCacheUsage        = "gauge.solr.jvm_memory_pools_Code-Cache_usage"
	gaugeSolrJvmMemoryPoolsMetaspaceUsage        = "gauge.solr.jvm_memory_pools_Metaspace_usage"
	gaugeSolrJvmMemoryPoolsParEdenSpaceUsage     = "gauge.solr.jvm_memory_pools_Par-Eden-Space_usage"
	gaugeSolrJvmMemoryPoolsParSurvivorSpaceUsage = "gauge.solr.jvm_memory_pools_Par-Survivor-Space_usage"
	gaugeSolrJvmTotalMemory                      = "gauge.solr.jvm_total_memory"
	gaugeSolrJvmTotalMemoryUsed                  = "gauge.solr.jvm_total_memory_used"
	gaugeSolrNodeMetricRequestTime               = "gauge.solr.node_metric_request_time"
	gaugeSolrQueryResultCacheCumulativeHitratio  = "gauge.solr.query_result_cache_cumulative_hitratio"
	gaugeSolrReplicationHandlerResponse          = "gauge.solr.replication_handler_response"
	gaugeSolrSearchQueryResponse                 = "gauge.solr.search_query_response"
	gaugeSolrSearcherWarmup                      = "gauge.solr.searcher_warmup"
	gaugeSolrShardCumulativeDocs                 = "gauge.solr.shard_cumulative_docs"
	gaugeSolrUpdateRequestHandlerResponse        = "gauge.solr.update_request_handler_response"
	gaugeSolrZookeeperRequestTime                = "gauge.solr.zookeeper_request_time"
)

var metricSet = map[string]monitors.MetricInfo{
	counterSolrHTTP2xxResponses:                  {Type: datapoint.Count},
	counterSolrHTTP4xxResponses:                  {Type: datapoint.Count},
	counterSolrHTTP5xxResponses:                  {Type: datapoint.Count},
	counterSolrHTTPRequests:                      {Type: datapoint.Count},
	counterSolrJvmClassesLoaded:                  {Type: datapoint.Count},
	counterSolrNodeCollectionsRequests:           {Type: datapoint.Count},
	counterSolrNodeCoresRequests:                 {Type: datapoint.Count},
	counterSolrNodeMetricRequestCount:            {Type: datapoint.Count},
	counterSolrNodeMetricsRequests:               {Type: datapoint.Count},
	counterSolrNodeZookeeperRequests:             {Type: datapoint.Count},
	counterSolrOpenFileDescriptorCount:           {Type: datapoint.Count},
	counterSolrReplicationHandlerRequests:        {Type: datapoint.Count},
	counterSolrSearchQueryRequests:               {Type: datapoint.Count},
	counterSolrUpdateHandlerRequests:             {Type: datapoint.Count},
	counterSolrZookeeperErrors:                   {Type: datapoint.Count},
	gaugeSolrCoreDeletedDocs:                     {Type: datapoint.Gauge},
	gaugeSolrCoreIndexSize:                       {Type: datapoint.Gauge},
	gaugeSolrCoreMaxDocs:                         {Type: datapoint.Gauge},
	gaugeSolrCoreNumDocs:                         {Type: datapoint.Gauge},
	gaugeSolrCoreTotalspace:                      {Type: datapoint.Gauge},
	gaugeSolrCoreUsablespace:                     {Type: datapoint.Gauge},
	gaugeSolrDocumentCacheCumulativeHitratio:     {Type: datapoint.Gauge},
	gaugeSolrFieldValueCacheCumulativeHitratio:   {Type: datapoint.Gauge},
	gaugeSolrHTTPActiveRequests:                  {Type: datapoint.Gauge},
	gaugeSolrJettyGetRequestLatency:              {Type: datapoint.Gauge},
	gaugeSolrJettyPostRequestLatency:             {Type: datapoint.Gauge},
	gaugeSolrJettyRequestLatency:                 {Type: datapoint.Gauge},
	gaugeSolrJvmGcCmsCount:                       {Type: datapoint.Gauge},
	gaugeSolrJvmGcCmsTime:                        {Type: datapoint.Gauge},
	gaugeSolrJvmGcParnewCount:                    {Type: datapoint.Gauge},
	gaugeSolrJvmGcParnewTime:                     {Type: datapoint.Gauge},
	gaugeSolrJvmHeapUsage:                        {Type: datapoint.Gauge},
	gaugeSolrJvmMappedMemoryCapacity:             {Type: datapoint.Gauge},
	gaugeSolrJvmMappedMemoryUsed:                 {Type: datapoint.Gauge},
	gaugeSolrJvmMemoryPoolsCodeCacheUsage:        {Type: datapoint.Gauge},
	gaugeSolrJvmMemoryPoolsMetaspaceUsage:        {Type: datapoint.Gauge},
	gaugeSolrJvmMemoryPoolsParEdenSpaceUsage:     {Type: datapoint.Gauge},
	gaugeSolrJvmMemoryPoolsParSurvivorSpaceUsage: {Type: datapoint.Gauge},
	gaugeSolrJvmTotalMemory:                      {Type: datapoint.Gauge},
	gaugeSolrJvmTotalMemoryUsed:                  {Type: datapoint.Gauge},
	gaugeSolrNodeMetricRequestTime:               {Type: datapoint.Gauge},
	gaugeSolrQueryResultCacheCumulativeHitratio:  {Type: datapoint.Gauge},
	gaugeSolrReplicationHandlerResponse:          {Type: datapoint.Gauge},
	gaugeSolrSearchQueryResponse:                 {Type: datapoint.Gauge},
	gaugeSolrSearcherWarmup:                      {Type: datapoint.Gauge},
	gaugeSolrShardCumulativeDocs:                 {Type: datapoint.Gauge},
	gaugeSolrUpdateRequestHandlerResponse:        {Type: datapoint.Gauge},
	gaugeSolrZookeeperRequestTime:                {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	counterSolrHTTP2xxResponses:                  true,
	counterSolrHTTP4xxResponses:                  true,
	counterSolrHTTP5xxResponses:                  true,
	counterSolrHTTPRequests:                      true,
	counterSolrNodeCollectionsRequests:           true,
	counterSolrNodeCoresRequests:                 true,
	counterSolrNodeMetricsRequests:               true,
	counterSolrNodeZookeeperRequests:             true,
	counterSolrSearchQueryRequests:               true,
	counterSolrUpdateHandlerRequests:             true,
	gaugeSolrCoreDeletedDocs:                     true,
	gaugeSolrCoreIndexSize:                       true,
	gaugeSolrCoreMaxDocs:                         true,
	gaugeSolrCoreNumDocs:                         true,
	gaugeSolrCoreTotalspace:                      true,
	gaugeSolrCoreUsablespace:                     true,
	gaugeSolrDocumentCacheCumulativeHitratio:     true,
	gaugeSolrFieldValueCacheCumulativeHitratio:   true,
	gaugeSolrJettyRequestLatency:                 true,
	gaugeSolrJvmGcCmsCount:                       true,
	gaugeSolrJvmGcCmsTime:                        true,
	gaugeSolrJvmGcParnewCount:                    true,
	gaugeSolrJvmGcParnewTime:                     true,
	gaugeSolrJvmHeapUsage:                        true,
	gaugeSolrJvmMemoryPoolsCodeCacheUsage:        true,
	gaugeSolrJvmMemoryPoolsMetaspaceUsage:        true,
	gaugeSolrJvmMemoryPoolsParEdenSpaceUsage:     true,
	gaugeSolrJvmMemoryPoolsParSurvivorSpaceUsage: true,
	gaugeSolrJvmTotalMemory:                      true,
	gaugeSolrJvmTotalMemoryUsed:                  true,
	gaugeSolrQueryResultCacheCumulativeHitratio:  true,
	gaugeSolrSearchQueryResponse:                 true,
	gaugeSolrSearcherWarmup:                      true,
	gaugeSolrUpdateRequestHandlerResponse:        true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "collectd/solr",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     true,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
