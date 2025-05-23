// Copyright OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package smartagentextension

import (
	"fmt"
	"path/filepath"

	"github.com/signalfx/defaults"
	"go.opentelemetry.io/collector/confmap"
	"gopkg.in/yaml.v2"

	saconfig "github.com/signalfx/signalfx-agent/pkg/core/config"
)

var _ confmap.Unmarshaler = (*Config)(nil)

type Config struct {
	// Agent uses yaml, which mapstructure doesn't support.
	// Custom unmarshaller required for yaml and SFx defaults usage.
	saconfig.Config `mapstructure:"-,squash"`
}

func (cfg *Config) Unmarshal(componentParser *confmap.Conf) error {
	allSettings := componentParser.ToStringMap()

	configDirSet := false
	if collectd, ok := allSettings["collectd"]; ok {
		if collectdBlock, ok := collectd.(map[string]any); ok {
			if _, ok := collectdBlock["configDir"]; ok {
				configDirSet = true
			}
		}
	}

	config, err := smartAgentConfigFromSettingsMap(allSettings)
	if err != nil {
		return err
	}

	if config.BundleDir == "" {
		config.BundleDir = cfg.Config.BundleDir
	}
	config.Collectd.BundleDir = config.BundleDir

	if !configDirSet {
		config.Collectd.ConfigDir = filepath.Join(config.Collectd.BundleDir, "run", "collectd")
	}

	cfg.Config = *config
	return nil
}

func smartAgentConfigFromSettingsMap(settings map[string]any) (*saconfig.Config, error) {
	var config saconfig.Config
	var collectdSettings map[string]any
	var ok bool
	if collectdSettings, ok = settings["collectd"].(map[string]any); !ok {
		collectdSettings = map[string]any{}
	}

	settings["collectd"] = collectdSettings

	asBytes, err := yaml.Marshal(settings)
	if err != nil {
		return nil, fmt.Errorf("failed constructing raw Smart Agent config: %w", err)
	}

	err = yaml.UnmarshalStrict(asBytes, &config)
	if err != nil {
		return nil, fmt.Errorf("failed creating Smart Agent config: %w", err)
	}

	err = defaults.Set(&config)
	if err != nil {
		return nil, fmt.Errorf("failed setting config defaults: %w", err)
	}

	// The default on CollectdConfig is 0, use the default if this is the case.
	if config.Collectd.IntervalSeconds == 0 {
		config.Collectd.IntervalSeconds = defaultIntervalSeconds
	}

	config.Collectd.BundleDir = config.BundleDir
	return &config, nil
}
