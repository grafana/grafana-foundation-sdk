// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VariableModel] = (*VariableModelBuilder)(nil)

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
type VariableModelBuilder struct {
	internal *VariableModel
	errors   map[string]cog.BuildErrors
}

func NewVariableModelBuilder() *VariableModelBuilder {
	resource := &VariableModel{}
	builder := &VariableModelBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *VariableModelBuilder) Build() (VariableModel, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("VariableModel", err)...)
	}

	if len(errs) != 0 {
		return VariableModel{}, errs
	}

	return *builder.internal, nil
}

// Type of variable
func (builder *VariableModelBuilder) Type(typeArg VariableType) *VariableModelBuilder {
	builder.internal.Type = typeArg

	return builder
}

// Name of variable
func (builder *VariableModelBuilder) Name(name string) *VariableModelBuilder {
	builder.internal.Name = name

	return builder
}

// Optional display name
func (builder *VariableModelBuilder) Label(label string) *VariableModelBuilder {
	builder.internal.Label = &label

	return builder
}

// Visibility configuration for the variable
func (builder *VariableModelBuilder) Hide(hide VariableHide) *VariableModelBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Whether the variable value should be managed by URL query params or not
func (builder *VariableModelBuilder) SkipUrlSync(skipUrlSync bool) *VariableModelBuilder {
	builder.internal.SkipUrlSync = &skipUrlSync

	return builder
}

// Description of variable. It can be defined but `null`.
func (builder *VariableModelBuilder) Description(description string) *VariableModelBuilder {
	builder.internal.Description = &description

	return builder
}

// Query used to fetch values for a variable
func (builder *VariableModelBuilder) Query(query StringOrAny) *VariableModelBuilder {
	builder.internal.Query = &query

	return builder
}

// Data source used to fetch values for a variable. It can be defined but `null`.
func (builder *VariableModelBuilder) Datasource(datasource DataSourceRef) *VariableModelBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Shows current selected variable text/value on the dashboard
func (builder *VariableModelBuilder) Current(current VariableOption) *VariableModelBuilder {
	builder.internal.Current = &current

	return builder
}

// Whether multiple values can be selected or not from variable value list
func (builder *VariableModelBuilder) Multi(multi bool) *VariableModelBuilder {
	builder.internal.Multi = &multi

	return builder
}

// Options that can be selected for a variable.
func (builder *VariableModelBuilder) Options(options []VariableOption) *VariableModelBuilder {
	builder.internal.Options = options

	return builder
}

func (builder *VariableModelBuilder) Refresh(refresh VariableRefresh) *VariableModelBuilder {
	builder.internal.Refresh = &refresh

	return builder
}

// Options sort order
func (builder *VariableModelBuilder) Sort(sort VariableSort) *VariableModelBuilder {
	builder.internal.Sort = &sort

	return builder
}

func (builder *VariableModelBuilder) applyDefaults() {
	builder.SkipUrlSync(false)
	builder.Multi(false)
}
