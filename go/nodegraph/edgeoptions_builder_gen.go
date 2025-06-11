// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package nodegraph

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[EdgeOptions] = (*EdgeOptionsBuilder)(nil)

type EdgeOptionsBuilder struct {
	internal *EdgeOptions
	errors   cog.BuildErrors
}

func NewEdgeOptionsBuilder() *EdgeOptionsBuilder {
	resource := NewEdgeOptions()
	builder := &EdgeOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *EdgeOptionsBuilder) Build() (EdgeOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return EdgeOptions{}, err
	}

	if len(builder.errors) > 0 {
		return EdgeOptions{}, cog.MakeBuildErrors("nodegraph.edgeOptions", builder.errors)
	}

	return *builder.internal, nil
}

// Unit for the main stat to override what ever is set in the data frame.
func (builder *EdgeOptionsBuilder) MainStatUnit(mainStatUnit string) *EdgeOptionsBuilder {
	builder.internal.MainStatUnit = &mainStatUnit

	return builder
}

// Unit for the secondary stat to override what ever is set in the data frame.
func (builder *EdgeOptionsBuilder) SecondaryStatUnit(secondaryStatUnit string) *EdgeOptionsBuilder {
	builder.internal.SecondaryStatUnit = &secondaryStatUnit

	return builder
}
