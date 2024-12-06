<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Thresholds configuration for the panel
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\ThresholdsConfig>
 */
class ThresholdsConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\ThresholdsConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\ThresholdsConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\ThresholdsConfig
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Thresholds mode.
     */
    public function mode(\Grafana\Foundation\Dashboard\ThresholdsMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }
    /**
     * Must be sorted by 'value', first value is always -Infinity
     * @param array<\Grafana\Foundation\Dashboard\Threshold> $steps
     */
    public function steps(array $steps): static
    {
        $this->internal->steps = $steps;
    
        return $this;
    }

}
