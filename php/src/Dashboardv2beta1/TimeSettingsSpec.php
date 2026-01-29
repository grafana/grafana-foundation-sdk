<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Time configuration
 * It defines the default time config for the time picker, the refresh picker for the specific dashboard.
 */
class TimeSettingsSpec implements \JsonSerializable
{
    /**
     * Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
     */
    public ?string $timezone;

    /**
     * Start time range for dashboard.
     * Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
     */
    public string $from;

    /**
     * End time range for dashboard.
     * Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
     */
    public string $to;

    /**
     * Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
     * v1: refresh
     */
    public string $autoRefresh;

    /**
     * Interval options available in the refresh picker dropdown.
     * v1: timepicker.refresh_intervals
     * @var array<string>
     */
    public array $autoRefreshIntervals;

    /**
     * Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
     * v1: timepicker.quick_ranges , not exposed in the UI
     * @var array<\Grafana\Foundation\Dashboardv2beta1\TimeRangeOption>|null
     */
    public ?array $quickRanges;

    /**
     * Whether timepicker is visible or not.
     * v1: timepicker.hidden
     */
    public bool $hideTimepicker;

    /**
     * Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
     */
    public ?\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpecWeekStart $weekStart;

    /**
     * The month that the fiscal year starts on. 0 = January, 11 = December
     */
    public int $fiscalYearStartMonth;

    /**
     * Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
     * v1: timepicker.nowDelay
     */
    public ?string $nowDelay;

    /**
     * @param string|null $timezone
     * @param string|null $from
     * @param string|null $to
     * @param string|null $autoRefresh
     * @param array<string>|null $autoRefreshIntervals
     * @param array<\Grafana\Foundation\Dashboardv2beta1\TimeRangeOption>|null $quickRanges
     * @param bool|null $hideTimepicker
     * @param \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpecWeekStart|null $weekStart
     * @param int|null $fiscalYearStartMonth
     * @param string|null $nowDelay
     */
    public function __construct(?string $timezone = null, ?string $from = null, ?string $to = null, ?string $autoRefresh = null, ?array $autoRefreshIntervals = null, ?array $quickRanges = null, ?bool $hideTimepicker = null, ?\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpecWeekStart $weekStart = null, ?int $fiscalYearStartMonth = null, ?string $nowDelay = null)
    {
        $this->timezone = $timezone;
        $this->from = $from ?: "now-6h";
        $this->to = $to ?: "now";
        $this->autoRefresh = $autoRefresh ?: "";
        $this->autoRefreshIntervals = $autoRefreshIntervals ?: ["5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"];
        $this->quickRanges = $quickRanges;
        $this->hideTimepicker = $hideTimepicker ?: false;
        $this->weekStart = $weekStart;
        $this->fiscalYearStartMonth = $fiscalYearStartMonth ?: 0;
        $this->nowDelay = $nowDelay;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{timezone?: string, from?: string, to?: string, autoRefresh?: string, autoRefreshIntervals?: array<string>, quickRanges?: array<mixed>, hideTimepicker?: bool, weekStart?: string, fiscalYearStartMonth?: int, nowDelay?: string} $inputData */
        $data = $inputData;
        return new self(
            timezone: $data["timezone"] ?? null,
            from: $data["from"] ?? null,
            to: $data["to"] ?? null,
            autoRefresh: $data["autoRefresh"] ?? null,
            autoRefreshIntervals: $data["autoRefreshIntervals"] ?? null,
            quickRanges: array_filter(array_map((function($input) {
    	/** @var array{display?: string, from?: string, to?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\TimeRangeOption::fromArray($val);
    }), $data["quickRanges"] ?? [])),
            hideTimepicker: $data["hideTimepicker"] ?? null,
            weekStart: isset($data["weekStart"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpecWeekStart::fromValue($input); })($data["weekStart"]) : null,
            fiscalYearStartMonth: $data["fiscalYearStartMonth"] ?? null,
            nowDelay: $data["nowDelay"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->from = $this->from;
        $data->to = $this->to;
        $data->autoRefresh = $this->autoRefresh;
        $data->autoRefreshIntervals = $this->autoRefreshIntervals;
        $data->hideTimepicker = $this->hideTimepicker;
        $data->fiscalYearStartMonth = $this->fiscalYearStartMonth;
        if (isset($this->timezone)) {
            $data->timezone = $this->timezone;
        }
        if (isset($this->quickRanges)) {
            $data->quickRanges = $this->quickRanges;
        }
        if (isset($this->weekStart)) {
            $data->weekStart = $this->weekStart;
        }
        if (isset($this->nowDelay)) {
            $data->nowDelay = $this->nowDelay;
        }
        return $data;
    }
}
