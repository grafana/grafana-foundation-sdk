<?php

namespace Grafana\Foundation\Common;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HeatmapCalculationOptions>
 */
class HeatmapCalculationOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\HeatmapCalculationOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\HeatmapCalculationOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\HeatmapCalculationOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The number of buckets to use for the xAxis in the heatmap
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HeatmapCalculationBucketConfig> $xBuckets
     */
    public function xBuckets(\Grafana\Foundation\Cog\Builder $xBuckets): static
    {
        $xBucketsResource = $xBuckets->build();
        $this->internal->xBuckets = $xBucketsResource;
    
        return $this;
    }
    /**
     * The number of buckets to use for the yAxis in the heatmap
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HeatmapCalculationBucketConfig> $yBuckets
     */
    public function yBuckets(\Grafana\Foundation\Cog\Builder $yBuckets): static
    {
        $yBucketsResource = $yBuckets->build();
        $this->internal->yBuckets = $yBucketsResource;
    
        return $this;
    }

}
