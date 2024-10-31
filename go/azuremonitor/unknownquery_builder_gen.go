// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[UnknownQuery] = (*UnknownQueryBuilder)(nil)

type UnknownQueryBuilder struct {
	internal *UnknownQuery
	errors   map[string]cog.BuildErrors
}

func NewUnknownQueryBuilder() *UnknownQueryBuilder {
	resource := &UnknownQuery{}
	builder := &UnknownQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Kind = "UnknownQuery"

	return builder
}

func (builder *UnknownQueryBuilder) Build() (UnknownQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return UnknownQuery{}, err
	}

	return *builder.internal, nil
}

func (builder *UnknownQueryBuilder) RawQuery(rawQuery string) *UnknownQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}

func (builder *UnknownQueryBuilder) applyDefaults() {
}
