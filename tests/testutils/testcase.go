// Copyright Splunk, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutils

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type CollectorBuilder func(Collector) Collector

// A Testcase is a central helper utility to provide Container, OTLPReceiverSink, ResourceMetrics,
// SplunkOtelCollector, and ObservedLogs to integration tests with minimal boilerplate.  It also embeds testing.TB
// for easy testing and testify usage.
type Testcase struct {
	testing.TB
	Logger                              *zap.Logger
	ObservedLogs                        *observer.ObservedLogs
	OTLPReceiverSink                    *OTLPReceiverSink
	OTLPEndpoint                        string
	OTLPEndpointForCollector            string
	ID                                  string
	OTLPReceiverShouldBindAllInterfaces bool
}

// NewTestcase is the recommended constructor that will automatically configure an OTLPReceiverSink
// with available endpoint and ObservedLogs.
func NewTestcase(t testing.TB) *Testcase {
	tc := Testcase{TB: t}
	var logCore zapcore.Core
	logCore, tc.ObservedLogs = observer.New(zap.DebugLevel)
	tc.Logger = zap.New(logCore)

	tc.setOTLPEndpoint()
	var err error
	tc.OTLPReceiverSink, err = NewOTLPReceiverSink().WithEndpoint(tc.OTLPEndpoint).Build()
	require.NoError(tc, err)
	require.NoError(tc, tc.OTLPReceiverSink.Start())

	id, err := uuid.NewRandom()
	require.NoError(tc, err)
	tc.ID = id.String()
	return &tc
}

func (t *Testcase) setOTLPEndpoint() {
	otlpPort := GetAvailablePort(t)
	otlpHost := "localhost"
	t.OTLPEndpoint = fmt.Sprintf("%s:%d", otlpHost, otlpPort)
	t.OTLPEndpointForCollector = t.OTLPEndpoint
}

// Builds and starts all provided Container builder instances, returning them and a validating stop function.
func (t *Testcase) Containers(builders ...Container) (containers []*Container, stop func()) {
	for _, builder := range builders {
		containers = append(containers, builder.Build())
	}

	for _, container := range containers {
		assert.NoError(t, container.Start(context.Background()))
	}

	stop = func() {
		for _, container := range containers {
			assert.NoError(t, container.Stop(context.Background(), nil))
			assert.NoError(t, container.Terminate(context.Background()))
		}
	}

	return
}

// SplunkOtelCollector builds and starts a collector container or process using the desired config filename
// (assuming it's in the ./testdata directory) returning it and a validating shutdown function.
func (t *Testcase) SplunkOtelCollector(configFilename string, builders ...CollectorBuilder) (collector Collector, shutdown func()) {
	return t.splunkOtelCollector(configFilename, builders...)
}

// SplunkOtelCollectorContainer is the same as SplunkOtelCollector but returns *CollectorContainer.
// If SPLUNK_OTEL_COLLECTOR_IMAGE isn't set, tests that call this will be skipped.
func (t *Testcase) SplunkOtelCollectorContainer(configFilename string, builders ...CollectorBuilder) (collector *CollectorContainer, shutdown func()) {
	cc := NewCollectorContainer().WithImage(GetCollectorImageOrSkipTest(t))
	// TODO why does darwin need special business logic?  Is this a proxy to say "is local dev"? Regardless need why comment
	if runtime.GOOS == "darwin" {
		port := strings.Split(t.OTLPEndpointForCollector, ":")[1]
		t.OTLPEndpointForCollector = fmt.Sprintf("host.docker.internal:%s", port)
	}

	var c Collector
	c, shutdown = t.newCollector(&cc, configFilename, builders...)
	return c.(*CollectorContainer), shutdown
}

// SplunkOtelCollectorProcess is the same as SplunkOtelCollector but returns *CollectorProcess.
func (t *Testcase) SplunkOtelCollectorProcess(configFilename string, builders ...CollectorBuilder) (collector *CollectorProcess, shutdown func()) {
	cp := NewCollectorProcess()

	var c Collector
	c, shutdown = t.newCollector(&cp, configFilename, builders...)
	return c.(*CollectorProcess), shutdown
}

