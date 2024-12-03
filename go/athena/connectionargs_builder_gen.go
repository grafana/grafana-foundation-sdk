// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package athena

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ConnectionArgs] = (*ConnectionArgsBuilder)(nil)

type ConnectionArgsBuilder struct {
	internal *ConnectionArgs
	errors   map[string]cog.BuildErrors
}

func NewConnectionArgsBuilder() *ConnectionArgsBuilder {
	resource := NewConnectionArgs()
	builder := &ConnectionArgsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ConnectionArgsBuilder) Build() (ConnectionArgs, error) {
	if err := builder.internal.Validate(); err != nil {
		return ConnectionArgs{}, err
	}

	return *builder.internal, nil
}

func (builder *ConnectionArgsBuilder) Region(region string) *ConnectionArgsBuilder {
	builder.internal.Region = &region

	return builder
}

func (builder *ConnectionArgsBuilder) Catalog(catalog string) *ConnectionArgsBuilder {
	builder.internal.Catalog = &catalog

	return builder
}

func (builder *ConnectionArgsBuilder) Database(database string) *ConnectionArgsBuilder {
	builder.internal.Database = &database

	return builder
}

func (builder *ConnectionArgsBuilder) ResultReuseEnabled(resultReuseEnabled bool) *ConnectionArgsBuilder {
	builder.internal.ResultReuseEnabled = &resultReuseEnabled

	return builder
}

func (builder *ConnectionArgsBuilder) ResultReuseMaxAgeInMinutes(resultReuseMaxAgeInMinutes float64) *ConnectionArgsBuilder {
	builder.internal.ResultReuseMaxAgeInMinutes = &resultReuseMaxAgeInMinutes

	return builder
}