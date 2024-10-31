<?php

namespace Grafana\Foundation\Geomap;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Geomap\ControlsOptions>
 */
class ControlsOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Geomap\ControlsOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Geomap\ControlsOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Geomap\ControlsOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Zoom (upper left)
     */
    public function showZoom(bool $showZoom): static
    {
        $this->internal->showZoom = $showZoom;
    
        return $this;
    }
    /**
     * let the mouse wheel zoom
     */
    public function mouseWheelZoom(bool $mouseWheelZoom): static
    {
        $this->internal->mouseWheelZoom = $mouseWheelZoom;
    
        return $this;
    }
    /**
     * Lower right
     */
    public function showAttribution(bool $showAttribution): static
    {
        $this->internal->showAttribution = $showAttribution;
    
        return $this;
    }
    /**
     * Scale options
     */
    public function showScale(bool $showScale): static
    {
        $this->internal->showScale = $showScale;
    
        return $this;
    }
    /**
     * Show debug
     */
    public function showDebug(bool $showDebug): static
    {
        $this->internal->showDebug = $showDebug;
    
        return $this;
    }
    /**
     * Show measure
     */
    public function showMeasure(bool $showMeasure): static
    {
        $this->internal->showMeasure = $showMeasure;
    
        return $this;
    }

}
