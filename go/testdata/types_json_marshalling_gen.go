// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

func VariantConfig() cogvariants.DataqueryConfig {
	return cogvariants.DataqueryConfig{
		Identifier: "",
		DataqueryUnmarshaler: func(raw []byte) (cogvariants.Dataquery, error) {
			dataquery := Dataquery{}

			if err := json.Unmarshal(raw, &dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
	}
}
