// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package logs

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Options struct {
	ShowLabels         bool                     `json:"showLabels"`
	ShowCommonLabels   bool                     `json:"showCommonLabels"`
	ShowTime           bool                     `json:"showTime"`
	WrapLogMessage     bool                     `json:"wrapLogMessage"`
	PrettifyLogMessage bool                     `json:"prettifyLogMessage"`
	EnableLogDetails   bool                     `json:"enableLogDetails"`
	SortOrder          common.LogsSortOrder     `json:"sortOrder"`
	DedupStrategy      common.LogsDedupStrategy `json:"dedupStrategy"`
}

func (resource *Options) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "showLabels"
	if fields["showLabels"] != nil {
		if string(fields["showLabels"]) != "null" {
			if err := json.Unmarshal(fields["showLabels"], &resource.ShowLabels); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showLabels", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showLabels", errors.New("required field is null"))...)

		}
		delete(fields, "showLabels")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showLabels", errors.New("required field is missing from input"))...)
	}
	// Field "showCommonLabels"
	if fields["showCommonLabels"] != nil {
		if string(fields["showCommonLabels"]) != "null" {
			if err := json.Unmarshal(fields["showCommonLabels"], &resource.ShowCommonLabels); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showCommonLabels", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showCommonLabels", errors.New("required field is null"))...)

		}
		delete(fields, "showCommonLabels")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showCommonLabels", errors.New("required field is missing from input"))...)
	}
	// Field "showTime"
	if fields["showTime"] != nil {
		if string(fields["showTime"]) != "null" {
			if err := json.Unmarshal(fields["showTime"], &resource.ShowTime); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showTime", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showTime", errors.New("required field is null"))...)

		}
		delete(fields, "showTime")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showTime", errors.New("required field is missing from input"))...)
	}
	// Field "wrapLogMessage"
	if fields["wrapLogMessage"] != nil {
		if string(fields["wrapLogMessage"]) != "null" {
			if err := json.Unmarshal(fields["wrapLogMessage"], &resource.WrapLogMessage); err != nil {
				errs = append(errs, cog.MakeBuildErrors("wrapLogMessage", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("wrapLogMessage", errors.New("required field is null"))...)

		}
		delete(fields, "wrapLogMessage")
	} else {
		errs = append(errs, cog.MakeBuildErrors("wrapLogMessage", errors.New("required field is missing from input"))...)
	}
	// Field "prettifyLogMessage"
	if fields["prettifyLogMessage"] != nil {
		if string(fields["prettifyLogMessage"]) != "null" {
			if err := json.Unmarshal(fields["prettifyLogMessage"], &resource.PrettifyLogMessage); err != nil {
				errs = append(errs, cog.MakeBuildErrors("prettifyLogMessage", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("prettifyLogMessage", errors.New("required field is null"))...)

		}
		delete(fields, "prettifyLogMessage")
	} else {
		errs = append(errs, cog.MakeBuildErrors("prettifyLogMessage", errors.New("required field is missing from input"))...)
	}
	// Field "enableLogDetails"
	if fields["enableLogDetails"] != nil {
		if string(fields["enableLogDetails"]) != "null" {
			if err := json.Unmarshal(fields["enableLogDetails"], &resource.EnableLogDetails); err != nil {
				errs = append(errs, cog.MakeBuildErrors("enableLogDetails", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("enableLogDetails", errors.New("required field is null"))...)

		}
		delete(fields, "enableLogDetails")
	} else {
		errs = append(errs, cog.MakeBuildErrors("enableLogDetails", errors.New("required field is missing from input"))...)
	}
	// Field "sortOrder"
	if fields["sortOrder"] != nil {
		if string(fields["sortOrder"]) != "null" {
			if err := json.Unmarshal(fields["sortOrder"], &resource.SortOrder); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sortOrder", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("sortOrder", errors.New("required field is null"))...)

		}
		delete(fields, "sortOrder")
	} else {
		errs = append(errs, cog.MakeBuildErrors("sortOrder", errors.New("required field is missing from input"))...)
	}
	// Field "dedupStrategy"
	if fields["dedupStrategy"] != nil {
		if string(fields["dedupStrategy"]) != "null" {
			if err := json.Unmarshal(fields["dedupStrategy"], &resource.DedupStrategy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dedupStrategy", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("dedupStrategy", errors.New("required field is null"))...)

		}
		delete(fields, "dedupStrategy")
	} else {
		errs = append(errs, cog.MakeBuildErrors("dedupStrategy", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Options", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (resource Options) Equals(other Options) bool {
	if resource.ShowLabels != other.ShowLabels {
		return false
	}
	if resource.ShowCommonLabels != other.ShowCommonLabels {
		return false
	}
	if resource.ShowTime != other.ShowTime {
		return false
	}
	if resource.WrapLogMessage != other.WrapLogMessage {
		return false
	}
	if resource.PrettifyLogMessage != other.PrettifyLogMessage {
		return false
	}
	if resource.EnableLogDetails != other.EnableLogDetails {
		return false
	}
	if resource.SortOrder != other.SortOrder {
		return false
	}
	if resource.DedupStrategy != other.DedupStrategy {
		return false
	}

	return true
}

// Validate checks any constraint that may be defined for this type
// and returns all violations.
func (resource Options) Validate() error {
	return nil
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "logs",
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
