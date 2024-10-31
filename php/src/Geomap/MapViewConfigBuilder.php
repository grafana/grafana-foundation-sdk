<?php

namespace Grafana\Foundation\Geomap;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Geomap\MapViewConfig>
 */
class MapViewConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Geomap\MapViewConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Geomap\MapViewConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Geomap\MapViewConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function id(string $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    public function lat(int $lat): static
    {
        $this->internal->lat = $lat;
    
        return $this;
    }
    public function lon(int $lon): static
    {
        $this->internal->lon = $lon;
    
        return $this;
    }
    public function zoom(int $zoom): static
    {
        $this->internal->zoom = $zoom;
    
        return $this;
    }
    public function minZoom(int $minZoom): static
    {
        $this->internal->minZoom = $minZoom;
    
        return $this;
    }
    public function maxZoom(int $maxZoom): static
    {
        $this->internal->maxZoom = $maxZoom;
    
        return $this;
    }
    public function padding(int $padding): static
    {
        $this->internal->padding = $padding;
    
        return $this;
    }
    public function allLayers(bool $allLayers): static
    {
        $this->internal->allLayers = $allLayers;
    
        return $this;
    }
    public function lastOnly(bool $lastOnly): static
    {
        $this->internal->lastOnly = $lastOnly;
    
        return $this;
    }
    public function layer(string $layer): static
    {
        $this->internal->layer = $layer;
    
        return $this;
    }
    public function shared(bool $shared): static
    {
        $this->internal->shared = $shared;
    
        return $this;
    }

}
