<?php

namespace Grafana\Foundation\Alerting;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\TimeIntervalTimeRange>
 */
class TimeIntervalTimeRangeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\TimeIntervalTimeRange $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alerting\TimeIntervalTimeRange();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alerting\TimeIntervalTimeRange
     */
    public function build()
    {
        return $this->internal;
    }

    public function endTime(string $endTime): static
    {
        $this->internal->endTime = $endTime;
    
        return $this;
    }
    public function startTime(string $startTime): static
    {
        $this->internal->startTime = $startTime;
    
        return $this;
    }

}
