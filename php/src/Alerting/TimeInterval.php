<?php

namespace Grafana\Foundation\Alerting;

class TimeInterval implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Alerting\TimeRange>|null
     */
    public ?array $times;

    /**
     * @var array<\Grafana\Foundation\Alerting\WeekdayRange>|null
     */
    public ?array $weekdays;

    /**
     * @var array<\Grafana\Foundation\Alerting\DayOfMonthRange>|null
     */
    public ?array $daysOfMonth;

    /**
     * @var array<\Grafana\Foundation\Alerting\MonthRange>|null
     */
    public ?array $months;

    /**
     * @var array<\Grafana\Foundation\Alerting\YearRange>|null
     */
    public ?array $years;

    public string $location;

    /**
     * @param array<\Grafana\Foundation\Alerting\TimeRange>|null $times
     * @param array<\Grafana\Foundation\Alerting\WeekdayRange>|null $weekdays
     * @param array<\Grafana\Foundation\Alerting\DayOfMonthRange>|null $daysOfMonth
     * @param array<\Grafana\Foundation\Alerting\MonthRange>|null $months
     * @param array<\Grafana\Foundation\Alerting\YearRange>|null $years
     * @param string|null $location
     */
    public function __construct(?array $times = null, ?array $weekdays = null, ?array $daysOfMonth = null, ?array $months = null, ?array $years = null, ?string $location = null)
    {
        $this->times = $times;
        $this->weekdays = $weekdays;
        $this->daysOfMonth = $daysOfMonth;
        $this->months = $months;
        $this->years = $years;
        $this->location = $location ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{times?: array<mixed>, weekdays?: array<mixed>, days_of_month?: array<mixed>, months?: array<mixed>, years?: array<mixed>, location?: string} $inputData */
        $data = $inputData;
        return new self(
            times: array_filter(array_map((function($input) {
    	/** @var array{from?: string, to?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\TimeRange::fromArray($val);
    }), $data["times"] ?? [])),
            weekdays: array_filter(array_map((function($input) {
    	/** @var array{begin?: int, end?: int} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\WeekdayRange::fromArray($val);
    }), $data["weekdays"] ?? [])),
            daysOfMonth: array_filter(array_map((function($input) {
    	/** @var array{begin?: int, end?: int} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\DayOfMonthRange::fromArray($val);
    }), $data["days_of_month"] ?? [])),
            months: array_filter(array_map((function($input) {
    	/** @var array{begin?: int, end?: int} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\MonthRange::fromArray($val);
    }), $data["months"] ?? [])),
            years: array_filter(array_map((function($input) {
    	/** @var array{begin?: int, end?: int} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\YearRange::fromArray($val);
    }), $data["years"] ?? [])),
            location: $data["location"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->location = $this->location;
        if (isset($this->times)) {
            $data->times = $this->times;
        }
        if (isset($this->weekdays)) {
            $data->weekdays = $this->weekdays;
        }
        if (isset($this->daysOfMonth)) {
            $data->days_of_month = $this->daysOfMonth;
        }
        if (isset($this->months)) {
            $data->months = $this->months;
        }
        if (isset($this->years)) {
            $data->years = $this->years;
        }
        return $data;
    }
}