func (t *Testcase) splunkOtelCollector(configFilename string, builders ...CollectorBuilder) (collector Collector, shutdown func()) {
	if image := os.Getenv("SPLUNK_OTEL_COLLECTOR_IMAGE"); strings.TrimSpace(image) != "" {
		return t.SplunkOtelCollectorContainer(configFilename, builders...)
	}
	return t.SplunkOtelCollectorProcess(configFilename, builders...)
}

func (t *Testcase) newCollector(initial Collector, configFilename string, builders ...CollectorBuilder) (collector Collector, shutdown func()) {
	collector = initial

	coverDest := os.Getenv("CONTAINER_COVER_DEST")
	coverSrc := os.Getenv("CONTAINER_COVER_SRC")

	envVars := map[string]string{
		"GOCOVERDIR":     coverDest,
		"OTLP_ENDPOINT":  t.OTLPEndpointForCollector,
		"SPLUNK_TEST_ID": t.ID,
	}

	if configFilename != "" {
		collector = collector.WithConfigPath(
			path.Join(".", "testdata", configFilename),
		)
	}

	collector = collector.WithEnv(envVars).WithLogLevel("debug").WithLogger(t.Logger)

	for _, builder := range builders {
		collector = builder(collector)
	}

	splunkEnv := map[string]string{}
	for _, s := range os.Environ() {
		split := strings.Split(s, "=")
		if strings.HasPrefix(strings.ToUpper(split[0]), "SPLUNK_") {
			splunkEnv[split[0]] = split[1]
		}
	}
	collector = collector.WithEnv(splunkEnv)

	if coverSrc != "" && coverDest != "" {
		collector = collector.WithMount(coverSrc, coverDest)
	}

	var err error
	collector, err = collector.Build()
	require.NoError(t, err)
	require.NotNil(t, collector)
	require.NoError(t, collector.Start())

	return collector, func() { require.NoError(t, collector.Shutdown()) }
}

// PrintLogsOnFailure will print all ObserverLogs messages if the test has failed.  It's intended to be
// deferred after Testcase creation.
// There is a bug in testcontainers-go so it's not certain these are complete:
// https://github.com/testcontainers/testcontainers-go/pull/323
func (t *Testcase) PrintLogsOnFailure() {
	if !t.Failed() {
		return
	}
	fmt.Printf("Logs: \n")
	for _, statement := range t.ObservedLogs.All() {
		fmt.Printf("%v\n", statement)
	}
}

// Validating shutdown helper for the Testcase's OTLPReceiverSink
func (t *Testcase) ShutdownOTLPReceiverSink() {
	require.NoError(t, t.OTLPReceiverSink.Shutdown())
}

func CheckMetricsPresence(t *testing.T, metricNames []string, configFile string) {
	tc := NewTestcase(t)
	defer tc.PrintLogsOnFailure()
	defer tc.ShutdownOTLPReceiverSink()

	_, shutdown := tc.SplunkOtelCollectorContainer(configFile)
	tc.Cleanup(shutdown)

	missingMetrics := make(map[string]any, len(metricNames))
	for _, m := range metricNames {
		missingMetrics[m] = struct{}{}
	}

	assert.EventuallyWithT(tc, func(tt *assert.CollectT) {
		for i := 0; i < len(tc.OTLPReceiverSink.AllMetrics()); i++ {
			m := tc.OTLPReceiverSink.AllMetrics()[i]
			for j := 0; j < m.ResourceMetrics().Len(); j++ {
				rm := m.ResourceMetrics().At(j)
				for k := 0; k < rm.ScopeMetrics().Len(); k++ {
					sm := rm.ScopeMetrics().At(k)
					for l := 0; l < sm.Metrics().Len(); l++ {
						delete(missingMetrics, sm.Metrics().At(l).Name())
					}
				}
			}
		}
		msg := "Missing metrics:\n"
		for k := range missingMetrics {
			msg += fmt.Sprintf("- %q\n", k)
		}
		assert.Len(tt, missingMetrics, 0, msg)
	}, 1*time.Minute, 1*time.Second)
}
