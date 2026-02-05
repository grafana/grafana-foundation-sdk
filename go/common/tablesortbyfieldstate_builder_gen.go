// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TableSortByFieldState] = (*TableSortByFieldStateBuilder)(nil)

// Sort by field state
type TableSortByFieldStateBuilder struct {
	internal *TableSortByFieldState
	errors   cog.BuildErrors
}

func NewTableSortByFieldStateBuilder() *TableSortByFieldStateBuilder {
	resource := NewTableSortByFieldState()
	builder := &TableSortByFieldStateBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TableSortByFieldStateBuilder) Build() (TableSortByFieldState, error) {
	if err := builder.internal.Validate(); err != nil {
		return TableSortByFieldState{}, err
	}

	if len(builder.errors) > 0 {
		return TableSortByFieldState{}, cog.MakeBuildErrors("common.tableSortByFieldState", builder.errors)
	}

	return *builder.internal, nil
}

// Sets the display name of the field to sort by
func (builder *TableSortByFieldStateBuilder) DisplayName(displayName string) *TableSortByFieldStateBuilder {
	builder.internal.DisplayName = displayName

	return builder
}

// Flag used to indicate descending sort order
func (builder *TableSortByFieldStateBuilder) Desc(desc bool) *TableSortByFieldStateBuilder {
	builder.internal.Desc = &desc

	return builder
}
