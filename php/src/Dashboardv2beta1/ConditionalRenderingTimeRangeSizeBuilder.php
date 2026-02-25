<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingTimeRangeSizeKind>
 */
class ConditionalRenderingTimeRangeSizeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingTimeRangeSizeKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingTimeRangeSizeKind();
    $this->internal->kind = "ConditionalRenderingTimeRangeSize";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingTimeRangeSizeKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function value(string $value): static
    {
        $this->internal->spec->value = $value;
    
        return $this;
    }

}
