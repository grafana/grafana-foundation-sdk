<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind>
 */
class ConditionalRenderingGroupBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind();
    $this->internal->kind = "ConditionalRenderingGroup";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function visibility(\Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupSpecVisibility $visibility): static
    {
        $this->internal->spec->visibility = $visibility;
    
        return $this;
    }

    public function condition(\Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupSpecCondition $condition): static
    {
        $this->internal->spec->condition = $condition;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ConditionalRenderingVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ConditionalRenderingDataKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ConditionalRenderingTimeRangeSizeKind>> $items
     */
    public function items(array $items): static
    {
            $itemsResources = [];
            foreach ($items as $r1) {
                    $itemsResources[] = $r1->build();
            }
        $this->internal->spec->items = $itemsResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ConditionalRenderingVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ConditionalRenderingDataKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ConditionalRenderingTimeRangeSizeKind> $item
     */
    public function item( $item): static
    {
        $itemResource = $item->build();
        $this->internal->spec->items[] = $itemResource;
    
        return $this;
    }

}
