// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bigquery

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorProperty] = (*QueryEditorPropertyBuilder)(nil)

type QueryEditorPropertyBuilder struct {
	internal *QueryEditorProperty
	errors   cog.BuildErrors
}

func NewQueryEditorPropertyBuilder() *QueryEditorPropertyBuilder {
	resource := NewQueryEditorProperty()
	builder := &QueryEditorPropertyBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *QueryEditorPropertyBuilder) Build() (QueryEditorProperty, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryEditorProperty{}, err
	}

	if len(builder.errors) > 0 {
		return QueryEditorProperty{}, cog.MakeBuildErrors("bigquery.queryEditorProperty", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryEditorPropertyBuilder) Name(name string) *QueryEditorPropertyBuilder {
	builder.internal.Name = &name

	return builder
}
