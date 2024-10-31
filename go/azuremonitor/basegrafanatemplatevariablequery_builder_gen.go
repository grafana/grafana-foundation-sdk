// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BaseGrafanaTemplateVariableQuery] = (*BaseGrafanaTemplateVariableQueryBuilder)(nil)

type BaseGrafanaTemplateVariableQueryBuilder struct {
	internal *BaseGrafanaTemplateVariableQuery
	errors   map[string]cog.BuildErrors
}

func NewBaseGrafanaTemplateVariableQueryBuilder() *BaseGrafanaTemplateVariableQueryBuilder {
	resource := &BaseGrafanaTemplateVariableQuery{}
	builder := &BaseGrafanaTemplateVariableQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *BaseGrafanaTemplateVariableQueryBuilder) Build() (BaseGrafanaTemplateVariableQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return BaseGrafanaTemplateVariableQuery{}, err
	}

	return *builder.internal, nil
}

func (builder *BaseGrafanaTemplateVariableQueryBuilder) RawQuery(rawQuery string) *BaseGrafanaTemplateVariableQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}

func (builder *BaseGrafanaTemplateVariableQueryBuilder) applyDefaults() {
}
