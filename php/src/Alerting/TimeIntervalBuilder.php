<?php

namespace Grafana\Foundation\Alerting;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\TimeInterval>
 */
class TimeIntervalBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\TimeInterval $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alerting\TimeInterval();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alerting\TimeInterval
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\TimeRange>> $times
     */
    public function times(array $times): static
    {
            $timesResources = [];
            foreach ($times as $r1) {
                    $timesResources[] = $r1->build();
            }
        $this->internal->times = $timesResources;
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\WeekdayRange>> $weekdays
     */
    public function weekdays(array $weekdays): static
    {
            $weekdaysResources = [];
            foreach ($weekdays as $r1) {
                    $weekdaysResources[] = $r1->build();
            }
        $this->internal->weekdays = $weekdaysResources;
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\DayOfMonthRange>> $daysOfMonth
     */
    public function daysOfMonth(array $daysOfMonth): static
    {
            $daysOfMonthResources = [];
            foreach ($daysOfMonth as $r1) {
                    $daysOfMonthResources[] = $r1->build();
            }
        $this->internal->daysOfMonth = $daysOfMonthResources;
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\MonthRange>> $months
     */
    public function months(array $months): static
    {
            $monthsResources = [];
            foreach ($months as $r1) {
                    $monthsResources[] = $r1->build();
            }
        $this->internal->months = $monthsResources;
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\YearRange>> $years
     */
    public function years(array $years): static
    {
            $yearsResources = [];
            foreach ($years as $r1) {
                    $yearsResources[] = $r1->build();
            }
        $this->internal->years = $yearsResources;
    
        return $this;
    }
    public function location(string $location): static
    {
        $this->internal->location = $location;
    
        return $this;
    }

}
