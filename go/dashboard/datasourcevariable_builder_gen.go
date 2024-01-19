// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VariableModel] = (*DatasourceVariableBuilder)(nil)

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
type DatasourceVariableBuilder struct {
	internal *VariableModel
	errors   map[string]cog.BuildErrors
}

func NewDatasourceVariableBuilder(name string) *DatasourceVariableBuilder {
	resource := &VariableModel{}
	builder := &DatasourceVariableBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Name = name
	builder.internal.Type = "datasource"

	return builder
}

func (builder *DatasourceVariableBuilder) Build() (VariableModel, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("DatasourceVariable", err)...)
	}

	if len(errs) != 0 {
		return VariableModel{}, errs
	}

	return *builder.internal, nil
}

// Unique numeric identifier for the variable.
func (builder *DatasourceVariableBuilder) Id(id string) *DatasourceVariableBuilder {
	builder.internal.Id = id

	return builder
}

// Name of variable
func (builder *DatasourceVariableBuilder) Name(name string) *DatasourceVariableBuilder {
	builder.internal.Name = name

	return builder
}

// Optional display name
func (builder *DatasourceVariableBuilder) Label(label string) *DatasourceVariableBuilder {
	builder.internal.Label = &label

	return builder
}

// Visibility configuration for the variable
func (builder *DatasourceVariableBuilder) Hide(hide VariableHide) *DatasourceVariableBuilder {
	builder.internal.Hide = hide

	return builder
}

// Description of variable. It can be defined but `null`.
func (builder *DatasourceVariableBuilder) Description(description string) *DatasourceVariableBuilder {
	builder.internal.Description = &description

	return builder
}

// Query used to fetch values for a variable
func (builder *DatasourceVariableBuilder) Type(string string) *DatasourceVariableBuilder {
	if builder.internal.Query == nil {
		builder.internal.Query = &StringOrAny{}
	}
	builder.internal.Query.String = &string

	return builder
}

// Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.
func (builder *DatasourceVariableBuilder) AllFormat(allFormat string) *DatasourceVariableBuilder {
	builder.internal.AllFormat = &allFormat

	return builder
}

// Shows current selected variable text/value on the dashboard
func (builder *DatasourceVariableBuilder) Current(current VariableOption) *DatasourceVariableBuilder {
	builder.internal.Current = &current

	return builder
}

// Whether multiple values can be selected or not from variable value list
func (builder *DatasourceVariableBuilder) Multi(multi bool) *DatasourceVariableBuilder {
	builder.internal.Multi = &multi

	return builder
}

func (builder *DatasourceVariableBuilder) AllValue(allValue string) *DatasourceVariableBuilder {
	builder.internal.AllValue = &allValue

	return builder
}

func (builder *DatasourceVariableBuilder) Regex(regex string) *DatasourceVariableBuilder {
	builder.internal.Regex = &regex

	return builder
}

func (builder *DatasourceVariableBuilder) IncludeAll(includeAll bool) *DatasourceVariableBuilder {
	builder.internal.IncludeAll = &includeAll

	return builder
}

func (builder *DatasourceVariableBuilder) applyDefaults() {
}
