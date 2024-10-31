---
title: <span class="badge object-type-class"></span> NotificationSettings
---
# <span class="badge object-type-class"></span> NotificationSettings

## Definition

```php
class NotificationSettings implements \JsonSerializable
{
    /**
     * Override the labels by which incoming alerts are grouped together. For example, multiple alerts coming in for
     * cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels
     * use the special value '...' as the sole label name.
     * This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what
     * you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
     * Must include 'alertname' and 'grafana_folder' if not using '...'.
     * @var array<string>|null
     */
    public ?array $groupBy;

    /**
     * Override how long to wait before sending a notification about new alerts that are added to a group of alerts for
     * which an initial notification has already been sent. (Usually ~5m or more.)
     */
    public ?string $groupInterval;

    /**
     * Override how long to initially wait to send a notification for a group of alerts. Allows to wait for an
     * inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.)
     */
    public ?string $groupWait;

    /**
     * Override the times when notifications should be muted. These must match the name of a mute time interval defined
     * in the alertmanager configuration mute_time_intervals section. When muted it will not send any notifications, but
     * otherwise acts normally.
     * @var array<string>|null
     */
    public ?array $muteTimeIntervals;

    /**
     * Name of the receiver to send notifications to.
     */
    public string $receiver;

    /**
     * Override how long to wait before sending a notification again if it has already been sent successfully for an
     * alert. (Usually ~3h or more).
     * Note that this parameter is implicitly bound by Alertmanager's `--data.retention` configuration flag.
     * Notifications will be resent after either repeat_interval or the data retention period have passed, whichever
     * occurs first. `repeat_interval` should not be less than `group_interval`.
     */
    public ?string $repeatInterval;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [NotificationSettingsBuilder](./builder-NotificationSettingsBuilder.md)
