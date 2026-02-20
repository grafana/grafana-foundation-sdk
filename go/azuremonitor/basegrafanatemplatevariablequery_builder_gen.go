// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BaseGrafanaTemplateVariableQuery] = (*BaseGrafanaTemplateVariableQueryBuilder)(nil)

type BaseGrafanaTemplateVariableQueryBuilder struct {
	internal *BaseGrafanaTemplateVariableQuery
	errors   cog.BuildErrors
}

func NewBaseGrafanaTemplateVariableQueryBuilder() *BaseGrafanaTemplateVariableQueryBuilder {
	resource := NewBaseGrafanaTemplateVariableQuery()
	builder := &BaseGrafanaTemplateVariableQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BaseGrafanaTemplateVariableQueryBuilder) Build() (BaseGrafanaTemplateVariableQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return BaseGrafanaTemplateVariableQuery{}, err
	}

	if len(builder.errors) > 0 {
		return BaseGrafanaTemplateVariableQuery{}, cog.MakeBuildErrors("azuremonitor.baseGrafanaTemplateVariableQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BaseGrafanaTemplateVariableQueryBuilder) RecordError(path string, err error) *BaseGrafanaTemplateVariableQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *BaseGrafanaTemplateVariableQueryBuilder) RawQuery(rawQuery string) *BaseGrafanaTemplateVariableQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}
