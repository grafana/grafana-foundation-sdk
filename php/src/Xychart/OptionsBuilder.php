<?php

namespace Grafana\Foundation\Xychart;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Xychart\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Xychart\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Xychart\Options
     */
    public function build()
    {
        return $this->internal;
    }

    public function mapping(\Grafana\Foundation\Xychart\SeriesMapping $mapping): static
    {
        $this->internal->mapping = $mapping;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend
     */
    public function legend(\Grafana\Foundation\Cog\Builder $legend): static
    {
        $legendResource = $legend->build();
        $this->internal->legend = $legendResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTooltipOptions> $tooltip
     */
    public function tooltip(\Grafana\Foundation\Cog\Builder $tooltip): static
    {
        $tooltipResource = $tooltip->build();
        $this->internal->tooltip = $tooltipResource;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XYSeriesConfig>> $series
     */
    public function series(array $series): static
    {
            $seriesResources = [];
            foreach ($series as $r1) {
                    $seriesResources[] = $r1->build();
            }
        $this->internal->series = $seriesResources;
    
        return $this;
    }

}
