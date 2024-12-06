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
	resource := NewNotificationSettings()
	builder := &NotificationSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *NotificationSettingsBuilder) Build() (NotificationSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return NotificationSettings{}, err
	}

	return *builder.internal, nil
}

// Override the labels by which incoming alerts are grouped together. For example, multiple alerts coming in for
// cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels
// use the special value '...' as the sole label name.
// This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what
// you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
// Must include 'alertname' and 'grafana_folder' if not using '...'.
func (builder *NotificationSettingsBuilder) GroupBy(groupBy []string) *NotificationSettingsBuilder {
	builder.internal.GroupBy = groupBy

	return builder
}

// Override how long to wait before sending a notification about new alerts that are added to a group of alerts for
// which an initial notification has already been sent. (Usually ~5m or more.)
func (builder *NotificationSettingsBuilder) GroupInterval(groupInterval string) *NotificationSettingsBuilder {
	builder.internal.GroupInterval = &groupInterval

	return builder
}

// Override how long to initially wait to send a notification for a group of alerts. Allows to wait for an
// inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.)
func (builder *NotificationSettingsBuilder) GroupWait(groupWait string) *NotificationSettingsBuilder {
	builder.internal.GroupWait = &groupWait

	return builder
}

// Override the times when notifications should be muted. These must match the name of a mute time interval defined
// in the alertmanager configuration mute_time_intervals section. When muted it will not send any notifications, but
// otherwise acts normally.
func (builder *NotificationSettingsBuilder) MuteTimeIntervals(muteTimeIntervals []string) *NotificationSettingsBuilder {
	builder.internal.MuteTimeIntervals = muteTimeIntervals

	return builder
}

// Name of the receiver to send notifications to.
func (builder *NotificationSettingsBuilder) Receiver(receiver string) *NotificationSettingsBuilder {
	builder.internal.Receiver = receiver

	return builder
}

// Override how long to wait before sending a notification again if it has already been sent successfully for an
// alert. (Usually ~3h or more).
// Note that this parameter is implicitly bound by Alertmanager's `--data.retention` configuration flag.
// Notifications will be resent after either repeat_interval or the data retention period have passed, whichever
// occurs first. `repeat_interval` should not be less than `group_interval`.
func (builder *NotificationSettingsBuilder) RepeatInterval(repeatInterval string) *NotificationSettingsBuilder {
	builder.internal.RepeatInterval = &repeatInterval

	return builder
}
