// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[FetchOptions] = (*FetchOptionsBuilder)(nil)

type FetchOptionsBuilder struct {
	internal *FetchOptions
	errors   cog.BuildErrors
}

func NewFetchOptionsBuilder() *FetchOptionsBuilder {
	resource := NewFetchOptions()
	builder := &FetchOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *FetchOptionsBuilder) Build() (FetchOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return FetchOptions{}, err
	}

	if len(builder.errors) > 0 {
		return FetchOptions{}, cog.MakeBuildErrors("dashboardv2beta1.fetchOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *FetchOptionsBuilder) RecordError(path string, err error) *FetchOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *FetchOptionsBuilder) Method(method HttpRequestMethod) *FetchOptionsBuilder {
	builder.internal.Method = method

	return builder
}

func (builder *FetchOptionsBuilder) Url(url string) *FetchOptionsBuilder {
	builder.internal.Url = url

	return builder
}

func (builder *FetchOptionsBuilder) Body(body string) *FetchOptionsBuilder {
	builder.internal.Body = &body

	return builder
}

// These are 2D arrays of strings, each representing a key-value pair
// We are defining them this way because we can't generate a go struct that
// that would have exactly two strings in each sub-array
func (builder *FetchOptionsBuilder) QueryParams(queryParams [][]string) *FetchOptionsBuilder {
	builder.internal.QueryParams = queryParams

	return builder
}

func (builder *FetchOptionsBuilder) Headers(headers [][]string) *FetchOptionsBuilder {
	builder.internal.Headers = headers

	return builder
}
