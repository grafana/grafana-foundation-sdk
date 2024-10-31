<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HideSeriesConfig>
 */
class HideSeriesConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\HideSeriesConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\HideSeriesConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\HideSeriesConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function tooltip(bool $tooltip): static
    {
        $this->internal->tooltip = $tooltip;
    
        return $this;
    }
    public function legend(bool $legend): static
    {
        $this->internal->legend = $legend;
    
        return $this;
    }
    public function viz(bool $viz): static
    {
        $this->internal->viz = $viz;
    
        return $this;
    }

}
