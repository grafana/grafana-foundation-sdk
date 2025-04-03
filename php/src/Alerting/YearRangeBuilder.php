<?php

namespace Grafana\Foundation\Alerting;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\YearRange>
 */
class YearRangeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\YearRange $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alerting\YearRange();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alerting\YearRange
     */
    public function build()
    {
        return $this->internal;
    }

    public function begin(int $begin): static
    {
        $this->internal->begin = $begin;
    
        return $this;
    }

    public function end(int $end): static
    {
        $this->internal->end = $end;
    
        return $this;
    }

}
