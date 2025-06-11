---
title: <span class="badge object-type-struct"></span> NotificationSettings
---
# <span class="badge object-type-struct"></span> NotificationSettings

## Definition

```go
type NotificationSettings struct {
    // Override the times when notifications should not be muted. These must match the name of a mute time interval defined
    // in the alertmanager configuration time_intervals section. All notifications will be suppressed unless they are sent
    // at the time that matches any interval.
    ActiveTimeIntervals []string `json:"active_time_intervals,omitempty"`
    // Override the labels by which incoming alerts are grouped together. For example, multiple alerts coming in for
    // cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels
    // use the special value '...' as the sole label name.
    // This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what
    // you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
    // Must include 'alertname' and 'grafana_folder' if not using '...'.
    GroupBy []string `json:"group_by,omitempty"`
    // Override how long to wait before sending a notification about new alerts that are added to a group of alerts for
    // which an initial notification has already been sent. (Usually ~5m or more.)
    GroupInterval *string `json:"group_interval,omitempty"`
    // Override how long to initially wait to send a notification for a group of alerts. Allows to wait for an
    // inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.)
    GroupWait *string `json:"group_wait,omitempty"`
    // Override the times when notifications should be muted. These must match the name of a mute time interval defined
    // in the alertmanager configuration time_intervals section. When muted it will not send any notifications, but
    // otherwise acts normally.
    MuteTimeIntervals []string `json:"mute_time_intervals,omitempty"`
    // Name of the receiver to send notifications to.
    Receiver string `json:"receiver"`
    // Override how long to wait before sending a notification again if it has already been sent successfully for an
    // alert. (Usually ~3h or more).
    // Note that this parameter is implicitly bound by Alertmanager's `--data.retention` configuration flag.
    // Notifications will be resent after either repeat_interval or the data retention period have passed, whichever
    // occurs first. `repeat_interval` should not be less than `group_interval`.
    RepeatInterval *string `json:"repeat_interval,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NotificationSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (notificationSettings *NotificationSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `NotificationSettings` objects.

```go
func (notificationSettings *NotificationSettings) Equals(other NotificationSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `NotificationSettings` fields for violations and returns them.

```go
func (notificationSettings *NotificationSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [NotificationSettingsBuilder](./builder-NotificationSettingsBuilder.md)
