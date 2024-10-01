// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package candlestick

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
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

type CandlestickColors struct {
	Up   string `json:"up"`
	Down string `json:"down"`
	Flat string `json:"flat"`
}

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

type FieldConfig = common.GraphFieldConfig

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
		FieldConfigUnmarshaler: func(raw []byte) (any, error) {
			fieldConfig := &FieldConfig{}

			if err := json.Unmarshal(raw, fieldConfig); err != nil {
				return nil, err
			}

			return fieldConfig, nil
		},
	}
}
