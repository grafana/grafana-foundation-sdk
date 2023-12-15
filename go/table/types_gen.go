// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package table

import (
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
