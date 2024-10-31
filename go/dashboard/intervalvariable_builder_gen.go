// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VariableModel] = (*IntervalVariableBuilder)(nil)

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
type IntervalVariableBuilder struct {
	internal *VariableModel
	errors   map[string]cog.BuildErrors
}

func NewIntervalVariableBuilder(name string) *IntervalVariableBuilder {
	resource := &VariableModel{}
	builder := &IntervalVariableBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Name = name
	builder.internal.Type = "interval"

	return builder
}

func (builder *IntervalVariableBuilder) Build() (VariableModel, error) {
	if err := builder.internal.Validate(); err != nil {
		return VariableModel{}, err
	}

	return *builder.internal, nil
}

// Unique numeric identifier for the variable.
func (builder *IntervalVariableBuilder) Id(id string) *IntervalVariableBuilder {
	builder.internal.Id = id

	return builder
}

// Name of variable
func (builder *IntervalVariableBuilder) Name(name string) *IntervalVariableBuilder {
	builder.internal.Name = name

	return builder
}

// Optional display name
func (builder *IntervalVariableBuilder) Label(label string) *IntervalVariableBuilder {
	builder.internal.Label = &label

	return builder
}

// Visibility configuration for the variable
func (builder *IntervalVariableBuilder) Hide(hide VariableHide) *IntervalVariableBuilder {
	builder.internal.Hide = hide

	return builder
}

// Description of variable. It can be defined but `null`.
func (builder *IntervalVariableBuilder) Description(description string) *IntervalVariableBuilder {
	builder.internal.Description = &description

	return builder
}

// Query used to fetch values for a variable
func (builder *IntervalVariableBuilder) Values(query StringOrMap) *IntervalVariableBuilder {
	builder.internal.Query = &query

	return builder
}

// Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.
func (builder *IntervalVariableBuilder) AllFormat(allFormat string) *IntervalVariableBuilder {
	builder.internal.AllFormat = &allFormat

	return builder
}

// Shows current selected variable text/value on the dashboard
func (builder *IntervalVariableBuilder) Current(current VariableOption) *IntervalVariableBuilder {
	builder.internal.Current = &current

	return builder
}

// Options that can be selected for a variable.
func (builder *IntervalVariableBuilder) Options(options []VariableOption) *IntervalVariableBuilder {
	builder.internal.Options = options

	return builder
}

func (builder *IntervalVariableBuilder) applyDefaults() {
}
