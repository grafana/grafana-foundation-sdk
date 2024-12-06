<?php

namespace Grafana\Foundation\Common;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\MapLayerOptions>
 */
class MapLayerOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\MapLayerOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\MapLayerOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\MapLayerOptions
     */
    public function build()
    {
        return $this->internal;
    }

    public function type(string $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    /**
     * configured unique display name
     */
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    /**
     * Custom options depending on the type
     * @param mixed $config
     */
    public function config( $config): static
    {
        $this->internal->config = $config;
    
        return $this;
    }
    /**
     * Common method to define geometry fields
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\FrameGeometrySource> $location
     */
    public function location(\Grafana\Foundation\Cog\Builder $location): static
    {
        $locationResource = $location->build();
        $this->internal->location = $locationResource;
    
        return $this;
    }
    /**
     * Defines a frame MatcherConfig that may filter data for the given layer
     * @param mixed $filterData
     */
    public function filterData( $filterData): static
    {
        $this->internal->filterData = $filterData;
    
        return $this;
    }
    /**
     * Common properties:
     * https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
     * Layer opacity (0-1)
     */
    public function opacity(int $opacity): static
    {
        $this->internal->opacity = $opacity;
    
        return $this;
    }
    /**
     * Check tooltip (defaults to true)
     */
    public function tooltip(bool $tooltip): static
    {
        $this->internal->tooltip = $tooltip;
    
        return $this;
    }

}
