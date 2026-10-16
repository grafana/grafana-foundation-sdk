<?php

namespace Grafana\Foundation\Alerting;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\NotificationSettings>
 */
class NotificationSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\NotificationSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alerting\NotificationSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alerting\NotificationSettings
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Override the times when notifications should not be muted. These must match the name of a mute time interval defined
     * in the alertmanager configuration time_intervals section. All notifications will be suppressed unless they are sent
     * at the time that matches any interval.
     * @param array<string> $activeTimeIntervals
     */
    public function activeTimeIntervals(array $activeTimeIntervals): static
    {
        $this->internal->activeTimeIntervals = $activeTimeIntervals;
    
        return $this;
    }

    /**
     * Override the labels by which incoming alerts are grouped together. For example, multiple alerts coming in for
     * cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels
     * use the special value '...' as the sole label name.
     * This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what
     * you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
     * Must include 'alertname' and 'grafana_folder' if not using '...'.
     * @param array<string> $groupBy
     */
    public function groupBy(array $groupBy): static
    {
        $this->internal->groupBy = $groupBy;
    
        return $this;
    }

    /**
     * Override how long to wait before sending a notification about new alerts that are added to a group of alerts for
     * which an initial notification has already been sent. (Usually ~5m or more.)
     */
    public function groupInterval(string $groupInterval): static
    {
        $this->internal->groupInterval = $groupInterval;
    
        return $this;
    }

    /**
     * Override how long to initially wait to send a notification for a group of alerts. Allows to wait for an
     * inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.)
     */
    public function groupWait(string $groupWait): static
    {
        $this->internal->groupWait = $groupWait;
    
        return $this;
    }

    /**
     * Override the times when notifications should be muted. These must match the name of a mute time interval defined
     * in the alertmanager configuration time_intervals section. When muted it will not send any notifications, but
     * otherwise acts normally.
     * @param array<string> $muteTimeIntervals
     */
    public function muteTimeIntervals(array $muteTimeIntervals): static
    {
        $this->internal->muteTimeIntervals = $muteTimeIntervals;
    
        return $this;
    }

    /**
     * Name of the receiver to send notifications to.
     */
    public function receiver(string $receiver): static
    {
        $this->internal->receiver = $receiver;
    
        return $this;
    }

    /**
     * Override how long to wait before sending a notification again if it has already been sent successfully for an
     * alert. (Usually ~3h or more).
     * Note that this parameter is implicitly bound by Alertmanager's `--data.retention` configuration flag.
     * Notifications will be resent after either repeat_interval or the data retention period have passed, whichever
     * occurs first. `repeat_interval` should not be less than `group_interval`.
     */
    public function repeatInterval(string $repeatInterval): static
    {
        $this->internal->repeatInterval = $repeatInterval;
    
        return $this;
    }

}
