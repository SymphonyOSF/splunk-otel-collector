// Code generated by monitor-code-gen. DO NOT EDIT.

package logstash

import (
	"github.com/signalfx/golib/v3/datapoint"

	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "logstash"

const (
	groupEvents     = "events"
	groupHotThreads = "hot_threads"
	groupJvm        = "jvm"
	groupOs         = "os"
	groupPipeline   = "pipeline"
	groupPlugins    = "plugins"
	groupProcess    = "process"
	groupReloads    = "reloads"
)

var groupSet = map[string]bool{
	groupEvents:     true,
	groupHotThreads: true,
	groupJvm:        true,
	groupOs:         true,
	groupPipeline:   true,
	groupPlugins:    true,
	groupProcess:    true,
	groupReloads:    true,
}

const (
	nodeHotThreadsHotThreadsBusiestThreads                         = "node.hot_threads.hot_threads.busiest_threads"
	nodeHotThreadsHotThreadsThreadsPercentOfCPUTime                = "node.hot_threads.hot_threads.threads.percent_of_cpu_time"
	nodeOsOsAvailableProcessors                                    = "node.os.os.available_processors"
	nodePipelinesBatchDelay                                        = "node.pipelines.batch_delay"
	nodePipelinesBatchSize                                         = "node.pipelines.batch_size"
	nodePipelinesWorkers                                           = "node.pipelines.workers"
	nodePluginsTotal                                               = "node.plugins.total"
	nodeStatsEventsEventsDurationInMillis                          = "node.stats.events.events.duration_in_millis"
	nodeStatsEventsEventsFiltered                                  = "node.stats.events.events.filtered"
	nodeStatsEventsEventsIn                                        = "node.stats.events.events.in"
	nodeStatsEventsEventsOut                                       = "node.stats.events.events.out"
	nodeStatsEventsEventsQueuePushDurationInMillis                 = "node.stats.events.events.queue_push_duration_in_millis"
	nodeStatsJvmJvmGcCollectorsOldCollectionCount                  = "node.stats.jvm.jvm.gc.collectors.old.collection_count"
	nodeStatsJvmJvmGcCollectorsOldCollectionTimeInMillis           = "node.stats.jvm.jvm.gc.collectors.old.collection_time_in_millis"
	nodeStatsJvmJvmGcCollectorsYoungCollectionCount                = "node.stats.jvm.jvm.gc.collectors.young.collection_count"
	nodeStatsJvmJvmGcCollectorsYoungCollectionTimeInMillis         = "node.stats.jvm.jvm.gc.collectors.young.collection_time_in_millis"
	nodeStatsJvmJvmMemHeapCommittedInBytes                         = "node.stats.jvm.jvm.mem.heap_committed_in_bytes"
	nodeStatsJvmJvmMemHeapMaxInBytes                               = "node.stats.jvm.jvm.mem.heap_max_in_bytes"
	nodeStatsJvmJvmMemHeapUsedInBytes                              = "node.stats.jvm.jvm.mem.heap_used_in_bytes"
	nodeStatsJvmJvmMemHeapUsedPercent                              = "node.stats.jvm.jvm.mem.heap_used_percent"
	nodeStatsJvmJvmMemNonHeapCommittedInBytes                      = "node.stats.jvm.jvm.mem.non_heap_committed_in_bytes"
	nodeStatsJvmJvmMemNonHeapUsedInBytes                           = "node.stats.jvm.jvm.mem.non_heap_used_in_bytes"
	nodeStatsJvmJvmMemPoolsOldCommittedInBytes                     = "node.stats.jvm.jvm.mem.pools.old.committed_in_bytes"
	nodeStatsJvmJvmMemPoolsOldMaxInBytes                           = "node.stats.jvm.jvm.mem.pools.old.max_in_bytes"
	nodeStatsJvmJvmMemPoolsOldPeakMaxInBytes                       = "node.stats.jvm.jvm.mem.pools.old.peak_max_in_bytes"
	nodeStatsJvmJvmMemPoolsOldPeakUsedInBytes                      = "node.stats.jvm.jvm.mem.pools.old.peak_used_in_bytes"
	nodeStatsJvmJvmMemPoolsOldUsedInBytes                          = "node.stats.jvm.jvm.mem.pools.old.used_in_bytes"
	nodeStatsJvmJvmMemPoolsSurvivorCommittedInBytes                = "node.stats.jvm.jvm.mem.pools.survivor.committed_in_bytes"
	nodeStatsJvmJvmMemPoolsSurvivorMaxInBytes                      = "node.stats.jvm.jvm.mem.pools.survivor.max_in_bytes"
	nodeStatsJvmJvmMemPoolsSurvivorPeakMaxInBytes                  = "node.stats.jvm.jvm.mem.pools.survivor.peak_max_in_bytes"
	nodeStatsJvmJvmMemPoolsSurvivorPeakUsedInBytes                 = "node.stats.jvm.jvm.mem.pools.survivor.peak_used_in_bytes"
	nodeStatsJvmJvmMemPoolsSurvivorUsedInBytes                     = "node.stats.jvm.jvm.mem.pools.survivor.used_in_bytes"
	nodeStatsJvmJvmMemPoolsYoungCommittedInBytes                   = "node.stats.jvm.jvm.mem.pools.young.committed_in_bytes"
	nodeStatsJvmJvmMemPoolsYoungMaxInBytes                         = "node.stats.jvm.jvm.mem.pools.young.max_in_bytes"
	nodeStatsJvmJvmMemPoolsYoungPeakMaxInBytes                     = "node.stats.jvm.jvm.mem.pools.young.peak_max_in_bytes"
	nodeStatsJvmJvmMemPoolsYoungPeakUsedInBytes                    = "node.stats.jvm.jvm.mem.pools.young.peak_used_in_bytes"
	nodeStatsJvmJvmMemPoolsYoungUsedInBytes                        = "node.stats.jvm.jvm.mem.pools.young.used_in_bytes"
	nodeStatsJvmJvmThreadsCount                                    = "node.stats.jvm.jvm.threads.count"
	nodeStatsJvmJvmThreadsPeakCount                                = "node.stats.jvm.jvm.threads.peak_count"
	nodeStatsJvmJvmUptimeInMillis                                  = "node.stats.jvm.jvm.uptime_in_millis"
	nodeStatsOsOsCgroupCPUCfsPeriodMicros                          = "node.stats.os.os.cgroup.cpu.cfs_period_micros"
	nodeStatsOsOsCgroupCPUCfsQuotaMicros                           = "node.stats.os.os.cgroup.cpu.cfs_quota_micros"
	nodeStatsOsOsCgroupCPUStatNumberOfElapsedPeriods               = "node.stats.os.os.cgroup.cpu.stat.number_of_elapsed_periods"
	nodeStatsOsOsCgroupCPUStatNumberOfTimesThrottled               = "node.stats.os.os.cgroup.cpu.stat.number_of_times_throttled"
	nodeStatsOsOsCgroupCPUStatTimeThrottledNanos                   = "node.stats.os.os.cgroup.cpu.stat.time_throttled_nanos"
	nodeStatsOsOsCgroupCpuacctUsageNanos                           = "node.stats.os.os.cgroup.cpuacct.usage_nanos"
	nodeStatsPipelinesEventsDurationInMillis                       = "node.stats.pipelines.events.duration_in_millis"
	nodeStatsPipelinesEventsFiltered                               = "node.stats.pipelines.events.filtered"
	nodeStatsPipelinesEventsIn                                     = "node.stats.pipelines.events.in"
	nodeStatsPipelinesEventsOut                                    = "node.stats.pipelines.events.out"
	nodeStatsPipelinesEventsQueuePushDurationInMillis              = "node.stats.pipelines.events.queue_push_duration_in_millis"
	nodeStatsPipelinesPluginsCodecsDecodeDurationInMillis          = "node.stats.pipelines.plugins.codecs.decode.duration_in_millis"
	nodeStatsPipelinesPluginsCodecsDecodeOut                       = "node.stats.pipelines.plugins.codecs.decode.out"
	nodeStatsPipelinesPluginsCodecsDecodeWritesIn                  = "node.stats.pipelines.plugins.codecs.decode.writes_in"
	nodeStatsPipelinesPluginsCodecsEncodeDurationInMillis          = "node.stats.pipelines.plugins.codecs.encode.duration_in_millis"
	nodeStatsPipelinesPluginsCodecsEncodeWritesIn                  = "node.stats.pipelines.plugins.codecs.encode.writes_in"
	nodeStatsPipelinesPluginsFiltersEventsDurationInMillis         = "node.stats.pipelines.plugins.filters.events.duration_in_millis"
	nodeStatsPipelinesPluginsFiltersEventsIn                       = "node.stats.pipelines.plugins.filters.events.in"
	nodeStatsPipelinesPluginsFiltersEventsOut                      = "node.stats.pipelines.plugins.filters.events.out"
	nodeStatsPipelinesPluginsInputsEventsOut                       = "node.stats.pipelines.plugins.inputs.events.out"
	nodeStatsPipelinesPluginsInputsEventsQueuePushDurationInMillis = "node.stats.pipelines.plugins.inputs.events.queue_push_duration_in_millis"
	nodeStatsPipelinesPluginsOutputsEventsDurationInMillis         = "node.stats.pipelines.plugins.outputs.events.duration_in_millis"
	nodeStatsPipelinesPluginsOutputsEventsIn                       = "node.stats.pipelines.plugins.outputs.events.in"
	nodeStatsPipelinesPluginsOutputsEventsOut                      = "node.stats.pipelines.plugins.outputs.events.out"
	nodeStatsPipelinesQueueEventsCount                             = "node.stats.pipelines.queue.events_count"
	nodeStatsPipelinesQueueMaxQueueSizeInBytes                     = "node.stats.pipelines.queue.max_queue_size_in_bytes"
	nodeStatsPipelinesQueueQueueSizeInBytes                        = "node.stats.pipelines.queue.queue_size_in_bytes"
	nodeStatsPipelinesReloadsFailures                              = "node.stats.pipelines.reloads.failures"
	nodeStatsPipelinesReloadsSuccesses                             = "node.stats.pipelines.reloads.successes"
	nodeStatsProcessProcessCPULoadAverage15m                       = "node.stats.process.process.cpu.load_average.15m"
	nodeStatsProcessProcessCPULoadAverage1m                        = "node.stats.process.process.cpu.load_average.1m"
	nodeStatsProcessProcessCPULoadAverage5m                        = "node.stats.process.process.cpu.load_average.5m"
	nodeStatsProcessProcessCPUPercent                              = "node.stats.process.process.cpu.percent"
	nodeStatsProcessProcessCPUTotalInMillis                        = "node.stats.process.process.cpu.total_in_millis"
	nodeStatsProcessProcessMaxFileDescriptors                      = "node.stats.process.process.max_file_descriptors"
	nodeStatsProcessProcessMemTotalVirtualInBytes                  = "node.stats.process.process.mem.total_virtual_in_bytes"
	nodeStatsProcessProcessOpenFileDescriptors                     = "node.stats.process.process.open_file_descriptors"
	nodeStatsProcessProcessPeakOpenFileDescriptors                 = "node.stats.process.process.peak_open_file_descriptors"
	nodeStatsReloadsReloadsFailures                                = "node.stats.reloads.reloads.failures"
	nodeStatsReloadsReloadsSuccesses                               = "node.stats.reloads.reloads.successes"
)

var metricSet = map[string]monitors.MetricInfo{
	nodeHotThreadsHotThreadsBusiestThreads:                         {Type: datapoint.Gauge, Group: groupHotThreads},
	nodeHotThreadsHotThreadsThreadsPercentOfCPUTime:                {Type: datapoint.Gauge, Group: groupHotThreads},
	nodeOsOsAvailableProcessors:                                    {Type: datapoint.Gauge, Group: groupOs},
	nodePipelinesBatchDelay:                                        {Type: datapoint.Gauge, Group: groupPipeline},
	nodePipelinesBatchSize:                                         {Type: datapoint.Gauge, Group: groupPipeline},
	nodePipelinesWorkers:                                           {Type: datapoint.Gauge, Group: groupPipeline},
	nodePluginsTotal:                                               {Type: datapoint.Gauge, Group: groupPlugins},
	nodeStatsEventsEventsDurationInMillis:                          {Type: datapoint.Counter, Group: groupEvents},
	nodeStatsEventsEventsFiltered:                                  {Type: datapoint.Counter, Group: groupEvents},
	nodeStatsEventsEventsIn:                                        {Type: datapoint.Counter, Group: groupEvents},
	nodeStatsEventsEventsOut:                                       {Type: datapoint.Counter, Group: groupEvents},
	nodeStatsEventsEventsQueuePushDurationInMillis:                 {Type: datapoint.Counter, Group: groupEvents},
	nodeStatsJvmJvmGcCollectorsOldCollectionCount:                  {Type: datapoint.Counter, Group: groupJvm},
	nodeStatsJvmJvmGcCollectorsOldCollectionTimeInMillis:           {Type: datapoint.Counter, Group: groupJvm},
	nodeStatsJvmJvmGcCollectorsYoungCollectionCount:                {Type: datapoint.Counter, Group: groupJvm},
	nodeStatsJvmJvmGcCollectorsYoungCollectionTimeInMillis:         {Type: datapoint.Counter, Group: groupJvm},
	nodeStatsJvmJvmMemHeapCommittedInBytes:                         {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemHeapMaxInBytes:                               {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemHeapUsedInBytes:                              {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemHeapUsedPercent:                              {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemNonHeapCommittedInBytes:                      {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemNonHeapUsedInBytes:                           {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsOldCommittedInBytes:                     {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsOldMaxInBytes:                           {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsOldPeakMaxInBytes:                       {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsOldPeakUsedInBytes:                      {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsOldUsedInBytes:                          {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsSurvivorCommittedInBytes:                {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsSurvivorMaxInBytes:                      {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsSurvivorPeakMaxInBytes:                  {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsSurvivorPeakUsedInBytes:                 {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsSurvivorUsedInBytes:                     {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsYoungCommittedInBytes:                   {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsYoungMaxInBytes:                         {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsYoungPeakMaxInBytes:                     {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsYoungPeakUsedInBytes:                    {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmMemPoolsYoungUsedInBytes:                        {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmThreadsCount:                                    {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmThreadsPeakCount:                                {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsJvmJvmUptimeInMillis:                                  {Type: datapoint.Gauge, Group: groupJvm},
	nodeStatsOsOsCgroupCPUCfsPeriodMicros:                          {Type: datapoint.Gauge, Group: groupOs},
	nodeStatsOsOsCgroupCPUCfsQuotaMicros:                           {Type: datapoint.Gauge, Group: groupOs},
	nodeStatsOsOsCgroupCPUStatNumberOfElapsedPeriods:               {Type: datapoint.Counter, Group: groupOs},
	nodeStatsOsOsCgroupCPUStatNumberOfTimesThrottled:               {Type: datapoint.Counter, Group: groupOs},
	nodeStatsOsOsCgroupCPUStatTimeThrottledNanos:                   {Type: datapoint.Counter, Group: groupOs},
	nodeStatsOsOsCgroupCpuacctUsageNanos:                           {Type: datapoint.Gauge, Group: groupOs},
	nodeStatsPipelinesEventsDurationInMillis:                       {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesEventsFiltered:                               {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesEventsIn:                                     {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesEventsOut:                                    {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesEventsQueuePushDurationInMillis:              {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesPluginsCodecsDecodeDurationInMillis:          {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesPluginsCodecsDecodeOut:                       {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesPluginsCodecsDecodeWritesIn:                  {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesPluginsCodecsEncodeDurationInMillis:          {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesPluginsCodecsEncodeWritesIn:                  {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesPluginsFiltersEventsDurationInMillis:         {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesPluginsFiltersEventsIn:                       {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesPluginsFiltersEventsOut:                      {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesPluginsInputsEventsOut:                       {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesPluginsInputsEventsQueuePushDurationInMillis: {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesPluginsOutputsEventsDurationInMillis:         {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesPluginsOutputsEventsIn:                       {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesPluginsOutputsEventsOut:                      {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesQueueEventsCount:                             {Type: datapoint.Gauge, Group: groupPipeline},
	nodeStatsPipelinesQueueMaxQueueSizeInBytes:                     {Type: datapoint.Gauge, Group: groupPipeline},
	nodeStatsPipelinesQueueQueueSizeInBytes:                        {Type: datapoint.Gauge, Group: groupPipeline},
	nodeStatsPipelinesReloadsFailures:                              {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsPipelinesReloadsSuccesses:                             {Type: datapoint.Counter, Group: groupPipeline},
	nodeStatsProcessProcessCPULoadAverage15m:                       {Type: datapoint.Gauge, Group: groupProcess},
	nodeStatsProcessProcessCPULoadAverage1m:                        {Type: datapoint.Gauge, Group: groupProcess},
	nodeStatsProcessProcessCPULoadAverage5m:                        {Type: datapoint.Gauge, Group: groupProcess},
	nodeStatsProcessProcessCPUPercent:                              {Type: datapoint.Gauge, Group: groupProcess},
	nodeStatsProcessProcessCPUTotalInMillis:                        {Type: datapoint.Counter, Group: groupProcess},
	nodeStatsProcessProcessMaxFileDescriptors:                      {Type: datapoint.Gauge, Group: groupProcess},
	nodeStatsProcessProcessMemTotalVirtualInBytes:                  {Type: datapoint.Gauge, Group: groupProcess},
	nodeStatsProcessProcessOpenFileDescriptors:                     {Type: datapoint.Gauge, Group: groupProcess},
	nodeStatsProcessProcessPeakOpenFileDescriptors:                 {Type: datapoint.Gauge, Group: groupProcess},
	nodeStatsReloadsReloadsFailures:                                {Type: datapoint.Counter, Group: groupReloads},
	nodeStatsReloadsReloadsSuccesses:                               {Type: datapoint.Counter, Group: groupReloads},
}

var defaultMetrics = map[string]bool{
	nodePluginsTotal:                                       true,
	nodeStatsEventsEventsDurationInMillis:                  true,
	nodeStatsEventsEventsFiltered:                          true,
	nodeStatsEventsEventsIn:                                true,
	nodeStatsEventsEventsOut:                               true,
	nodeStatsEventsEventsQueuePushDurationInMillis:         true,
	nodeStatsJvmJvmMemHeapCommittedInBytes:                 true,
	nodeStatsJvmJvmMemHeapMaxInBytes:                       true,
	nodeStatsJvmJvmMemHeapUsedInBytes:                      true,
	nodeStatsJvmJvmMemNonHeapUsedInBytes:                   true,
	nodeStatsJvmJvmThreadsCount:                            true,
	nodeStatsJvmJvmThreadsPeakCount:                        true,
	nodeStatsPipelinesEventsDurationInMillis:               true,
	nodeStatsPipelinesEventsFiltered:                       true,
	nodeStatsPipelinesEventsIn:                             true,
	nodeStatsPipelinesEventsOut:                            true,
	nodeStatsPipelinesPluginsFiltersEventsDurationInMillis: true,
	nodeStatsPipelinesPluginsFiltersEventsIn:               true,
	nodeStatsPipelinesPluginsFiltersEventsOut:              true,
	nodeStatsPipelinesPluginsInputsEventsOut:               true,
	nodeStatsPipelinesPluginsOutputsEventsDurationInMillis: true,
	nodeStatsPipelinesPluginsOutputsEventsIn:               true,
	nodeStatsPipelinesPluginsOutputsEventsOut:              true,
	nodeStatsProcessProcessCPUPercent:                      true,
}

var groupMetricsMap = map[string][]string{
	groupEvents: {
		nodeStatsEventsEventsDurationInMillis,
		nodeStatsEventsEventsFiltered,
		nodeStatsEventsEventsIn,
		nodeStatsEventsEventsOut,
		nodeStatsEventsEventsQueuePushDurationInMillis,
	},
	groupHotThreads: {
		nodeHotThreadsHotThreadsBusiestThreads,
		nodeHotThreadsHotThreadsThreadsPercentOfCPUTime,
	},
	groupJvm: {
		nodeStatsJvmJvmGcCollectorsOldCollectionCount,
		nodeStatsJvmJvmGcCollectorsOldCollectionTimeInMillis,
		nodeStatsJvmJvmGcCollectorsYoungCollectionCount,
		nodeStatsJvmJvmGcCollectorsYoungCollectionTimeInMillis,
		nodeStatsJvmJvmMemHeapCommittedInBytes,
		nodeStatsJvmJvmMemHeapMaxInBytes,
		nodeStatsJvmJvmMemHeapUsedInBytes,
		nodeStatsJvmJvmMemHeapUsedPercent,
		nodeStatsJvmJvmMemNonHeapCommittedInBytes,
		nodeStatsJvmJvmMemNonHeapUsedInBytes,
		nodeStatsJvmJvmMemPoolsOldCommittedInBytes,
		nodeStatsJvmJvmMemPoolsOldMaxInBytes,
		nodeStatsJvmJvmMemPoolsOldPeakMaxInBytes,
		nodeStatsJvmJvmMemPoolsOldPeakUsedInBytes,
		nodeStatsJvmJvmMemPoolsOldUsedInBytes,
		nodeStatsJvmJvmMemPoolsSurvivorCommittedInBytes,
		nodeStatsJvmJvmMemPoolsSurvivorMaxInBytes,
		nodeStatsJvmJvmMemPoolsSurvivorPeakMaxInBytes,
		nodeStatsJvmJvmMemPoolsSurvivorPeakUsedInBytes,
		nodeStatsJvmJvmMemPoolsSurvivorUsedInBytes,
		nodeStatsJvmJvmMemPoolsYoungCommittedInBytes,
		nodeStatsJvmJvmMemPoolsYoungMaxInBytes,
		nodeStatsJvmJvmMemPoolsYoungPeakMaxInBytes,
		nodeStatsJvmJvmMemPoolsYoungPeakUsedInBytes,
		nodeStatsJvmJvmMemPoolsYoungUsedInBytes,
		nodeStatsJvmJvmThreadsCount,
		nodeStatsJvmJvmThreadsPeakCount,
		nodeStatsJvmJvmUptimeInMillis,
	},
	groupOs: {
		nodeOsOsAvailableProcessors,
		nodeStatsOsOsCgroupCPUCfsPeriodMicros,
		nodeStatsOsOsCgroupCPUCfsQuotaMicros,
		nodeStatsOsOsCgroupCPUStatNumberOfElapsedPeriods,
		nodeStatsOsOsCgroupCPUStatNumberOfTimesThrottled,
		nodeStatsOsOsCgroupCPUStatTimeThrottledNanos,
		nodeStatsOsOsCgroupCpuacctUsageNanos,
	},
	groupPipeline: {
		nodePipelinesBatchDelay,
		nodePipelinesBatchSize,
		nodePipelinesWorkers,
		nodeStatsPipelinesEventsDurationInMillis,
		nodeStatsPipelinesEventsFiltered,
		nodeStatsPipelinesEventsIn,
		nodeStatsPipelinesEventsOut,
		nodeStatsPipelinesEventsQueuePushDurationInMillis,
		nodeStatsPipelinesPluginsCodecsDecodeDurationInMillis,
		nodeStatsPipelinesPluginsCodecsDecodeOut,
		nodeStatsPipelinesPluginsCodecsDecodeWritesIn,
		nodeStatsPipelinesPluginsCodecsEncodeDurationInMillis,
		nodeStatsPipelinesPluginsCodecsEncodeWritesIn,
		nodeStatsPipelinesPluginsFiltersEventsDurationInMillis,
		nodeStatsPipelinesPluginsFiltersEventsIn,
		nodeStatsPipelinesPluginsFiltersEventsOut,
		nodeStatsPipelinesPluginsInputsEventsOut,
		nodeStatsPipelinesPluginsInputsEventsQueuePushDurationInMillis,
		nodeStatsPipelinesPluginsOutputsEventsDurationInMillis,
		nodeStatsPipelinesPluginsOutputsEventsIn,
		nodeStatsPipelinesPluginsOutputsEventsOut,
		nodeStatsPipelinesQueueEventsCount,
		nodeStatsPipelinesQueueMaxQueueSizeInBytes,
		nodeStatsPipelinesQueueQueueSizeInBytes,
		nodeStatsPipelinesReloadsFailures,
		nodeStatsPipelinesReloadsSuccesses,
	},
	groupPlugins: {
		nodePluginsTotal,
	},
	groupProcess: {
		nodeStatsProcessProcessCPULoadAverage15m,
		nodeStatsProcessProcessCPULoadAverage1m,
		nodeStatsProcessProcessCPULoadAverage5m,
		nodeStatsProcessProcessCPUPercent,
		nodeStatsProcessProcessCPUTotalInMillis,
		nodeStatsProcessProcessMaxFileDescriptors,
		nodeStatsProcessProcessMemTotalVirtualInBytes,
		nodeStatsProcessProcessOpenFileDescriptors,
		nodeStatsProcessProcessPeakOpenFileDescriptors,
	},
	groupReloads: {
		nodeStatsReloadsReloadsFailures,
		nodeStatsReloadsReloadsSuccesses,
	},
}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "logstash",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
