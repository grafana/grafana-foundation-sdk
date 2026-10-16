<?php

namespace Grafana\Foundation\Alerting;

class NotificationSettings implements \JsonSerializable
{
    /**
     * Override the times when notifications should not be muted. These must match the name of a mute time interval defined
     * in the alertmanager configuration time_intervals section. All notifications will be suppressed unless they are sent
     * at the time that matches any interval.
     * @var array<string>|null
     */
    public ?array $activeTimeIntervals;

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
     * in the alertmanager configuration time_intervals section. When muted it will not send any notifications, but
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

    /**
     * @param array<string>|null $activeTimeIntervals
     * @param array<string>|null $groupBy
     * @param string|null $groupInterval
     * @param string|null $groupWait
     * @param array<string>|null $muteTimeIntervals
     * @param string|null $receiver
     * @param string|null $repeatInterval
     */
    public function __construct(?array $activeTimeIntervals = null, ?array $groupBy = null, ?string $groupInterval = null, ?string $groupWait = null, ?array $muteTimeIntervals = null, ?string $receiver = null, ?string $repeatInterval = null)
    {
        $this->activeTimeIntervals = $activeTimeIntervals;
        $this->groupBy = $groupBy;
        $this->groupInterval = $groupInterval;
        $this->groupWait = $groupWait;
        $this->muteTimeIntervals = $muteTimeIntervals;
        $this->receiver = $receiver ?: "";
        $this->repeatInterval = $repeatInterval;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{active_time_intervals?: array<string>, group_by?: array<string>, group_interval?: string, group_wait?: string, mute_time_intervals?: array<string>, receiver?: string, repeat_interval?: string} $inputData */
        $data = $inputData;
        return new self(
            activeTimeIntervals: $data["active_time_intervals"] ?? null,
            groupBy: $data["group_by"] ?? null,
            groupInterval: $data["group_interval"] ?? null,
            groupWait: $data["group_wait"] ?? null,
            muteTimeIntervals: $data["mute_time_intervals"] ?? null,
            receiver: $data["receiver"] ?? null,
            repeatInterval: $data["repeat_interval"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->receiver = $this->receiver;
        if (isset($this->activeTimeIntervals)) {
            $data->active_time_intervals = $this->activeTimeIntervals;
        }
        if (isset($this->groupBy)) {
            $data->group_by = $this->groupBy;
        }
        if (isset($this->groupInterval)) {
            $data->group_interval = $this->groupInterval;
        }
        if (isset($this->groupWait)) {
            $data->group_wait = $this->groupWait;
        }
        if (isset($this->muteTimeIntervals)) {
            $data->mute_time_intervals = $this->muteTimeIntervals;
        }
        if (isset($this->repeatInterval)) {
            $data->repeat_interval = $this->repeatInterval;
        }
        return $data;
    }
}
