// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	resource "github.com/grafana/grafana-foundation-sdk/go/resource"
)

var _ cog.Builder[AlertRuleSpec] = (*AlertRuleBuilder)(nil)

type AlertRuleBuilder struct {
	internal *AlertRuleSpec
	errors   cog.BuildErrors
}

func NewAlertRuleBuilder() *AlertRuleBuilder {
	resource := NewAlertRuleSpec()
	builder := &AlertRuleBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

// Creates a resource manifest from an AlertRule.
func AlertRuleManifest(name string, folder cog.Builder[AlertRuleSpec]) *resource.ManifestBuilder {
	builder := resource.NewManifestBuilder()

	builder.ApiVersion("rules.alerting.grafana.app/v1beta1")

	builder.Kind("AlertRule")

	builder.Metadata(resource.Named(name))
	folderResource, err := folder.Build()
	if err != nil {
		builder.RecordError("AlertRuleManifest(folder)", err)
		return builder
	}
	builder.Spec(folderResource)

	return builder
}

func (builder *AlertRuleBuilder) Build() (AlertRuleSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return AlertRuleSpec{}, err
	}

	if len(builder.errors) > 0 {
		return AlertRuleSpec{}, cog.MakeBuildErrors("alertingv1beta1.alertRule", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AlertRuleBuilder) RecordError(path string, err error) *AlertRuleBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *AlertRuleBuilder) Title(title string) *AlertRuleBuilder {
	builder.internal.Title = title

	return builder
}

func (builder *AlertRuleBuilder) Paused(paused bool) *AlertRuleBuilder {
	builder.internal.Paused = &paused

	return builder
}

func (builder *AlertRuleBuilder) Trigger(trigger cog.Builder[IntervalTrigger]) *AlertRuleBuilder {
	triggerResource, err := trigger.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Trigger = triggerResource

	return builder
}

func (builder *AlertRuleBuilder) Labels(labels map[string]TemplateString) *AlertRuleBuilder {
	builder.internal.Labels = labels

	return builder
}

func (builder *AlertRuleBuilder) Annotations(annotations map[string]TemplateString) *AlertRuleBuilder {
	builder.internal.Annotations = annotations

	return builder
}

func (builder *AlertRuleBuilder) For(forArg string) *AlertRuleBuilder {
	builder.internal.For = &forArg

	return builder
}

func (builder *AlertRuleBuilder) KeepFiringFor(keepFiringFor string) *AlertRuleBuilder {
	builder.internal.KeepFiringFor = &keepFiringFor

	return builder
}

func (builder *AlertRuleBuilder) MissingSeriesEvalsToResolve(missingSeriesEvalsToResolve int64) *AlertRuleBuilder {
	builder.internal.MissingSeriesEvalsToResolve = &missingSeriesEvalsToResolve

	return builder
}

func (builder *AlertRuleBuilder) NoDataState(noDataState NoDataState) *AlertRuleBuilder {
	builder.internal.NoDataState = noDataState

	return builder
}

func (builder *AlertRuleBuilder) ExecErrState(execErrState ExecErrState) *AlertRuleBuilder {
	builder.internal.ExecErrState = execErrState

	return builder
}

func (builder *AlertRuleBuilder) NotificationSettings(notificationSettings NotificationSettings) *AlertRuleBuilder {
	builder.internal.NotificationSettings = &notificationSettings

	return builder
}

func (builder *AlertRuleBuilder) Expressions(expressions ExpressionMap) *AlertRuleBuilder {
	builder.internal.Expressions = expressions

	return builder
}

func (builder *AlertRuleBuilder) PanelRef(panelRef cog.Builder[PanelRef]) *AlertRuleBuilder {
	panelRefResource, err := panelRef.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.PanelRef = &panelRefResource

	return builder
}
