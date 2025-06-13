// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NotificationTemplate] = (*NotificationTemplateBuilder)(nil)

type NotificationTemplateBuilder struct {
	internal *NotificationTemplate
	errors   cog.BuildErrors
}

func NewNotificationTemplateBuilder() *NotificationTemplateBuilder {
	resource := NewNotificationTemplate()
	builder := &NotificationTemplateBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *NotificationTemplateBuilder) Build() (NotificationTemplate, error) {
	if err := builder.internal.Validate(); err != nil {
		return NotificationTemplate{}, err
	}

	if len(builder.errors) > 0 {
		return NotificationTemplate{}, cog.MakeBuildErrors("alerting.notificationTemplate", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *NotificationTemplateBuilder) Name(name string) *NotificationTemplateBuilder {
	builder.internal.Name = &name

	return builder
}

func (builder *NotificationTemplateBuilder) Provenance(provenance Provenance) *NotificationTemplateBuilder {
	builder.internal.Provenance = &provenance

	return builder
}

func (builder *NotificationTemplateBuilder) Template(template string) *NotificationTemplateBuilder {
	builder.internal.Template = &template

	return builder
}
