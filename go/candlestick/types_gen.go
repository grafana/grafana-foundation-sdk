// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package candlestick

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type VizDisplayMode string

const (
	VizDisplayModeCandlesVolume VizDisplayMode = "candles+volume"
	VizDisplayModeCandles       VizDisplayMode = "candles"
	VizDisplayModeVolume        VizDisplayMode = "volume"
)

type CandleStyle string

const (
	CandleStyleCandles  CandleStyle = "candles"
	CandleStyleOHLCBars CandleStyle = "ohlcbars"
)

type ColorStrategy string

const (
	ColorStrategyOpenClose  ColorStrategy = "open-close"
	ColorStrategyCloseClose ColorStrategy = "close-close"
)

type CandlestickFieldMap struct {
	// Corresponds to the starting value of the given period
	Open *string `json:"open,omitempty"`
	// Corresponds to the highest value of the given period
	High *string `json:"high,omitempty"`
	// Corresponds to the lowest value of the given period
	Low *string `json:"low,omitempty"`
	// Corresponds to the final (end) value of the given period
	Close *string `json:"close,omitempty"`
	// Corresponds to the sample count in the given period. (e.g. number of trades)
	Volume *string `json:"volume,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CandlestickFieldMap` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CandlestickFieldMap) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "open"
	if fields["open"] != nil {
		if string(fields["open"]) != "null" {
			if err := json.Unmarshal(fields["open"], &resource.Open); err != nil {
				errs = append(errs, cog.MakeBuildErrors("open", err)...)
			}

		}
		delete(fields, "open")

	}
	// Field "high"
	if fields["high"] != nil {
		if string(fields["high"]) != "null" {
			if err := json.Unmarshal(fields["high"], &resource.High); err != nil {
				errs = append(errs, cog.MakeBuildErrors("high", err)...)
			}

		}
		delete(fields, "high")

	}
	// Field "low"
	if fields["low"] != nil {
		if string(fields["low"]) != "null" {
			if err := json.Unmarshal(fields["low"], &resource.Low); err != nil {
				errs = append(errs, cog.MakeBuildErrors("low", err)...)
			}

		}
		delete(fields, "low")

	}
	// Field "close"
	if fields["close"] != nil {
		if string(fields["close"]) != "null" {
			if err := json.Unmarshal(fields["close"], &resource.Close); err != nil {
				errs = append(errs, cog.MakeBuildErrors("close", err)...)
			}

		}
		delete(fields, "close")

	}
	// Field "volume"
	if fields["volume"] != nil {
		if string(fields["volume"]) != "null" {
			if err := json.Unmarshal(fields["volume"], &resource.Volume); err != nil {
				errs = append(errs, cog.MakeBuildErrors("volume", err)...)
			}

		}
		delete(fields, "volume")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CandlestickFieldMap", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CandlestickFieldMap` objects.
func (resource CandlestickFieldMap) Equals(other CandlestickFieldMap) bool {
	if resource.Open == nil && other.Open != nil || resource.Open != nil && other.Open == nil {
		return false
	}

	if resource.Open != nil {
		if *resource.Open != *other.Open {
			return false
		}
	}
	if resource.High == nil && other.High != nil || resource.High != nil && other.High == nil {
		return false
	}

	if resource.High != nil {
		if *resource.High != *other.High {
			return false
		}
	}
	if resource.Low == nil && other.Low != nil || resource.Low != nil && other.Low == nil {
		return false
	}

	if resource.Low != nil {
		if *resource.Low != *other.Low {
			return false
		}
	}
	if resource.Close == nil && other.Close != nil || resource.Close != nil && other.Close == nil {
		return false
	}

	if resource.Close != nil {
		if *resource.Close != *other.Close {
			return false
		}
	}
	if resource.Volume == nil && other.Volume != nil || resource.Volume != nil && other.Volume == nil {
		return false
	}

	if resource.Volume != nil {
		if *resource.Volume != *other.Volume {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CandlestickFieldMap` fields for violations and returns them.
func (resource CandlestickFieldMap) Validate() error {
	return nil
}

type CandlestickColors struct {
	Up   string `json:"up"`
	Down string `json:"down"`
	Flat string `json:"flat"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CandlestickColors` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CandlestickColors) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "up"
	if fields["up"] != nil {
		if string(fields["up"]) != "null" {
			if err := json.Unmarshal(fields["up"], &resource.Up); err != nil {
				errs = append(errs, cog.MakeBuildErrors("up", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("up", errors.New("required field is null"))...)

		}
		delete(fields, "up")
	} else {
		errs = append(errs, cog.MakeBuildErrors("up", errors.New("required field is missing from input"))...)
	}
	// Field "down"
	if fields["down"] != nil {
		if string(fields["down"]) != "null" {
			if err := json.Unmarshal(fields["down"], &resource.Down); err != nil {
				errs = append(errs, cog.MakeBuildErrors("down", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("down", errors.New("required field is null"))...)

		}
		delete(fields, "down")
	} else {
		errs = append(errs, cog.MakeBuildErrors("down", errors.New("required field is missing from input"))...)
	}
	// Field "flat"
	if fields["flat"] != nil {
		if string(fields["flat"]) != "null" {
			if err := json.Unmarshal(fields["flat"], &resource.Flat); err != nil {
				errs = append(errs, cog.MakeBuildErrors("flat", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("flat", errors.New("required field is null"))...)

		}
		delete(fields, "flat")
	} else {
		errs = append(errs, cog.MakeBuildErrors("flat", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CandlestickColors", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CandlestickColors` objects.
func (resource CandlestickColors) Equals(other CandlestickColors) bool {
	if resource.Up != other.Up {
		return false
	}
	if resource.Down != other.Down {
		return false
	}
	if resource.Flat != other.Flat {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CandlestickColors` fields for violations and returns them.
func (resource CandlestickColors) Validate() error {
	return nil
}

type Options struct {
	// Sets which dimensions are used for the visualization
	Mode VizDisplayMode `json:"mode"`
	// Sets the style of the candlesticks
	CandleStyle CandleStyle `json:"candleStyle"`
	// Sets the color strategy for the candlesticks
	ColorStrategy ColorStrategy `json:"colorStrategy"`
	// Map fields to appropriate dimension
	Fields CandlestickFieldMap `json:"fields"`
	// Set which colors are used when the price movement is up or down
	Colors CandlestickColors       `json:"colors"`
	Legend common.VizLegendOptions `json:"legend"`
	// When enabled, all fields will be sent to the graph
	IncludeAllFields *bool `json:"includeAllFields,omitempty"`
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
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is null"))...)

		}
		delete(fields, "mode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is missing from input"))...)
	}
	// Field "candleStyle"
	if fields["candleStyle"] != nil {
		if string(fields["candleStyle"]) != "null" {
			if err := json.Unmarshal(fields["candleStyle"], &resource.CandleStyle); err != nil {
				errs = append(errs, cog.MakeBuildErrors("candleStyle", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("candleStyle", errors.New("required field is null"))...)

		}
		delete(fields, "candleStyle")
	} else {
		errs = append(errs, cog.MakeBuildErrors("candleStyle", errors.New("required field is missing from input"))...)
	}
	// Field "colorStrategy"
	if fields["colorStrategy"] != nil {
		if string(fields["colorStrategy"]) != "null" {
			if err := json.Unmarshal(fields["colorStrategy"], &resource.ColorStrategy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("colorStrategy", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("colorStrategy", errors.New("required field is null"))...)

		}
		delete(fields, "colorStrategy")
	} else {
		errs = append(errs, cog.MakeBuildErrors("colorStrategy", errors.New("required field is missing from input"))...)
	}
	// Field "fields"
	if fields["fields"] != nil {
		if string(fields["fields"]) != "null" {

			resource.Fields = CandlestickFieldMap{}
			if err := resource.Fields.UnmarshalJSONStrict(fields["fields"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fields", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("fields", errors.New("required field is null"))...)

		}
		delete(fields, "fields")
	} else {
		errs = append(errs, cog.MakeBuildErrors("fields", errors.New("required field is missing from input"))...)
	}
	// Field "colors"
	if fields["colors"] != nil {
		if string(fields["colors"]) != "null" {

			resource.Colors = CandlestickColors{}
			if err := resource.Colors.UnmarshalJSONStrict(fields["colors"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("colors", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("colors", errors.New("required field is null"))...)

		}
		delete(fields, "colors")
	} else {
		errs = append(errs, cog.MakeBuildErrors("colors", errors.New("required field is missing from input"))...)
	}
	// Field "legend"
	if fields["legend"] != nil {
		if string(fields["legend"]) != "null" {

			resource.Legend = common.VizLegendOptions{}
			if err := resource.Legend.UnmarshalJSONStrict(fields["legend"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("legend", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("legend", errors.New("required field is null"))...)

		}
		delete(fields, "legend")
	} else {
		errs = append(errs, cog.MakeBuildErrors("legend", errors.New("required field is missing from input"))...)
	}
	// Field "includeAllFields"
	if fields["includeAllFields"] != nil {
		if string(fields["includeAllFields"]) != "null" {
			if err := json.Unmarshal(fields["includeAllFields"], &resource.IncludeAllFields); err != nil {
				errs = append(errs, cog.MakeBuildErrors("includeAllFields", err)...)
			}

		}
		delete(fields, "includeAllFields")

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
	if resource.Mode != other.Mode {
		return false
	}
	if resource.CandleStyle != other.CandleStyle {
		return false
	}
	if resource.ColorStrategy != other.ColorStrategy {
		return false
	}
	if !resource.Fields.Equals(other.Fields) {
		return false
	}
	if !resource.Colors.Equals(other.Colors) {
		return false
	}
	if !resource.Legend.Equals(other.Legend) {
		return false
	}
	if resource.IncludeAllFields == nil && other.IncludeAllFields != nil || resource.IncludeAllFields != nil && other.IncludeAllFields == nil {
		return false
	}

	if resource.IncludeAllFields != nil {
		if *resource.IncludeAllFields != *other.IncludeAllFields {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Fields.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("fields", err)...)
	}
	if err := resource.Colors.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("colors", err)...)
	}
	if err := resource.Legend.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("legend", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type FieldConfig = common.GraphFieldConfig

// VariantConfig returns the configuration related to candlestick panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "candlestick",
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
		FieldConfigUnmarshaler: func(raw []byte) (any, error) {
			fieldConfig := &FieldConfig{}

			if err := json.Unmarshal(raw, fieldConfig); err != nil {
				return nil, err
			}

			return fieldConfig, nil
		},
		StrictFieldConfigUnmarshaler: func(raw []byte) (any, error) {
			fieldConfig := &FieldConfig{}

			if err := fieldConfig.UnmarshalJSONStrict(raw); err != nil {
				return nil, err
			}

			return fieldConfig, nil
		},
		GoConverter: func(inputPanel any) string {
			if panel, ok := inputPanel.(*dashboard.Panel); ok {
				return PanelConverter(*panel)
			}

			return PanelConverter(inputPanel.(dashboard.Panel))
		},
	}
}
