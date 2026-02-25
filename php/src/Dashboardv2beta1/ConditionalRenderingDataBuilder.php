<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataKind>
 */
class ConditionalRenderingDataBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataKind();
    $this->internal->kind = "ConditionalRenderingData";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function value(bool $value): static
    {
        $this->internal->spec->value = $value;
    
        return $this;
    }

}
