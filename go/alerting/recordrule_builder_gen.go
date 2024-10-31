// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RecordRule] = (*RecordRuleBuilder)(nil)

type RecordRuleBuilder struct {
	internal *RecordRule
	errors   map[string]cog.BuildErrors
}

func NewRecordRuleBuilder() *RecordRuleBuilder {
	resource := &RecordRule{}
	builder := &RecordRuleBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *RecordRuleBuilder) Build() (RecordRule, error) {
	if err := builder.internal.Validate(); err != nil {
		return RecordRule{}, err
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

func (builder *RecordRuleBuilder) applyDefaults() {
}
