<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\BarConfig>
 */
class BarConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\BarConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\BarConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\BarConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function barAlignment(\Grafana\Foundation\Common\BarAlignment $barAlignment): static
    {
        $this->internal->barAlignment = $barAlignment;
    
        return $this;
    }
    public function barWidthFactor(float $barWidthFactor): static
    {
        $this->internal->barWidthFactor = $barWidthFactor;
    
        return $this;
    }
    public function barMaxWidth(float $barMaxWidth): static
    {
        $this->internal->barMaxWidth = $barMaxWidth;
    
        return $this;
    }

}
