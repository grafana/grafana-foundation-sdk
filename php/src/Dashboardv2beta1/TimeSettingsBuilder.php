<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Time configuration
 * It defines the default time config for the time picker, the refresh picker for the specific dashboard.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec>
 */
class TimeSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
     */
    public function timezone(string $timezone): static
    {
        $this->internal->timezone = $timezone;
    
        return $this;
    }

    /**
     * Start time range for dashboard.
     * Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
     */
    public function from(string $from): static
    {
        $this->internal->from = $from;
    
        return $this;
    }

    /**
     * End time range for dashboard.
     * Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
     */
    public function to(string $to): static
    {
        $this->internal->to = $to;
    
        return $this;
    }

    /**
     * Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
     * v1: refresh
     */
    public function autoRefresh(string $autoRefresh): static
    {
        $this->internal->autoRefresh = $autoRefresh;
    
        return $this;
    }

    /**
     * Interval options available in the refresh picker dropdown.
     * v1: timepicker.refresh_intervals
     * @param array<string> $autoRefreshIntervals
     */
    public function autoRefreshIntervals(array $autoRefreshIntervals): static
    {
        $this->internal->autoRefreshIntervals = $autoRefreshIntervals;
    
        return $this;
    }

    /**
     * Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
     * v1: timepicker.quick_ranges , not exposed in the UI
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TimeRangeOption>> $quickRanges
     */
    public function quickRanges(array $quickRanges): static
    {
            $quickRangesResources = [];
            foreach ($quickRanges as $r1) {
                    $quickRangesResources[] = $r1->build();
            }
        $this->internal->quickRanges = $quickRangesResources;
    
        return $this;
    }

    /**
     * Whether timepicker is visible or not.
     * v1: timepicker.hidden
     */
    public function hideTimepicker(bool $hideTimepicker): static
    {
        $this->internal->hideTimepicker = $hideTimepicker;
    
        return $this;
    }

    /**
     * Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
     */
    public function weekStart(\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpecWeekStart $weekStart): static
    {
        $this->internal->weekStart = $weekStart;
    
        return $this;
    }

    /**
     * The month that the fiscal year starts on. 0 = January, 11 = December
     */
    public function fiscalYearStartMonth(int $fiscalYearStartMonth): static
    {
        $this->internal->fiscalYearStartMonth = $fiscalYearStartMonth;
    
        return $this;
    }

    /**
     * Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
     * v1: timepicker.nowDelay
     */
    public function nowDelay(string $nowDelay): static
    {
        $this->internal->nowDelay = $nowDelay;
    
        return $this;
    }

}
