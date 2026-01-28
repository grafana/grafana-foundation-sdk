<?php

namespace Grafana\Foundation\Statetimeline;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Statetimeline\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Statetimeline\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Statetimeline\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Statetimeline\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Show timeline values on chart
     */
    public function showValue(\Grafana\Foundation\Common\VisibilityMode $showValue): static
    {
        $this->internal->showValue = $showValue;
    
        return $this;
    }

    /**
     * Controls the row height
     */
    public function rowHeight(float $rowHeight): static
    {
        if (!($rowHeight <= 1)) {
            throw new \ValueError('$rowHeight must be <= 1');
        }
        $this->internal->rowHeight = $rowHeight;
    
        return $this;
    }

    /**
     * Merge equal consecutive values
     */
    public function mergeValues(bool $mergeValues): static
    {
        $this->internal->mergeValues = $mergeValues;
    
        return $this;
    }

    /**
     * Controls value alignment on the timelines
     */
    public function alignValue(\Grafana\Foundation\Common\TimelineValueAlignment $alignValue): static
    {
        $this->internal->alignValue = $alignValue;
    
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
