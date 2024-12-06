<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\FillConfig>
 */
class FillConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\FillConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\FillConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\FillConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function fillColor(string $fillColor): static
    {
        $this->internal->fillColor = $fillColor;
    
        return $this;
    }
    public function fillOpacity(float $fillOpacity): static
    {
        $this->internal->fillOpacity = $fillOpacity;
    
        return $this;
    }
    public function fillBelowTo(string $fillBelowTo): static
    {
        $this->internal->fillBelowTo = $fillBelowTo;
    
        return $this;
    }

}
