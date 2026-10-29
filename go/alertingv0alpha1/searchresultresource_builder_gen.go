// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SearchResultResource] = (*SearchResultResourceBuilder)(nil)

// #SearchResultResource is the full identity of a hit.
type SearchResultResourceBuilder struct {
	internal *SearchResultResource
	errors   cog.BuildErrors
}

func NewSearchResultResourceBuilder() *SearchResultResourceBuilder {
	resource := NewSearchResultResource()
	builder := &SearchResultResourceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *SearchResultResourceBuilder) Build() (SearchResultResource, error) {
	if err := builder.internal.Validate(); err != nil {
		return SearchResultResource{}, err
	}

	if len(builder.errors) > 0 {
		return SearchResultResource{}, cog.MakeBuildErrors("alertingv0alpha1.searchResultResource", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SearchResultResourceBuilder) RecordError(path string, err error) *SearchResultResourceBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *SearchResultResourceBuilder) Group(group string) *SearchResultResourceBuilder {
	builder.internal.Group = group

	return builder
}

func (builder *SearchResultResourceBuilder) Resource(resource string) *SearchResultResourceBuilder {
	builder.internal.Resource = resource

	return builder
}

func (builder *SearchResultResourceBuilder) Kind(kind string) *SearchResultResourceBuilder {
	builder.internal.Kind = kind

	return builder
}

func (builder *SearchResultResourceBuilder) Name(name string) *SearchResultResourceBuilder {
	builder.internal.Name = name

	return builder
}
