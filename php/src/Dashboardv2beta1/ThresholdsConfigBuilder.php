<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ThresholdsConfig>
 */
class ThresholdsConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\ThresholdsConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\ThresholdsConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\ThresholdsConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(\Grafana\Foundation\Dashboardv2beta1\ThresholdsMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Dashboardv2beta1\Threshold> $steps
     */
    public function steps(array $steps): static
    {
        $this->internal->steps = $steps;
    
        return $this;
    }

}
