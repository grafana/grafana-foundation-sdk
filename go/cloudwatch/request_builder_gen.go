// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

var _ cog.Builder[variants.Dataquery] = (*RequestBuilder)(nil)

type RequestBuilder struct {
	internal *Request
	errors   cog.BuildErrors
}

func NewRequestBuilder() *RequestBuilder {
	resource := NewRequest()
	builder := &RequestBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RequestBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return Request{}, err
	}

	if len(builder.errors) > 0 {
		return Request{}, cog.MakeBuildErrors("cloudwatch.request", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RequestBuilder) RecordError(path string, err error) *RequestBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *RequestBuilder) MetricsQuery(metricsQuery cog.Builder[MetricsQuery]) *RequestBuilder {
	metricsQueryResource, err := metricsQuery.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.MetricsQuery = &metricsQueryResource

	return builder
}

func (builder *RequestBuilder) LogsQuery(logsQuery cog.Builder[LogsQuery]) *RequestBuilder {
	logsQueryResource, err := logsQuery.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.LogsQuery = &logsQueryResource

	return builder
}

func (builder *RequestBuilder) AnnotationQuery(annotationQuery cog.Builder[AnnotationQuery]) *RequestBuilder {
	annotationQueryResource, err := annotationQuery.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.AnnotationQuery = &annotationQueryResource

	return builder
}
