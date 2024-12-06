<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\StackableFieldConfig>
 */
class StackableFieldConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\StackableFieldConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\StackableFieldConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\StackableFieldConfig
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\StackingConfig> $stacking
     */
    public function stacking(\Grafana\Foundation\Cog\Builder $stacking): static
    {
        $stackingResource = $stacking->build();
        $this->internal->stacking = $stackingResource;
    
        return $this;
    }

}
