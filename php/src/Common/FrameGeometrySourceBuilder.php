<?php

namespace Grafana\Foundation\Common;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\FrameGeometrySource>
 */
class FrameGeometrySourceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\FrameGeometrySource $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\FrameGeometrySource();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\FrameGeometrySource
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(\Grafana\Foundation\Common\FrameGeometrySourceMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }
    /**
     * Field mappings
     */
    public function geohash(string $geohash): static
    {
        $this->internal->geohash = $geohash;
    
        return $this;
    }
    public function latitude(string $latitude): static
    {
        $this->internal->latitude = $latitude;
    
        return $this;
    }
    public function longitude(string $longitude): static
    {
        $this->internal->longitude = $longitude;
    
        return $this;
    }
    public function wkt(string $wkt): static
    {
        $this->internal->wkt = $wkt;
    
        return $this;
    }
    public function lookup(string $lookup): static
    {
        $this->internal->lookup = $lookup;
    
        return $this;
    }
    /**
     * Path to Gazetteer
     */
    public function gazetteer(string $gazetteer): static
    {
        $this->internal->gazetteer = $gazetteer;
    
        return $this;
    }

}
