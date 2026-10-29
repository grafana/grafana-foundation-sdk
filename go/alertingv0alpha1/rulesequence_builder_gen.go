// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	resource "github.com/grafana/grafana-foundation-sdk/go/resource"
)

var _ cog.Builder[RuleSequenceSpec] = (*RuleSequenceBuilder)(nil)

type RuleSequenceBuilder struct {
	internal *RuleSequenceSpec
	errors   cog.BuildErrors
}

func NewRuleSequenceBuilder() *RuleSequenceBuilder {
	resource := NewRuleSequenceSpec()
	builder := &RuleSequenceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

// Creates a resource manifest from a RuleSequence.
func RuleSequenceManifest(name string, folder cog.Builder[RuleSequenceSpec]) *resource.ManifestBuilder {
	builder := resource.NewManifestBuilder()

	builder.ApiVersion("rules.alerting.grafana.app/v0alpha1")

	builder.Kind("RuleSequence")

	builder.Metadata(resource.Named(name))
	folderResource, err := folder.Build()
	if err != nil {
		builder.RecordError("RuleSequenceManifest(folder)", err)
		return builder
	}
	builder.Spec(folderResource)

	return builder
}

func (builder *RuleSequenceBuilder) Build() (RuleSequenceSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return RuleSequenceSpec{}, err
	}

	if len(builder.errors) > 0 {
		return RuleSequenceSpec{}, cog.MakeBuildErrors("alertingv0alpha1.ruleSequence", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RuleSequenceBuilder) RecordError(path string, err error) *RuleSequenceBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *RuleSequenceBuilder) Trigger(trigger cog.Builder[IntervalTrigger]) *RuleSequenceBuilder {
	triggerResource, err := trigger.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Trigger = triggerResource

	return builder
}

func (builder *RuleSequenceBuilder) RecordingRules(recordingRules []cog.Builder[RuleRef]) *RuleSequenceBuilder {
	recordingRulesResources := make([]RuleRef, 0, len(recordingRules))
	for _, r1 := range recordingRules {
		recordingRulesDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		recordingRulesResources = append(recordingRulesResources, recordingRulesDepth1)
	}
	builder.internal.RecordingRules = recordingRulesResources

	return builder
}

func (builder *RuleSequenceBuilder) AlertingRules(alertingRules []cog.Builder[RuleRef]) *RuleSequenceBuilder {
	alertingRulesResources := make([]RuleRef, 0, len(alertingRules))
	for _, r1 := range alertingRules {
		alertingRulesDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		alertingRulesResources = append(alertingRulesResources, alertingRulesDepth1)
	}
	builder.internal.AlertingRules = alertingRulesResources

	return builder
}
