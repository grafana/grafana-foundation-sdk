// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AutoGridRepeatOptions] = (*AutoGridRepeatOptionsBuilder)(nil)

type AutoGridRepeatOptionsBuilder struct {
	internal *AutoGridRepeatOptions
	errors   cog.BuildErrors
}

func NewAutoGridRepeatOptionsBuilder() *AutoGridRepeatOptionsBuilder {
	resource := NewAutoGridRepeatOptions()
	builder := &AutoGridRepeatOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AutoGridRepeatOptionsBuilder) Build() (AutoGridRepeatOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return AutoGridRepeatOptions{}, err
	}

	if len(builder.errors) > 0 {
		return AutoGridRepeatOptions{}, cog.MakeBuildErrors("dashboardv2beta1.autoGridRepeatOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AutoGridRepeatOptionsBuilder) RecordError(path string, err error) *AutoGridRepeatOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *AutoGridRepeatOptionsBuilder) Value(value string) *AutoGridRepeatOptionsBuilder {
	builder.internal.Value = value

	return builder
}
