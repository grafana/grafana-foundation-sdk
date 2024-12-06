<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardRangeMapOptions>
 */
class DashboardRangeMapOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\DashboardRangeMapOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\DashboardRangeMapOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\DashboardRangeMapOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Min value of the range. It can be null which means -Infinity
     */
    public function from(float $from): static
    {
        $this->internal->from = $from;
    
        return $this;
    }
    /**
     * Max value of the range. It can be null which means +Infinity
     */
    public function to(float $to): static
    {
        $this->internal->to = $to;
    
        return $this;
    }
    /**
     * Config to apply when the value is within the range
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\ValueMappingResult> $result
     */
    public function result(\Grafana\Foundation\Cog\Builder $result): static
    {
        $resultResource = $result->build();
        $this->internal->result = $resultResource;
    
        return $this;
    }

}
