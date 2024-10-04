// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CellValues] = (*CellValuesBuilder)(nil)

// Controls cell value options
type CellValuesBuilder struct {
	internal *CellValues
	errors   map[string]cog.BuildErrors
}

func NewCellValuesBuilder() *CellValuesBuilder {
	resource := &CellValues{}
	builder := &CellValuesBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *CellValuesBuilder) Build() (CellValues, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("CellValues", err)...)
	}

	if len(errs) != 0 {
		return CellValues{}, errs
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

func (builder *CellValuesBuilder) applyDefaults() {
}
