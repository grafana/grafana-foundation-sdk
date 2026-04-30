<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ThresholdsConfig>
 */
class ThresholdsConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\ThresholdsConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\ThresholdsConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\ThresholdsConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(\Grafana\Foundation\Dashboardv2\ThresholdsMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Dashboardv2\Threshold> $steps
     */
    public function steps(array $steps): static
    {
        $this->internal->steps = $steps;
    
        return $this;
    }

}
