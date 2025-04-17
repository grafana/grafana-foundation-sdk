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
     * Quick ranges for time picker.
     * @var array<\Grafana\Foundation\Dashboard\TimeOption>|null
     */
    public ?array $quickRanges;

    /**
     * Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
     */
    public ?string $nowDelay;

    /**
     * @param bool|null $hidden
     * @param array<string>|null $refreshIntervals
     * @param array<\Grafana\Foundation\Dashboard\TimeOption>|null $quickRanges
     * @param string|null $nowDelay
     */
    public function __construct(?bool $hidden = null, ?array $refreshIntervals = null, ?array $quickRanges = null, ?string $nowDelay = null)
    {
        $this->hidden = $hidden;
        $this->refreshIntervals = $refreshIntervals;
        $this->quickRanges = $quickRanges;
        $this->nowDelay = $nowDelay;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{hidden?: bool, refresh_intervals?: array<string>, quick_ranges?: array<mixed>, nowDelay?: string} $inputData */
        $data = $inputData;
        return new self(
            hidden: $data["hidden"] ?? null,
            refreshIntervals: $data["refresh_intervals"] ?? null,
            quickRanges: array_filter(array_map((function($input) {
    	/** @var array{display?: string, from?: string, to?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\TimeOption::fromArray($val);
    }), $data["quick_ranges"] ?? [])),
            nowDelay: $data["nowDelay"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->hidden)) {
            $data->hidden = $this->hidden;
        }
        if (isset($this->refreshIntervals)) {
            $data->refresh_intervals = $this->refreshIntervals;
        }
        if (isset($this->quickRanges)) {
            $data->quick_ranges = $this->quickRanges;
        }
        if (isset($this->nowDelay)) {
            $data->nowDelay = $this->nowDelay;
        }
        return $data;
    }
}
