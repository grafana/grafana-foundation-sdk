// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RuleRef] = (*RuleRefBuilder)(nil)

type RuleRefBuilder struct {
	internal *RuleRef
	errors   cog.BuildErrors
}

func NewRuleRefBuilder() *RuleRefBuilder {
	resource := NewRuleRef()
	builder := &RuleRefBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RuleRefBuilder) Build() (RuleRef, error) {
	if err := builder.internal.Validate(); err != nil {
		return RuleRef{}, err
	}

	if len(builder.errors) > 0 {
		return RuleRef{}, cog.MakeBuildErrors("alertingv1beta1.ruleRef", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RuleRefBuilder) RecordError(path string, err error) *RuleRefBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// name is the metadata.name of an AlertRule or RecordingRule resource.
func (builder *RuleRefBuilder) Name(name RuleUID) *RuleRefBuilder {
	builder.internal.Name = name

	return builder
}
