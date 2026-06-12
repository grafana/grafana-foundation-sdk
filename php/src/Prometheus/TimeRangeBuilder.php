<?php

namespace Grafana\Foundation\Prometheus;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\TimeRange>
 */
class TimeRangeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Prometheus\TimeRange $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Prometheus\TimeRange();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Prometheus\TimeRange
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * From is the start time of the query.
     */
    public function from(string $from): static
    {
        $this->internal->from = $from;
    
        return $this;
    }

    /**
     * To is the end time of the query.
     */
    public function to(string $to): static
    {
        $this->internal->to = $to;
    
        return $this;
    }

}
