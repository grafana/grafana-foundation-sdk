// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package news

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type Options struct {
	// empty/missing will default to grafana blog
	FeedUrl   *string `json:"feedUrl,omitempty"`
	ShowImage *bool   `json:"showImage,omitempty"`
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "news",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
