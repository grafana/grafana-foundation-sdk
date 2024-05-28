// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package debug

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type UpdateConfig struct {
	Render        bool `json:"render"`
	DataChanged   bool `json:"dataChanged"`
	SchemaChanged bool `json:"schemaChanged"`
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

func VariantConfig() cogvariants.PanelcfgConfig {
	return cogvariants.PanelcfgConfig{
		Identifier: "debug",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
