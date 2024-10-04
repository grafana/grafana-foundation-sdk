// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package debug

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[UpdateConfig] = (*UpdateConfigBuilder)(nil)

type UpdateConfigBuilder struct {
	internal *UpdateConfig
	errors   map[string]cog.BuildErrors
}

func NewUpdateConfigBuilder() *UpdateConfigBuilder {
	resource := &UpdateConfig{}
	builder := &UpdateConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *UpdateConfigBuilder) Build() (UpdateConfig, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("UpdateConfig", err)...)
	}

	if len(errs) != 0 {
		return UpdateConfig{}, errs
	}

	return *builder.internal, nil
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

func (builder *UpdateConfigBuilder) applyDefaults() {
}
