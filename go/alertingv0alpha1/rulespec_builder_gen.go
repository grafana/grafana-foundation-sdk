// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RuleSpec] = (*RuleSpecBuilder)(nil)

type RuleSpecBuilder struct {
	internal *RuleSpec
	errors   cog.BuildErrors
}

func NewRuleSpecBuilder() *RuleSpecBuilder {
	resource := NewRuleSpec()
	builder := &RuleSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RuleSpecBuilder) Build() (RuleSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return RuleSpec{}, err
	}

	if len(builder.errors) > 0 {
		return RuleSpec{}, cog.MakeBuildErrors("alertingv0alpha1.ruleSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RuleSpecBuilder) RecordError(path string, err error) *RuleSpecBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *RuleSpecBuilder) Title(title string) *RuleSpecBuilder {
	builder.internal.Title = title

	return builder
}

func (builder *RuleSpecBuilder) Paused(paused bool) *RuleSpecBuilder {
	builder.internal.Paused = &paused

	return builder
}

func (builder *RuleSpecBuilder) Trigger(trigger cog.Builder[IntervalTrigger]) *RuleSpecBuilder {
	triggerResource, err := trigger.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Trigger = triggerResource

	return builder
}

func (builder *RuleSpecBuilder) Labels(labels map[string]TemplateString) *RuleSpecBuilder {
	builder.internal.Labels = labels

	return builder
}

func (builder *RuleSpecBuilder) Expressions(expressions ExpressionMap) *RuleSpecBuilder {
	builder.internal.Expressions = expressions

	return builder
}
