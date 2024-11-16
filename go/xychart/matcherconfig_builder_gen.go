// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MatcherConfig] = (*MatcherConfigBuilder)(nil)

// NOTE: (copied from dashboard_kind.cue, since not exported)
// Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
// It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
type MatcherConfigBuilder struct {
	internal *MatcherConfig
	errors   map[string]cog.BuildErrors
}

func NewMatcherConfigBuilder() *MatcherConfigBuilder {
	resource := NewMatcherConfig()
	builder := &MatcherConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *MatcherConfigBuilder) Build() (MatcherConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return MatcherConfig{}, err
	}

	return *builder.internal, nil
}

// The matcher id. This is used to find the matcher implementation from registry.
func (builder *MatcherConfigBuilder) Id(id string) *MatcherConfigBuilder {
	builder.internal.Id = id

	return builder
}

// The matcher options. This is specific to the matcher implementation.
func (builder *MatcherConfigBuilder) Options(options any) *MatcherConfigBuilder {
	builder.internal.Options = &options

	return builder
}
