<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TimeRangeOption>
 */
class TimeRangeOptionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\TimeRangeOption $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\TimeRangeOption();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\TimeRangeOption
     */
    public function build()
    {
        return $this->internal;
    }

    public function display(string $display): static
    {
        $this->internal->display = $display;
    
        return $this;
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
