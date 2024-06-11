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
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("RecordRule", err)...)
	}

	if len(errs) != 0 {
		return RecordRule{}, errs
	}

	return *builder.internal, nil
}

func (builder *RecordRuleBuilder) From(from string) *RecordRuleBuilder {
	builder.internal.From = from

	return builder
}

func (builder *RecordRuleBuilder) Metric(metric string) *RecordRuleBuilder {
	builder.internal.Metric = metric

	return builder
}

func (builder *RecordRuleBuilder) applyDefaults() {
}
