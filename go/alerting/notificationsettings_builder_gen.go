// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NotificationSettings] = (*NotificationSettingsBuilder)(nil)

type NotificationSettingsBuilder struct {
	internal *NotificationSettings
	errors   map[string]cog.BuildErrors
}

func NewNotificationSettingsBuilder() *NotificationSettingsBuilder {
	resource := &NotificationSettings{}
	builder := &NotificationSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *NotificationSettingsBuilder) Build() (NotificationSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("NotificationSettings", err)...)
	}

	if len(errs) != 0 {
		return NotificationSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *NotificationSettingsBuilder) GroupBy(groupBy []string) *NotificationSettingsBuilder {
	builder.internal.GroupBy = groupBy

	return builder
}

func (builder *NotificationSettingsBuilder) GroupInterval(groupInterval string) *NotificationSettingsBuilder {
	builder.internal.GroupInterval = &groupInterval

	return builder
}

func (builder *NotificationSettingsBuilder) GroupWait(groupWait string) *NotificationSettingsBuilder {
	builder.internal.GroupWait = &groupWait

	return builder
}

func (builder *NotificationSettingsBuilder) MuteTimeIntervals(muteTimeIntervals []string) *NotificationSettingsBuilder {
	builder.internal.MuteTimeIntervals = muteTimeIntervals

	return builder
}

func (builder *NotificationSettingsBuilder) Receiver(receiver string) *NotificationSettingsBuilder {
	builder.internal.Receiver = receiver

	return builder
}

func (builder *NotificationSettingsBuilder) RepeatInterval(repeatInterval string) *NotificationSettingsBuilder {
	builder.internal.RepeatInterval = &repeatInterval

	return builder
}

func (builder *NotificationSettingsBuilder) applyDefaults() {
	builder.GroupBy([]string{"alertname", "grafana_folder"})
}
