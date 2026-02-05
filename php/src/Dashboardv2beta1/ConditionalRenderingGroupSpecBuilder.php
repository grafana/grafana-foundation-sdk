<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpec>
 */
class ConditionalRenderingGroupSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpec
     */
    public function build()
    {
        return $this->internal;
    }

    public function visibility(\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecVisibility $visibility): static
    {
        $this->internal->visibility = $visibility;
    
        return $this;
    }

    public function condition(\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecCondition $condition): static
    {
        $this->internal->condition = $condition;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingTimeRangeSizeKind>> $items
     */
    public function items(array $items): static
    {
            $itemsResources = [];
            foreach ($items as $r1) {
                    $itemsResources[] = $r1->build();
            }
        $this->internal->items = $itemsResources;
    
        return $this;
    }

}
