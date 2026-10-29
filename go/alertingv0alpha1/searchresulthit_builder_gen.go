// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SearchResultHit] = (*SearchResultHitBuilder)(nil)

// #SearchResultHit is a single match: its identity, an optional relevance
// score (present only when the query included free text), and the requested
// fields.
type SearchResultHitBuilder struct {
	internal *SearchResultHit
	errors   cog.BuildErrors
}

func NewSearchResultHitBuilder() *SearchResultHitBuilder {
	resource := NewSearchResultHit()
	builder := &SearchResultHitBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *SearchResultHitBuilder) Build() (SearchResultHit, error) {
	if err := builder.internal.Validate(); err != nil {
		return SearchResultHit{}, err
	}

	if len(builder.errors) > 0 {
		return SearchResultHit{}, cog.MakeBuildErrors("alertingv0alpha1.searchResultHit", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SearchResultHitBuilder) RecordError(path string, err error) *SearchResultHitBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *SearchResultHitBuilder) Resource(resource cog.Builder[SearchResultResource]) *SearchResultHitBuilder {
	resourceResource, err := resource.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Resource = resourceResource

	return builder
}

func (builder *SearchResultHitBuilder) Score(score float64) *SearchResultHitBuilder {
	builder.internal.Score = &score

	return builder
}

func (builder *SearchResultHitBuilder) Fields(fields cog.Builder[RuleSearchHitFields]) *SearchResultHitBuilder {
	fieldsResource, err := fields.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Fields = fieldsResource

	return builder
}
