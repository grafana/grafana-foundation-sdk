// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Matcher] = (*MatcherBuilder)(nil)

type MatcherBuilder struct {
	internal *Matcher
	errors   map[string]cog.BuildErrors
}

func NewMatcherBuilder() *MatcherBuilder {
	resource := &Matcher{}
	builder := &MatcherBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *MatcherBuilder) Build() (Matcher, error) {
	if err := builder.internal.Validate(); err != nil {
		return Matcher{}, err
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

func (builder *MatcherBuilder) applyDefaults() {
}
