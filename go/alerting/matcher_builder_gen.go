// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Matcher] = (*MatcherBuilder)(nil)

type MatcherBuilder struct {
	internal *Matcher
	errors   cog.BuildErrors
}

func NewMatcherBuilder() *MatcherBuilder {
	resource := NewMatcher()
	builder := &MatcherBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MatcherBuilder) Build() (Matcher, error) {
	if err := builder.internal.Validate(); err != nil {
		return Matcher{}, err
	}

	if len(builder.errors) > 0 {
		return Matcher{}, cog.MakeBuildErrors("alerting.matcher", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MatcherBuilder) Name(name string) *MatcherBuilder {
	builder.internal.Name = &name

	return builder
}

func (builder *MatcherBuilder) Type(typeArg MatchType) *MatcherBuilder {
	builder.internal.Type = &typeArg

	return builder
}

func (builder *MatcherBuilder) Value(value string) *MatcherBuilder {
	builder.internal.Value = &value

	return builder
}
