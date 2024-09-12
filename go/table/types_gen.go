// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package table

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	// Represents the index of the selected frame
	FrameIndex float64 `json:"frameIndex"`
	// Controls whether the panel should show the header
	ShowHeader bool `json:"showHeader"`
	// Controls whether the header should show icons for the column types
	ShowTypeIcons *bool `json:"showTypeIcons,omitempty"`
	// Used to control row sorting
	SortBy []common.TableSortByFieldState `json:"sortBy,omitempty"`
	// Controls footer options
	Footer *common.TableFooterOptions `json:"footer,omitempty"`
	// Controls the height of the rows
	CellHeight *common.TableCellHeight `json:"cellHeight,omitempty"`
}

type FieldConfig = common.TableFieldOptions

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "table",
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
