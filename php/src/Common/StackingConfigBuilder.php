<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\StackingConfig>
 */
class StackingConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\StackingConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\StackingConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\StackingConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(\Grafana\Foundation\Common\StackingMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }
    public function group(string $group): static
    {
        $this->internal->group = $group;
    
        return $this;
    }

}
