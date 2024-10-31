// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NotificationTemplate] = (*NotificationTemplateBuilder)(nil)

type NotificationTemplateBuilder struct {
	internal *NotificationTemplate
	errors   map[string]cog.BuildErrors
}

func NewNotificationTemplateBuilder() *NotificationTemplateBuilder {
	resource := &NotificationTemplate{}
	builder := &NotificationTemplateBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *NotificationTemplateBuilder) Build() (NotificationTemplate, error) {
	if err := builder.internal.Validate(); err != nil {
		return NotificationTemplate{}, err
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

func (builder *NotificationTemplateBuilder) applyDefaults() {
}
