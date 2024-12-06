<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDistributionConfig>
 */
class ScaleDistributionConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\ScaleDistributionConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\ScaleDistributionConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\ScaleDistributionConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function type(\Grafana\Foundation\Common\ScaleDistribution $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    public function log(float $log): static
    {
        $this->internal->log = $log;
    
        return $this;
    }
    public function linearThreshold(float $linearThreshold): static
    {
        $this->internal->linearThreshold = $linearThreshold;
    
        return $this;
    }

}
