// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VariableModel] = (*QueryVariableBuilder)(nil)

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
type QueryVariableBuilder struct {
	internal *VariableModel
	errors   map[string]cog.BuildErrors
}

func NewQueryVariableBuilder(name string) *QueryVariableBuilder {
	resource := NewVariableModel()
	builder := &QueryVariableBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}
	builder.internal.Name = name
	builder.internal.Type = "query"

	return builder
}

func (builder *QueryVariableBuilder) Build() (VariableModel, error) {
	if err := builder.internal.Validate(); err != nil {
		return VariableModel{}, err
	}

	return *builder.internal, nil
}

// Name of variable
func (builder *QueryVariableBuilder) Name(name string) *QueryVariableBuilder {
	builder.internal.Name = name

	return builder
}

// Optional display name
func (builder *QueryVariableBuilder) Label(label string) *QueryVariableBuilder {
	builder.internal.Label = &label

	return builder
}

// Visibility configuration for the variable
func (builder *QueryVariableBuilder) Hide(hide VariableHide) *QueryVariableBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Description of variable. It can be defined but `null`.
func (builder *QueryVariableBuilder) Description(description string) *QueryVariableBuilder {
	builder.internal.Description = &description

	return builder
}

// Query used to fetch values for a variable
func (builder *QueryVariableBuilder) Query(query StringOrMap) *QueryVariableBuilder {
	builder.internal.Query = &query

	return builder
}

// Data source used to fetch values for a variable. It can be defined but `null`.
func (builder *QueryVariableBuilder) Datasource(datasource DataSourceRef) *QueryVariableBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Shows current selected variable text/value on the dashboard
func (builder *QueryVariableBuilder) Current(current VariableOption) *QueryVariableBuilder {
	builder.internal.Current = &current

	return builder
}

// Whether multiple values can be selected or not from variable value list
func (builder *QueryVariableBuilder) Multi(multi bool) *QueryVariableBuilder {
	builder.internal.Multi = &multi

	return builder
}

// Options that can be selected for a variable.
func (builder *QueryVariableBuilder) Options(options []VariableOption) *QueryVariableBuilder {
	builder.internal.Options = options

	return builder
}

func (builder *QueryVariableBuilder) Refresh(refresh VariableRefresh) *QueryVariableBuilder {
	builder.internal.Refresh = &refresh

	return builder
}

// Options sort order
func (builder *QueryVariableBuilder) Sort(sort VariableSort) *QueryVariableBuilder {
	builder.internal.Sort = &sort

	return builder
}
