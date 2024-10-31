<?php

namespace Grafana\Foundation\Alerting;

/**
 * TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
 * within the interval.
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
     * @param array<string> $daysOfMonth
     */
    public function daysOfMonth(array $daysOfMonth): static
    {
        $this->internal->daysOfMonth = $daysOfMonth;
    
        return $this;
    }
    public function location(string $location): static
    {
        $this->internal->location = $location;
    
        return $this;
    }
    /**
     * @param array<string> $months
     */
    public function months(array $months): static
    {
        $this->internal->months = $months;
    
        return $this;
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
     * @param array<string> $weekdays
     */
    public function weekdays(array $weekdays): static
    {
        $this->internal->weekdays = $weekdays;
    
        return $this;
    }
    /**
     * @param array<string> $years
     */
    public function years(array $years): static
    {
        $this->internal->years = $years;
    
        return $this;
    }

}
