<?php

namespace Grafana\Foundation\Common;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HeatmapCalculationBucketConfig>
 */
class HeatmapCalculationBucketConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\HeatmapCalculationBucketConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\HeatmapCalculationBucketConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\HeatmapCalculationBucketConfig
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Sets the bucket calculation mode
     */
    public function mode(\Grafana\Foundation\Common\HeatmapCalculationMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }
    /**
     * The number of buckets to use for the axis in the heatmap
     */
    public function value(string $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }
    /**
     * Controls the scale of the buckets
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDistributionConfig> $scale
     */
    public function scale(\Grafana\Foundation\Cog\Builder $scale): static
    {
        $scaleResource = $scale->build();
        $this->internal->scale = $scaleResource;
    
        return $this;
    }

}
