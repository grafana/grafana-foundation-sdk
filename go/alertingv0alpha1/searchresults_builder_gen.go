// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SearchResults] = (*SearchResultsBuilder)(nil)

// #SearchResults is the response envelope, mirroring
// search.grafana.app SearchResults.
type SearchResultsBuilder struct {
	internal *SearchResults
	errors   cog.BuildErrors
}

func NewSearchResultsBuilder() *SearchResultsBuilder {
	resource := NewSearchResults()
	builder := &SearchResultsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *SearchResultsBuilder) Build() (SearchResults, error) {
	if err := builder.internal.Validate(); err != nil {
		return SearchResults{}, err
	}

	if len(builder.errors) > 0 {
		return SearchResults{}, cog.MakeBuildErrors("alertingv0alpha1.searchResults", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SearchResultsBuilder) RecordError(path string, err error) *SearchResultsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *SearchResultsBuilder) Metadata(metadata cog.Builder[SearchResultsMetadata]) *SearchResultsBuilder {
	metadataResource, err := metadata.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Metadata = metadataResource

	return builder
}

func (builder *SearchResultsBuilder) Items(items []cog.Builder[SearchResultHit]) *SearchResultsBuilder {
	itemsResources := make([]SearchResultHit, 0, len(items))
	for _, r1 := range items {
		itemsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		itemsResources = append(itemsResources, itemsDepth1)
	}
	builder.internal.Items = itemsResources

	return builder
}

func (builder *SearchResultsBuilder) Facets(facets map[string][]cog.Builder[FacetValue]) *SearchResultsBuilder {
	facetsResource := make(map[string][]FacetValue, len(facets))
	for key1, val1 := range facets {
		facetsDepth1 := make([]FacetValue, 0, len(val1))
		for _, r2 := range val1 {
			facetsDepth2, err := r2.Build()
			if err != nil {
				builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
				return builder
			}
			facetsDepth1 = append(facetsDepth1, facetsDepth2)
		}

		facetsResource[key1] = facetsDepth1
	}
	builder.internal.Facets = facetsResource

	return builder
}
