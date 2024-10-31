<?php

namespace Grafana\Foundation\Alerting;

/**
 * TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
 * within the interval.
 */
class TimeInterval implements \JsonSerializable
{
    /**
     * @var array<string>|null
     */
    public ?array $daysOfMonth;

    public ?string $location;

    /**
     * @var array<string>|null
     */
    public ?array $months;

    /**
     * @var array<\Grafana\Foundation\Alerting\TimeRange>|null
     */
    public ?array $times;

    /**
     * @var array<string>|null
     */
    public ?array $weekdays;

    /**
     * @var array<string>|null
     */
    public ?array $years;

    /**
     * @param array<string>|null $daysOfMonth
     * @param string|null $location
     * @param array<string>|null $months
     * @param array<\Grafana\Foundation\Alerting\TimeRange>|null $times
     * @param array<string>|null $weekdays
     * @param array<string>|null $years
     */
    public function __construct(?array $daysOfMonth = null, ?string $location = null, ?array $months = null, ?array $times = null, ?array $weekdays = null, ?array $years = null)
    {
        $this->daysOfMonth = $daysOfMonth;
        $this->location = $location;
        $this->months = $months;
        $this->times = $times;
        $this->weekdays = $weekdays;
        $this->years = $years;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{days_of_month?: array<string>, location?: string, months?: array<string>, times?: array<mixed>, weekdays?: array<string>, years?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            daysOfMonth: $data["days_of_month"] ?? null,
            location: $data["location"] ?? null,
            months: $data["months"] ?? null,
            times: array_filter(array_map((function($input) {
    	/** @var array{from?: string, to?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\TimeRange::fromArray($val);
    }), $data["times"] ?? [])),
            weekdays: $data["weekdays"] ?? null,
            years: $data["years"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->daysOfMonth)) {
            $data["days_of_month"] = $this->daysOfMonth;
        }
        if (isset($this->location)) {
            $data["location"] = $this->location;
        }
        if (isset($this->months)) {
            $data["months"] = $this->months;
        }
        if (isset($this->times)) {
            $data["times"] = $this->times;
        }
        if (isset($this->weekdays)) {
            $data["weekdays"] = $this->weekdays;
        }
        if (isset($this->years)) {
            $data["years"] = $this->years;
        }
        return $data;
    }
}
