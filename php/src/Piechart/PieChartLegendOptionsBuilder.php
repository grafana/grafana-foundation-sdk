<?php

namespace Grafana\Foundation\Piechart;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Piechart\PieChartLegendOptions>
 */
class PieChartLegendOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Piechart\PieChartLegendOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Piechart\PieChartLegendOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Piechart\PieChartLegendOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Piechart\PieChartLegendValues> $values
     */
    public function values(array $values): static
    {
        $this->internal->values = $values;
    
        return $this;
    }
    public function displayMode(\Grafana\Foundation\Common\LegendDisplayMode $displayMode): static
    {
        $this->internal->displayMode = $displayMode;
    
        return $this;
    }
    public function placement(\Grafana\Foundation\Common\LegendPlacement $placement): static
    {
        $this->internal->placement = $placement;
    
        return $this;
    }
    public function showLegend(bool $showLegend): static
    {
        $this->internal->showLegend = $showLegend;
    
        return $this;
    }
    public function asTable(bool $asTable): static
    {
        $this->internal->asTable = $asTable;
    
        return $this;
    }
    public function isVisible(bool $isVisible): static
    {
        $this->internal->isVisible = $isVisible;
    
        return $this;
    }
    public function sortBy(string $sortBy): static
    {
        $this->internal->sortBy = $sortBy;
    
        return $this;
    }
    public function sortDesc(bool $sortDesc): static
    {
        $this->internal->sortDesc = $sortDesc;
    
        return $this;
    }
    public function width(float $width): static
    {
        $this->internal->width = $width;
    
        return $this;
    }
    /**
     * @param array<string> $calcs
     */
    public function calcs(array $calcs): static
    {
        $this->internal->calcs = $calcs;
    
        return $this;
    }

}
