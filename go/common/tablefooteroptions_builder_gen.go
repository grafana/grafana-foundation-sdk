// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TableFooterOptions] = (*TableFooterOptionsBuilder)(nil)

// Footer options
type TableFooterOptionsBuilder struct {
	internal *TableFooterOptions
	errors   map[string]cog.BuildErrors
}

func NewTableFooterOptionsBuilder() *TableFooterOptionsBuilder {
	resource := NewTableFooterOptions()
	builder := &TableFooterOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *TableFooterOptionsBuilder) Build() (TableFooterOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return TableFooterOptions{}, err
	}

	return *builder.internal, nil
}

func (builder *TableFooterOptionsBuilder) Show(show bool) *TableFooterOptionsBuilder {
	builder.internal.Show = show

	return builder
}

// actually 1 value
func (builder *TableFooterOptionsBuilder) Reducer(reducer []string) *TableFooterOptionsBuilder {
	builder.internal.Reducer = reducer

	return builder
}

func (builder *TableFooterOptionsBuilder) Fields(fields []string) *TableFooterOptionsBuilder {
	builder.internal.Fields = fields

	return builder
}

func (builder *TableFooterOptionsBuilder) EnablePagination(enablePagination bool) *TableFooterOptionsBuilder {
	builder.internal.EnablePagination = &enablePagination

	return builder
}

func (builder *TableFooterOptionsBuilder) CountRows(countRows bool) *TableFooterOptionsBuilder {
	builder.internal.CountRows = &countRows

	return builder
}
