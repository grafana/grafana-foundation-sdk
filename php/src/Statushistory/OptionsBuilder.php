<?php

namespace Grafana\Foundation\Statushistory;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Statushistory\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Statushistory\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Statushistory\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Statushistory\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Set the height of the rows
     */
    public function rowHeight(float $rowHeight): static
    {
        if (!($rowHeight >= 0)) {
            throw new \ValueError('$rowHeight must be >= 0');
        }
        if (!($rowHeight <= 1)) {
            throw new \ValueError('$rowHeight must be <= 1');
        }
        $this->internal->rowHeight = $rowHeight;
    
        return $this;
    }

    /**
     * Show values on the columns
     */
    public function showValue(\Grafana\Foundation\Common\VisibilityMode $showValue): static
    {
        $this->internal->showValue = $showValue;
    
        return $this;
    }

    /**
     * Controls the column width
     */
    public function colWidth(float $colWidth): static
    {
        if (!($colWidth <= 1)) {
            throw new \ValueError('$colWidth must be <= 1');
        }
        $this->internal->colWidth = $colWidth;
    
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
     * @param array<string> $timezone
     */
    public function timezone(array $timezone): static
    {
        $this->internal->timezone = $timezone;
    
        return $this;
    }

    /**
     * Enables pagination when > 0
     */
    public function perPage(float $perPage): static
    {
        if (!($perPage >= 1)) {
            throw new \ValueError('$perPage must be >= 1');
        }
        $this->internal->perPage = $perPage;
    
        return $this;
    }

}
