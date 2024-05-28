// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package candlestick

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
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

type CandlestickColors struct {
	Up   string `json:"up"`
	Down string `json:"down"`
	Flat string `json:"flat"`
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
	Colors  CandlestickColors        `json:"colors"`
	Legend  common.VizLegendOptions  `json:"legend"`
	Tooltip common.VizTooltipOptions `json:"tooltip"`
	// When enabled, all fields will be sent to the graph
	IncludeAllFields *bool `json:"includeAllFields,omitempty"`
}

type FieldConfig = common.GraphFieldConfig

func VariantConfig() cogvariants.PanelcfgConfig {
	return cogvariants.PanelcfgConfig{
		Identifier: "candlestick",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
		FieldConfigUnmarshaler: func(raw []byte) (any, error) {
			fieldConfig := FieldConfig{}

			if err := json.Unmarshal(raw, &fieldConfig); err != nil {
				return nil, err
			}

			return fieldConfig, nil
		},
	}
}
