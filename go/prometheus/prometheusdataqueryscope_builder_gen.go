// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PrometheusDataqueryScope] = (*PrometheusDataqueryScopeBuilder)(nil)

type PrometheusDataqueryScopeBuilder struct {
	internal *PrometheusDataqueryScope
	errors   map[string]cog.BuildErrors
}

func NewPrometheusDataqueryScopeBuilder() *PrometheusDataqueryScopeBuilder {
	resource := &PrometheusDataqueryScope{}
	builder := &PrometheusDataqueryScopeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *PrometheusDataqueryScopeBuilder) Build() (PrometheusDataqueryScope, error) {
	if err := builder.internal.Validate(); err != nil {
		return PrometheusDataqueryScope{}, err
	}

	return *builder.internal, nil
}

func (builder *PrometheusDataqueryScopeBuilder) Matchers(matchers string) *PrometheusDataqueryScopeBuilder {
	builder.internal.Matchers = matchers

	return builder
}

func (builder *PrometheusDataqueryScopeBuilder) applyDefaults() {
}
