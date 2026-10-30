// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	"time"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RecordingRuleStatus] = (*RecordingRuleStatusBuilder)(nil)

type RecordingRuleStatusBuilder struct {
	internal *RecordingRuleStatus
	errors   cog.BuildErrors
}

func NewRecordingRuleStatusBuilder() *RecordingRuleStatusBuilder {
	resource := NewRecordingRuleStatus()
	builder := &RecordingRuleStatusBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RecordingRuleStatusBuilder) Build() (RecordingRuleStatus, error) {
	if err := builder.internal.Validate(); err != nil {
		return RecordingRuleStatus{}, err
	}

	if len(builder.errors) > 0 {
		return RecordingRuleStatus{}, cog.MakeBuildErrors("alertingv1beta1.recordingRuleStatus", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RecordingRuleStatusBuilder) RecordError(path string, err error) *RecordingRuleStatusBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *RecordingRuleStatusBuilder) Health(health RecordingRuleHealth) *RecordingRuleStatusBuilder {
	builder.internal.Health = &health

	return builder
}

func (builder *RecordingRuleStatusBuilder) LastEvaluationTime(lastEvaluationTime time.Time) *RecordingRuleStatusBuilder {
	builder.internal.LastEvaluationTime = &lastEvaluationTime

	return builder
}

// duration of the last evaluation in seconds
func (builder *RecordingRuleStatusBuilder) EvaluationDuration(evaluationDuration float64) *RecordingRuleStatusBuilder {
	builder.internal.EvaluationDuration = &evaluationDuration

	return builder
}

func (builder *RecordingRuleStatusBuilder) LastError(lastError string) *RecordingRuleStatusBuilder {
	builder.internal.LastError = &lastError

	return builder
}
