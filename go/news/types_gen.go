// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package news

import (
	"encoding/json"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Options struct {
	// empty/missing will default to grafana blog
	FeedUrl   *string `json:"feedUrl,omitempty"`
	ShowImage *bool   `json:"showImage,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Options` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Options) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "feedUrl"
	if fields["feedUrl"] != nil {
		if string(fields["feedUrl"]) != "null" {
			if err := json.Unmarshal(fields["feedUrl"], &resource.FeedUrl); err != nil {
				errs = append(errs, cog.MakeBuildErrors("feedUrl", err)...)
			}

		}
		delete(fields, "feedUrl")

	}
	// Field "showImage"
	if fields["showImage"] != nil {
		if string(fields["showImage"]) != "null" {
			if err := json.Unmarshal(fields["showImage"], &resource.ShowImage); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showImage", err)...)
			}

		}
		delete(fields, "showImage")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Options", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Options` objects.
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

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	return nil
}

// VariantConfig returns the configuration related to news panels.
// This configuration describes how to unmarshal it, convert it to code, …
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
		StrictOptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := options.UnmarshalJSONStrict(raw); err != nil {
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
