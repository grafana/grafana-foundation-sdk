<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls exemplar options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Heatmap\ExemplarConfig>
 */
class ExemplarConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Heatmap\ExemplarConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Heatmap\ExemplarConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Heatmap\ExemplarConfig
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Sets the color of the exemplar markers
     */
    public function color(string $color): static
    {
        $this->internal->color = $color;
    
        return $this;
    }

}
