<?php

namespace Grafana\Foundation\Piechart;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Piechart\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Piechart\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Piechart\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Piechart\Options
     */
    public function build()
    {
        return $this->internal;
    }

    public function pieType(\Grafana\Foundation\Piechart\PieChartType $pieType): static
    {
        $this->internal->pieType = $pieType;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Piechart\PieChartLabels> $displayLabels
     */
    public function displayLabels(array $displayLabels): static
    {
        $this->internal->displayLabels = $displayLabels;
    
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
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ReduceDataOptions> $reduceOptions
     */
    public function reduceOptions(\Grafana\Foundation\Cog\Builder $reduceOptions): static
    {
        $reduceOptionsResource = $reduceOptions->build();
        $this->internal->reduceOptions = $reduceOptionsResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTextDisplayOptions> $text
     */
    public function text(\Grafana\Foundation\Cog\Builder $text): static
    {
        $textResource = $text->build();
        $this->internal->text = $textResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Piechart\PieChartLegendOptions> $legend
     */
    public function legend(\Grafana\Foundation\Cog\Builder $legend): static
    {
        $legendResource = $legend->build();
        $this->internal->legend = $legendResource;
    
        return $this;
    }

    public function orientation(\Grafana\Foundation\Common\VizOrientation $orientation): static
    {
        $this->internal->orientation = $orientation;
    
        return $this;
    }

}
