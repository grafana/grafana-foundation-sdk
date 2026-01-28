<?php

namespace Grafana\Foundation\Geomap;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Geomap\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Geomap\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Geomap\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Geomap\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Geomap\MapViewConfig> $view
     */
    public function view(\Grafana\Foundation\Cog\Builder $view): static
    {
        $viewResource = $view->build();
        $this->internal->view = $viewResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Geomap\ControlsOptions> $controls
     */
    public function controls(\Grafana\Foundation\Cog\Builder $controls): static
    {
        $controlsResource = $controls->build();
        $this->internal->controls = $controlsResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\MapLayerOptions> $basemap
     */
    public function basemap(\Grafana\Foundation\Cog\Builder $basemap): static
    {
        $basemapResource = $basemap->build();
        $this->internal->basemap = $basemapResource;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\MapLayerOptions>> $layers
     */
    public function layers(array $layers): static
    {
            $layersResources = [];
            foreach ($layers as $r1) {
                    $layersResources[] = $r1->build();
            }
        $this->internal->layers = $layersResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Geomap\TooltipOptions> $tooltip
     */
    public function tooltip(\Grafana\Foundation\Cog\Builder $tooltip): static
    {
        $tooltipResource = $tooltip->build();
        $this->internal->tooltip = $tooltipResource;
    
        return $this;
    }

}
