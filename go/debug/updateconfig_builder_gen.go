// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package debug

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[UpdateConfig] = (*UpdateConfigBuilder)(nil)

type UpdateConfigBuilder struct {
	internal *UpdateConfig
	errors   cog.BuildErrors
}

func NewUpdateConfigBuilder() *UpdateConfigBuilder {
	resource := NewUpdateConfig()
	builder := &UpdateConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *UpdateConfigBuilder) Build() (UpdateConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return UpdateConfig{}, err
	}

	if len(builder.errors) > 0 {
		return UpdateConfig{}, cog.MakeBuildErrors("debug.updateConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *UpdateConfigBuilder) RecordError(path string, err error) *UpdateConfigBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *UpdateConfigBuilder) Render(render bool) *UpdateConfigBuilder {
	builder.internal.Render = render

	return builder
}

func (builder *UpdateConfigBuilder) DataChanged(dataChanged bool) *UpdateConfigBuilder {
	builder.internal.DataChanged = dataChanged

	return builder
}

func (builder *UpdateConfigBuilder) SchemaChanged(schemaChanged bool) *UpdateConfigBuilder {
	builder.internal.SchemaChanged = schemaChanged

	return builder
}
