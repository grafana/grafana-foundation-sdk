// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[InfinityOptions] = (*InfinityOptionsBuilder)(nil)

type InfinityOptionsBuilder struct {
	internal *InfinityOptions
	errors   cog.BuildErrors
}

func NewInfinityOptionsBuilder() *InfinityOptionsBuilder {
	resource := NewInfinityOptions()
	builder := &InfinityOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *InfinityOptionsBuilder) Build() (InfinityOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return InfinityOptions{}, err
	}

	if len(builder.errors) > 0 {
		return InfinityOptions{}, cog.MakeBuildErrors("dashboardv2beta1.infinityOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *InfinityOptionsBuilder) Method(method HttpRequestMethod) *InfinityOptionsBuilder {
	builder.internal.Method = method

	return builder
}

func (builder *InfinityOptionsBuilder) Url(url string) *InfinityOptionsBuilder {
	builder.internal.Url = url

	return builder
}

func (builder *InfinityOptionsBuilder) Body(body string) *InfinityOptionsBuilder {
	builder.internal.Body = &body

	return builder
}

// These are 2D arrays of strings, each representing a key-value pair
// We are defining them this way because we can't generate a go struct that
// that would have exactly two strings in each sub-array
func (builder *InfinityOptionsBuilder) QueryParams(queryParams [][]string) *InfinityOptionsBuilder {
	builder.internal.QueryParams = queryParams

	return builder
}

func (builder *InfinityOptionsBuilder) DatasourceUid(datasourceUid string) *InfinityOptionsBuilder {
	builder.internal.DatasourceUid = datasourceUid

	return builder
}

func (builder *InfinityOptionsBuilder) Headers(headers [][]string) *InfinityOptionsBuilder {
	builder.internal.Headers = headers

	return builder
}
