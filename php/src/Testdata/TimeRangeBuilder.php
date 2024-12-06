<?php

namespace Grafana\Foundation\Testdata;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\TimeRange>
 */
class TimeRangeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Testdata\TimeRange $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Testdata\TimeRange();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Testdata\TimeRange
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
