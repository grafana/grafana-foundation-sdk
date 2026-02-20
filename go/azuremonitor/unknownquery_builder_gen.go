// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[UnknownQuery] = (*UnknownQueryBuilder)(nil)

type UnknownQueryBuilder struct {
	internal *UnknownQuery
	errors   cog.BuildErrors
}

func NewUnknownQueryBuilder() *UnknownQueryBuilder {
	resource := NewUnknownQuery()
	builder := &UnknownQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "UnknownQuery"

	return builder
}

func (builder *UnknownQueryBuilder) Build() (UnknownQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return UnknownQuery{}, err
	}

	if len(builder.errors) > 0 {
		return UnknownQuery{}, cog.MakeBuildErrors("azuremonitor.unknownQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *UnknownQueryBuilder) RecordError(path string, err error) *UnknownQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *UnknownQueryBuilder) RawQuery(rawQuery string) *UnknownQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}
