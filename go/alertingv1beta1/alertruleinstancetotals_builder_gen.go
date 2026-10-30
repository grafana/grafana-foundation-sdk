// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AlertRuleInstanceTotals] = (*AlertRuleInstanceTotalsBuilder)(nil)

// Count of alert instances per state. error also counts instances whose evaluation
// errored but were mapped to another state via execErrState, so it can overlap.
type AlertRuleInstanceTotalsBuilder struct {
	internal *AlertRuleInstanceTotals
	errors   cog.BuildErrors
}

func NewAlertRuleInstanceTotalsBuilder() *AlertRuleInstanceTotalsBuilder {
	resource := NewAlertRuleInstanceTotals()
	builder := &AlertRuleInstanceTotalsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AlertRuleInstanceTotalsBuilder) Build() (AlertRuleInstanceTotals, error) {
	if err := builder.internal.Validate(); err != nil {
		return AlertRuleInstanceTotals{}, err
	}

	if len(builder.errors) > 0 {
		return AlertRuleInstanceTotals{}, cog.MakeBuildErrors("alertingv1beta1.alertRuleInstanceTotals", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AlertRuleInstanceTotalsBuilder) RecordError(path string, err error) *AlertRuleInstanceTotalsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *AlertRuleInstanceTotalsBuilder) Healthy(healthy int64) *AlertRuleInstanceTotalsBuilder {
	builder.internal.Healthy = &healthy

	return builder
}

func (builder *AlertRuleInstanceTotalsBuilder) Firing(firing int64) *AlertRuleInstanceTotalsBuilder {
	builder.internal.Firing = &firing

	return builder
}

func (builder *AlertRuleInstanceTotalsBuilder) Pending(pending int64) *AlertRuleInstanceTotalsBuilder {
	builder.internal.Pending = &pending

	return builder
}

func (builder *AlertRuleInstanceTotalsBuilder) Recovering(recovering int64) *AlertRuleInstanceTotalsBuilder {
	builder.internal.Recovering = &recovering

	return builder
}

func (builder *AlertRuleInstanceTotalsBuilder) Nodata(nodata int64) *AlertRuleInstanceTotalsBuilder {
	builder.internal.Nodata = &nodata

	return builder
}

func (builder *AlertRuleInstanceTotalsBuilder) Error(errorArg int64) *AlertRuleInstanceTotalsBuilder {
	builder.internal.Error = &errorArg

	return builder
}
