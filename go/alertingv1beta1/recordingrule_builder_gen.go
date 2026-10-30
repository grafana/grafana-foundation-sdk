// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	resource "github.com/grafana/grafana-foundation-sdk/go/resource"
)

var _ cog.Builder[RecordingRuleSpec] = (*RecordingRuleBuilder)(nil)

type RecordingRuleBuilder struct {
	internal *RecordingRuleSpec
	errors   cog.BuildErrors
}

func NewRecordingRuleBuilder() *RecordingRuleBuilder {
	resource := NewRecordingRuleSpec()
	builder := &RecordingRuleBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

// Creates a resource manifest from a RecordingRule.
func RecordingRuleManifest(name string, folder cog.Builder[RecordingRuleSpec]) *resource.ManifestBuilder {
	builder := resource.NewManifestBuilder()

	builder.ApiVersion("rules.alerting.grafana.app/v1beta1")

	builder.Kind("RecordingRule")

	builder.Metadata(resource.Named(name))
	folderResource, err := folder.Build()
	if err != nil {
		builder.RecordError("RecordingRuleManifest(folder)", err)
		return builder
	}
	builder.Spec(folderResource)

	return builder
}

func (builder *RecordingRuleBuilder) Build() (RecordingRuleSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return RecordingRuleSpec{}, err
	}

	if len(builder.errors) > 0 {
		return RecordingRuleSpec{}, cog.MakeBuildErrors("alertingv1beta1.recordingRule", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RecordingRuleBuilder) RecordError(path string, err error) *RecordingRuleBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *RecordingRuleBuilder) Title(title string) *RecordingRuleBuilder {
	builder.internal.Title = title

	return builder
}

func (builder *RecordingRuleBuilder) Paused(paused bool) *RecordingRuleBuilder {
	builder.internal.Paused = &paused

	return builder
}

func (builder *RecordingRuleBuilder) Trigger(trigger cog.Builder[IntervalTrigger]) *RecordingRuleBuilder {
	triggerResource, err := trigger.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Trigger = triggerResource

	return builder
}

func (builder *RecordingRuleBuilder) Labels(labels map[string]TemplateString) *RecordingRuleBuilder {
	builder.internal.Labels = labels

	return builder
}

func (builder *RecordingRuleBuilder) Metric(metric MetricName) *RecordingRuleBuilder {
	builder.internal.Metric = metric

	return builder
}

func (builder *RecordingRuleBuilder) Expressions(expressions ExpressionMap) *RecordingRuleBuilder {
	builder.internal.Expressions = expressions

	return builder
}

func (builder *RecordingRuleBuilder) TargetDatasourceUID(targetDatasourceUID DatasourceUID) *RecordingRuleBuilder {
	builder.internal.TargetDatasourceUID = targetDatasourceUID

	return builder
}
