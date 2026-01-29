// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AdhocVariableKind] = (*AdhocVariableBuilder)(nil)

// Adhoc variable kind
type AdhocVariableBuilder struct {
	internal *AdhocVariableKind
	errors   cog.BuildErrors
}

func NewAdhocVariableBuilder(name string) *AdhocVariableBuilder {
	resource := NewAdhocVariableKind()
	builder := &AdhocVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "AdhocVariable"
	builder.internal.Spec.Name = name

	return builder
}

func (builder *AdhocVariableBuilder) Build() (AdhocVariableKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return AdhocVariableKind{}, err
	}

	if len(builder.errors) > 0 {
		return AdhocVariableKind{}, cog.MakeBuildErrors("dashboardv2beta1.adhocVariable", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AdhocVariableBuilder) Group(group string) *AdhocVariableBuilder {
	builder.internal.Group = group

	return builder
}

func (builder *AdhocVariableBuilder) Datasource(datasource cog.Builder[Dashboardv2beta1AdhocVariableKindDatasource]) *AdhocVariableBuilder {
	datasourceResource, err := datasource.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Datasource = &datasourceResource

	return builder
}

func (builder *AdhocVariableBuilder) Spec(spec AdhocVariableSpec) *AdhocVariableBuilder {
	builder.internal.Spec = spec

	return builder
}

func (builder *AdhocVariableBuilder) Name(name string) *AdhocVariableBuilder {
	builder.internal.Spec.Name = name

	return builder
}

func (builder *AdhocVariableBuilder) BaseFilters(baseFilters []cog.Builder[AdHocFilterWithLabels]) *AdhocVariableBuilder {
	baseFiltersResources := make([]AdHocFilterWithLabels, 0, len(baseFilters))
	for _, r1 := range baseFilters {
		baseFiltersDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		baseFiltersResources = append(baseFiltersResources, baseFiltersDepth1)
	}
	builder.internal.Spec.BaseFilters = baseFiltersResources

	return builder
}

func (builder *AdhocVariableBuilder) Filters(filters []cog.Builder[AdHocFilterWithLabels]) *AdhocVariableBuilder {
	filtersResources := make([]AdHocFilterWithLabels, 0, len(filters))
	for _, r1 := range filters {
		filtersDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		filtersResources = append(filtersResources, filtersDepth1)
	}
	builder.internal.Spec.Filters = filtersResources

	return builder
}

func (builder *AdhocVariableBuilder) DefaultKeys(defaultKeys []cog.Builder[MetricFindValue]) *AdhocVariableBuilder {
	defaultKeysResources := make([]MetricFindValue, 0, len(defaultKeys))
	for _, r1 := range defaultKeys {
		defaultKeysDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		defaultKeysResources = append(defaultKeysResources, defaultKeysDepth1)
	}
	builder.internal.Spec.DefaultKeys = defaultKeysResources

	return builder
}

func (builder *AdhocVariableBuilder) Label(label string) *AdhocVariableBuilder {
	builder.internal.Spec.Label = &label

	return builder
}

func (builder *AdhocVariableBuilder) Hide(hide VariableHide) *AdhocVariableBuilder {
	builder.internal.Spec.Hide = hide

	return builder
}

func (builder *AdhocVariableBuilder) SkipUrlSync(skipUrlSync bool) *AdhocVariableBuilder {
	builder.internal.Spec.SkipUrlSync = skipUrlSync

	return builder
}

func (builder *AdhocVariableBuilder) Description(description string) *AdhocVariableBuilder {
	builder.internal.Spec.Description = &description

	return builder
}

func (builder *AdhocVariableBuilder) AllowCustomValue(allowCustomValue bool) *AdhocVariableBuilder {
	builder.internal.Spec.AllowCustomValue = allowCustomValue

	return builder
}
