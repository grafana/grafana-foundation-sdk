<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Time picker configuration
 * It defines the default config for the time picker and the refresh picker for the specific dashboard.
 */
class TimePickerConfig implements \JsonSerializable
{
    /**
     * Whether timepicker is visible or not.
     */
    public ?bool $hidden;

    /**
     * Interval options available in the refresh picker dropdown.
     * @var array<string>|null
     */
    public ?array $refreshIntervals;

    /**
     * Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
     * @var array<string>|null
     */
    public ?array $timeOptions;

    /**
     * Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
     */
    public ?string $nowDelay;

    /**
     * @param bool|null $hidden
     * @param array<string>|null $refreshIntervals
     * @param array<string>|null $timeOptions
     * @param string|null $nowDelay
     */
    public function __construct(?bool $hidden = null, ?array $refreshIntervals = null, ?array $timeOptions = null, ?string $nowDelay = null)
    {
        $this->hidden = $hidden;
        $this->refreshIntervals = $refreshIntervals;
        $this->timeOptions = $timeOptions;
        $this->nowDelay = $nowDelay;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{hidden?: bool, refresh_intervals?: array<string>, time_options?: array<string>, nowDelay?: string} $inputData */
        $data = $inputData;
        return new self(
            hidden: $data["hidden"] ?? null,
            refreshIntervals: $data["refresh_intervals"] ?? null,
            timeOptions: $data["time_options"] ?? null,
            nowDelay: $data["nowDelay"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->hidden)) {
            $data["hidden"] = $this->hidden;
        }
        if (isset($this->refreshIntervals)) {
            $data["refresh_intervals"] = $this->refreshIntervals;
        }
        if (isset($this->timeOptions)) {
            $data["time_options"] = $this->timeOptions;
        }
        if (isset($this->nowDelay)) {
            $data["nowDelay"] = $this->nowDelay;
        }
        return $data;
    }
}
