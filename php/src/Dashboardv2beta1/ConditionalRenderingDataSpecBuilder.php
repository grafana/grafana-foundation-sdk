<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataSpec>
 */
class ConditionalRenderingDataSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataSpec
     */
    public function build()
    {
        return $this->internal;
    }

    public function value(bool $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }

}
