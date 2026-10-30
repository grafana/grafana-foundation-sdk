// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	"time"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AlertRuleStatus] = (*AlertRuleStatusBuilder)(nil)

type AlertRuleStatusBuilder struct {
	internal *AlertRuleStatus
	errors   cog.BuildErrors
}

func NewAlertRuleStatusBuilder() *AlertRuleStatusBuilder {
	resource := NewAlertRuleStatus()
	builder := &AlertRuleStatusBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AlertRuleStatusBuilder) Build() (AlertRuleStatus, error) {
	if err := builder.internal.Validate(); err != nil {
		return AlertRuleStatus{}, err
	}

	if len(builder.errors) > 0 {
		return AlertRuleStatus{}, cog.MakeBuildErrors("alertingv1beta1.alertRuleStatus", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AlertRuleStatusBuilder) RecordError(path string, err error) *AlertRuleStatusBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *AlertRuleStatusBuilder) Health(health AlertRuleHealth) *AlertRuleStatusBuilder {
	builder.internal.Health = &health

	return builder
}

func (builder *AlertRuleStatusBuilder) State(state AlertRuleState) *AlertRuleStatusBuilder {
	builder.internal.State = &state

	return builder
}

func (builder *AlertRuleStatusBuilder) StateReason(stateReason AlertRuleStateReason) *AlertRuleStatusBuilder {
	builder.internal.StateReason = &stateReason

	return builder
}

func (builder *AlertRuleStatusBuilder) LastEvaluationTime(lastEvaluationTime time.Time) *AlertRuleStatusBuilder {
	builder.internal.LastEvaluationTime = &lastEvaluationTime

	return builder
}

// duration of the last evaluation in seconds
func (builder *AlertRuleStatusBuilder) EvaluationDuration(evaluationDuration float64) *AlertRuleStatusBuilder {
	builder.internal.EvaluationDuration = &evaluationDuration

	return builder
}

func (builder *AlertRuleStatusBuilder) LastError(lastError string) *AlertRuleStatusBuilder {
	builder.internal.LastError = &lastError

	return builder
}

func (builder *AlertRuleStatusBuilder) Totals(totals cog.Builder[AlertRuleInstanceTotals]) *AlertRuleStatusBuilder {
	totalsResource, err := totals.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Totals = &totalsResource

	return builder
}
