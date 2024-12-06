// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TableAutoCellOptions] = (*TableAutoCellOptionsBuilder)(nil)

// Auto mode table cell options
type TableAutoCellOptionsBuilder struct {
	internal *TableAutoCellOptions
	errors   map[string]cog.BuildErrors
}

func NewTableAutoCellOptionsBuilder() *TableAutoCellOptionsBuilder {
	resource := NewTableAutoCellOptions()
	builder := &TableAutoCellOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}
	builder.internal.Type = "auto"

	return builder
}

func (builder *TableAutoCellOptionsBuilder) Build() (TableAutoCellOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return TableAutoCellOptions{}, err
	}

	return *builder.internal, nil
}

func (builder *TableAutoCellOptionsBuilder) WrapText(wrapText bool) *TableAutoCellOptionsBuilder {
	builder.internal.WrapText = &wrapText

	return builder
}
