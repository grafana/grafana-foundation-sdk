// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RuleSearchHitFields] = (*RuleSearchHitFieldsBuilder)(nil)

// #RuleSearchHitFields is the per-kind field payload returned on each hit.
// It carries the union of alert- and recording-rule search fields; only the
// fields relevant to a hit's kind are populated. This maps to the kind's
// declared searchFields.
type RuleSearchHitFieldsBuilder struct {
	internal *RuleSearchHitFields
	errors   cog.BuildErrors
}

func NewRuleSearchHitFieldsBuilder() *RuleSearchHitFieldsBuilder {
	resource := NewRuleSearchHitFields()
	builder := &RuleSearchHitFieldsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) Build() (RuleSearchHitFields, error) {
	if err := builder.internal.Validate(); err != nil {
		return RuleSearchHitFields{}, err
	}

	if len(builder.errors) > 0 {
		return RuleSearchHitFields{}, cog.MakeBuildErrors("alertingv0alpha1.ruleSearchHitFields", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RuleSearchHitFieldsBuilder) RecordError(path string, err error) *RuleSearchHitFieldsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *RuleSearchHitFieldsBuilder) Title(title string) *RuleSearchHitFieldsBuilder {
	builder.internal.Title = &title

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) Folder(folder string) *RuleSearchHitFieldsBuilder {
	builder.internal.Folder = &folder

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) Type(typeArg string) *RuleSearchHitFieldsBuilder {
	builder.internal.Type = &typeArg

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) Interval(interval string) *RuleSearchHitFieldsBuilder {
	builder.internal.Interval = &interval

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) Paused(paused bool) *RuleSearchHitFieldsBuilder {
	builder.internal.Paused = &paused

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) Labels(labels map[string]string) *RuleSearchHitFieldsBuilder {
	builder.internal.Labels = labels

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) DatasourceUIDs(datasourceUIDs []string) *RuleSearchHitFieldsBuilder {
	builder.internal.DatasourceUIDs = datasourceUIDs

	return builder
}

// Alert-rule fields.
func (builder *RuleSearchHitFieldsBuilder) Annotations(annotations map[string]string) *RuleSearchHitFieldsBuilder {
	builder.internal.Annotations = annotations

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) For(forArg string) *RuleSearchHitFieldsBuilder {
	builder.internal.For = &forArg

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) KeepFiringFor(keepFiringFor string) *RuleSearchHitFieldsBuilder {
	builder.internal.KeepFiringFor = &keepFiringFor

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) DashboardUID(dashboardUID string) *RuleSearchHitFieldsBuilder {
	builder.internal.DashboardUID = &dashboardUID

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) PanelID(panelID int64) *RuleSearchHitFieldsBuilder {
	builder.internal.PanelID = &panelID

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) Receiver(receiver string) *RuleSearchHitFieldsBuilder {
	builder.internal.Receiver = &receiver

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) NotificationType(notificationType string) *RuleSearchHitFieldsBuilder {
	builder.internal.NotificationType = &notificationType

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) RoutingTree(routingTree string) *RuleSearchHitFieldsBuilder {
	builder.internal.RoutingTree = &routingTree

	return builder
}

// Recording-rule fields.
func (builder *RuleSearchHitFieldsBuilder) Metric(metric string) *RuleSearchHitFieldsBuilder {
	builder.internal.Metric = &metric

	return builder
}

func (builder *RuleSearchHitFieldsBuilder) TargetDatasourceUID(targetDatasourceUID string) *RuleSearchHitFieldsBuilder {
	builder.internal.TargetDatasourceUID = &targetDatasourceUID

	return builder
}
