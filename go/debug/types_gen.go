// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package debug

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type UpdateConfig struct {
	Render        bool `json:"render"`
	DataChanged   bool `json:"dataChanged"`
	SchemaChanged bool `json:"schemaChanged"`
}

func (resource UpdateConfig) Equals(other UpdateConfig) bool {
	if resource.Render != other.Render {
		return false
	}
	if resource.DataChanged != other.DataChanged {
		return false
	}
	if resource.SchemaChanged != other.SchemaChanged {
		return false
	}

	return true
}

type DebugMode string

const (
	DebugModeRender     DebugMode = "render"
	DebugModeEvents     DebugMode = "events"
	DebugModeCursor     DebugMode = "cursor"
	DebugModeState      DebugMode = "State"
	DebugModeThrowError DebugMode = "ThrowError"
)

type Options struct {
	Mode     DebugMode     `json:"mode"`
	Counters *UpdateConfig `json:"counters,omitempty"`
}

func (resource Options) Equals(other Options) bool {
	if resource.Mode != other.Mode {
		return false
	}
	if resource.Counters == nil && other.Counters != nil || resource.Counters != nil && other.Counters == nil {
		return false
	}

	if resource.Counters != nil {
		if !resource.Counters.Equals(*other.Counters) {
			return false
		}
	}

	return true
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "debug",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := json.Unmarshal(raw, options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
