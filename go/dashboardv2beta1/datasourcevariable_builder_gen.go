// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DatasourceVariableKind] = (*DatasourceVariableBuilder)(nil)

// Datasource variable kind
type DatasourceVariableBuilder struct {
	internal *DatasourceVariableKind
	errors   cog.BuildErrors
}

func NewDatasourceVariableBuilder(name string) *DatasourceVariableBuilder {
	resource := NewDatasourceVariableKind()
	builder := &DatasourceVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "DatasourceVariable"
	builder.internal.Spec.Name = name

	return builder
}

func (builder *DatasourceVariableBuilder) Build() (DatasourceVariableKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return DatasourceVariableKind{}, err
	}

	if len(builder.errors) > 0 {
		return DatasourceVariableKind{}, cog.MakeBuildErrors("dashboardv2beta1.datasourceVariable", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DatasourceVariableBuilder) RecordError(path string, err error) *DatasourceVariableBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *DatasourceVariableBuilder) Name(name string) *DatasourceVariableBuilder {
	builder.internal.Spec.Name = name

	return builder
}

func (builder *DatasourceVariableBuilder) PluginId(pluginId string) *DatasourceVariableBuilder {
	builder.internal.Spec.PluginId = pluginId

	return builder
}

func (builder *DatasourceVariableBuilder) Refresh(refresh VariableRefresh) *DatasourceVariableBuilder {
	builder.internal.Spec.Refresh = refresh

	return builder
}

func (builder *DatasourceVariableBuilder) Regex(regex string) *DatasourceVariableBuilder {
	builder.internal.Spec.Regex = regex

	return builder
}

func (builder *DatasourceVariableBuilder) Current(current VariableOption) *DatasourceVariableBuilder {
	builder.internal.Spec.Current = current

	return builder
}

func (builder *DatasourceVariableBuilder) Options(options []VariableOption) *DatasourceVariableBuilder {
	builder.internal.Spec.Options = options

	return builder
}

func (builder *DatasourceVariableBuilder) Multi(multi bool) *DatasourceVariableBuilder {
	builder.internal.Spec.Multi = multi

	return builder
}

func (builder *DatasourceVariableBuilder) IncludeAll(includeAll bool) *DatasourceVariableBuilder {
	builder.internal.Spec.IncludeAll = includeAll

	return builder
}

func (builder *DatasourceVariableBuilder) AllValue(allValue string) *DatasourceVariableBuilder {
	builder.internal.Spec.AllValue = &allValue

	return builder
}

func (builder *DatasourceVariableBuilder) Label(label string) *DatasourceVariableBuilder {
	builder.internal.Spec.Label = &label

	return builder
}

func (builder *DatasourceVariableBuilder) Hide(hide VariableHide) *DatasourceVariableBuilder {
	builder.internal.Spec.Hide = hide

	return builder
}

func (builder *DatasourceVariableBuilder) SkipUrlSync(skipUrlSync bool) *DatasourceVariableBuilder {
	builder.internal.Spec.SkipUrlSync = skipUrlSync

	return builder
}

func (builder *DatasourceVariableBuilder) Description(description string) *DatasourceVariableBuilder {
	builder.internal.Spec.Description = &description

	return builder
}

func (builder *DatasourceVariableBuilder) AllowCustomValue(allowCustomValue bool) *DatasourceVariableBuilder {
	builder.internal.Spec.AllowCustomValue = allowCustomValue

	return builder
}
