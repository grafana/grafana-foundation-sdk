// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RecordRule] = (*RecordRuleBuilder)(nil)

type RecordRuleBuilder struct {
	internal *RecordRule
	errors   cog.BuildErrors
}

func NewRecordRuleBuilder() *RecordRuleBuilder {
	resource := NewRecordRule()
	builder := &RecordRuleBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RecordRuleBuilder) Build() (RecordRule, error) {
	if err := builder.internal.Validate(); err != nil {
		return RecordRule{}, err
	}

	if len(builder.errors) > 0 {
		return RecordRule{}, cog.MakeBuildErrors("alerting.recordRule", builder.errors)
	}

	return *builder.internal, nil
}

// Which expression node should be used as the input for the recorded metric.
func (builder *RecordRuleBuilder) From(from string) *RecordRuleBuilder {
	builder.internal.From = from

	return builder
}

// Name of the recorded metric.
func (builder *RecordRuleBuilder) Metric(metric string) *RecordRuleBuilder {
	builder.internal.Metric = metric

	return builder
}
