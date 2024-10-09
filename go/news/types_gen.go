// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package news

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Options struct {
	// empty/missing will default to grafana blog
	FeedUrl   *string `json:"feedUrl,omitempty"`
	ShowImage *bool   `json:"showImage,omitempty"`
}

func (resource Options) Equals(other Options) bool {
	if resource.FeedUrl == nil && other.FeedUrl != nil || resource.FeedUrl != nil && other.FeedUrl == nil {
		return false
	}

	if resource.FeedUrl != nil {
		if *resource.FeedUrl != *other.FeedUrl {
			return false
		}
	}
	if resource.ShowImage == nil && other.ShowImage != nil || resource.ShowImage != nil && other.ShowImage == nil {
		return false
	}

	if resource.ShowImage != nil {
		if *resource.ShowImage != *other.ShowImage {
			return false
		}
	}

	return true
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "news",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := json.Unmarshal(raw, options); err != nil {
				return nil, err
			}

			return options, nil
		},
		GoConverter: func(inputPanel any) string {
			if panel, ok := inputPanel.(*dashboard.Panel); ok {
				return PanelConverter(*panel)
			}

			return PanelConverter(inputPanel.(dashboard.Panel))
		},
	}
}
