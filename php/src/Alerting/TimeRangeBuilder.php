<?php

namespace Grafana\Foundation\Alerting;

/**
 * Redefining this to avoid an import cycle
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\TimeRange>
 */
class TimeRangeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\TimeRange $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alerting\TimeRange();
    }

    /**
     * @return \Grafana\Foundation\Alerting\TimeRange
     */
    public function build()
    {
        return $this->internal;
    }

    public function from(string $from): static
    {
        $this->internal->from = $from;
    
        return $this;
    }
    public function to(string $to): static
    {
        $this->internal->to = $to;
    
        return $this;
    }

}
