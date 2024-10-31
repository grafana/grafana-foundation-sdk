---
title: <span class="badge builder"></span> NotificationSettingsBuilder
---
# <span class="badge builder"></span> NotificationSettingsBuilder

## Constructor

```php
new NotificationSettingsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> groupBy

Override the labels by which incoming alerts are grouped together. For example, multiple alerts coming in for

cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels

use the special value '...' as the sole label name.

This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what

you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.

Must include 'alertname' and 'grafana_folder' if not using '...'.

@param array<string> $groupBy

```php
groupBy(array $groupBy)
```

### <span class="badge object-method"></span> groupInterval

Override how long to wait before sending a notification about new alerts that are added to a group of alerts for

which an initial notification has already been sent. (Usually ~5m or more.)

```php
groupInterval(string $groupInterval)
```

### <span class="badge object-method"></span> groupWait

Override how long to initially wait to send a notification for a group of alerts. Allows to wait for an

inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.)

```php
groupWait(string $groupWait)
```

### <span class="badge object-method"></span> muteTimeIntervals

Override the times when notifications should be muted. These must match the name of a mute time interval defined

in the alertmanager configuration mute_time_intervals section. When muted it will not send any notifications, but

otherwise acts normally.

@param array<string> $muteTimeIntervals

```php
muteTimeIntervals(array $muteTimeIntervals)
```

### <span class="badge object-method"></span> receiver

Name of the receiver to send notifications to.

```php
receiver(string $receiver)
```

### <span class="badge object-method"></span> repeatInterval

Override how long to wait before sending a notification again if it has already been sent successfully for an

alert. (Usually ~3h or more).

Note that this parameter is implicitly bound by Alertmanager's `--data.retention` configuration flag.

Notifications will be resent after either repeat_interval or the data retention period have passed, whichever

occurs first. `repeat_interval` should not be less than `group_interval`.

```php
repeatInterval(string $repeatInterval)
```

## See also

 * <span class="badge object-type-class"></span> [NotificationSettings](./object-NotificationSettings.md)
