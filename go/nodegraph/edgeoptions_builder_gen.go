// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package nodegraph

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[EdgeOptions] = (*EdgeOptionsBuilder)(nil)

type EdgeOptionsBuilder struct {
	internal *EdgeOptions
	errors   map[string]cog.BuildErrors
}

func NewEdgeOptionsBuilder() *EdgeOptionsBuilder {
	resource := &EdgeOptions{}
	builder := &EdgeOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *EdgeOptionsBuilder) Build() (EdgeOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return EdgeOptions{}, err
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

func (builder *EdgeOptionsBuilder) applyDefaults() {
}
