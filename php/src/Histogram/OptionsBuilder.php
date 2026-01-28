<?php

namespace Grafana\Foundation\Histogram;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Histogram\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Histogram\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Histogram\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Histogram\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Bucket count (approx)
     */
    public function bucketCount(int $bucketCount): static
    {
        if (!($bucketCount > 0)) {
            throw new \ValueError('$bucketCount must be > 0');
        }
        $this->internal->bucketCount = $bucketCount;
    
        return $this;
    }

    /**
     * Size of each bucket
     */
    public function bucketSize(int $bucketSize): static
    {
        $this->internal->bucketSize = $bucketSize;
    
        return $this;
    }

    /**
     * Offset buckets by this amount
     */
    public function bucketOffset(float $bucketOffset): static
    {
        $this->internal->bucketOffset = $bucketOffset;
    
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
     * Combines multiple series into a single histogram
     */
    public function combine(bool $combine): static
    {
        $this->internal->combine = $combine;
    
        return $this;
    }

}
