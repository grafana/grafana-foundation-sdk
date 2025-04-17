<?php

namespace Grafana\Foundation\Dashboard;

class TimePicker implements \JsonSerializable
{
    /**
     * Whether timepicker is visible or not.
     */
    public bool $hidden;

    /**
     * Interval options available in the refresh picker dropdown.
     * @var array<string>
     */
    public array $refreshIntervals;

    /**
     * Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.
     */
    public bool $collapse;

    /**
     * Whether timepicker is enabled or not. Has no effect on provisioned dashboard.
     */
    public bool $enable;

    /**
     * Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
     * @var array<string>
     */
    public array $timeOptions;

    /**
     * @param bool|null $hidden
     * @param array<string>|null $refreshIntervals
     * @param bool|null $collapse
     * @param bool|null $enable
     * @param array<string>|null $timeOptions
     */
    public function __construct(?bool $hidden = null, ?array $refreshIntervals = null, ?bool $collapse = null, ?bool $enable = null, ?array $timeOptions = null)
    {
        $this->hidden = $hidden ?: false;
        $this->refreshIntervals = $refreshIntervals ?: ["5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"];
        $this->collapse = $collapse ?: false;
        $this->enable = $enable ?: true;
        $this->timeOptions = $timeOptions ?: ["5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d"];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{hidden?: bool, refresh_intervals?: array<string>, collapse?: bool, enable?: bool, time_options?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            hidden: $data["hidden"] ?? null,
            refreshIntervals: $data["refresh_intervals"] ?? null,
            collapse: $data["collapse"] ?? null,
            enable: $data["enable"] ?? null,
            timeOptions: $data["time_options"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->hidden = $this->hidden;
        $data->refresh_intervals = $this->refreshIntervals;
        $data->collapse = $this->collapse;
        $data->enable = $this->enable;
        $data->time_options = $this->timeOptions;
        return $data;
    }
}
