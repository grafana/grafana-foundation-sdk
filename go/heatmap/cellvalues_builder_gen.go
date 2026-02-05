// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CellValues] = (*CellValuesBuilder)(nil)

// Controls cell value options
type CellValuesBuilder struct {
	internal *CellValues
	errors   cog.BuildErrors
}

func NewCellValuesBuilder() *CellValuesBuilder {
	resource := NewCellValues()
	builder := &CellValuesBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *CellValuesBuilder) Build() (CellValues, error) {
	if err := builder.internal.Validate(); err != nil {
		return CellValues{}, err
	}

	if len(builder.errors) > 0 {
		return CellValues{}, cog.MakeBuildErrors("heatmap.cellValues", builder.errors)
	}

	return *builder.internal, nil
}

// Controls the cell value unit
func (builder *CellValuesBuilder) Unit(unit string) *CellValuesBuilder {
	builder.internal.Unit = &unit

	return builder
}

// Controls the number of decimals for cell values
func (builder *CellValuesBuilder) Decimals(decimals float32) *CellValuesBuilder {
	builder.internal.Decimals = &decimals

	return builder
}
