// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PrometheusDataqueryScope] = (*PrometheusDataqueryScopeBuilder)(nil)

type PrometheusDataqueryScopeBuilder struct {
	internal *PrometheusDataqueryScope
	errors   cog.BuildErrors
}

func NewPrometheusDataqueryScopeBuilder() *PrometheusDataqueryScopeBuilder {
	resource := NewPrometheusDataqueryScope()
	builder := &PrometheusDataqueryScopeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *PrometheusDataqueryScopeBuilder) Build() (PrometheusDataqueryScope, error) {
	if err := builder.internal.Validate(); err != nil {
		return PrometheusDataqueryScope{}, err
	}

	if len(builder.errors) > 0 {
		return PrometheusDataqueryScope{}, cog.MakeBuildErrors("prometheus.prometheusDataqueryScope", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PrometheusDataqueryScopeBuilder) Matchers(matchers string) *PrometheusDataqueryScopeBuilder {
	builder.internal.Matchers = matchers

	return builder
}
