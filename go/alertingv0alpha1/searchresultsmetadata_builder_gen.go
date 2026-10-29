// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SearchResultsMetadata] = (*SearchResultsMetadataBuilder)(nil)

// #SearchResultsMetadata carries the paging token and total authorised match
// count.
type SearchResultsMetadataBuilder struct {
	internal *SearchResultsMetadata
	errors   cog.BuildErrors
}

func NewSearchResultsMetadataBuilder() *SearchResultsMetadataBuilder {
	resource := NewSearchResultsMetadata()
	builder := &SearchResultsMetadataBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *SearchResultsMetadataBuilder) Build() (SearchResultsMetadata, error) {
	if err := builder.internal.Validate(); err != nil {
		return SearchResultsMetadata{}, err
	}

	if len(builder.errors) > 0 {
		return SearchResultsMetadata{}, cog.MakeBuildErrors("alertingv0alpha1.searchResultsMetadata", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SearchResultsMetadataBuilder) RecordError(path string, err error) *SearchResultsMetadataBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *SearchResultsMetadataBuilder) Continue(continueArg string) *SearchResultsMetadataBuilder {
	builder.internal.Continue = &continueArg

	return builder
}

func (builder *SearchResultsMetadataBuilder) TotalHits(totalHits int64) *SearchResultsMetadataBuilder {
	builder.internal.TotalHits = &totalHits

	return builder
}

// Always read totalHits together with totalHitsRelation.
func (builder *SearchResultsMetadataBuilder) TotalHitsRelation(totalHitsRelation TotalHitsRelation) *SearchResultsMetadataBuilder {
	builder.internal.TotalHitsRelation = &totalHitsRelation

	return builder
}
