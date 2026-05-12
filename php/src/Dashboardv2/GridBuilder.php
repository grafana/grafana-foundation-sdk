<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\GridLayoutKind>
 */
class GridBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\GridLayoutKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\GridLayoutKind();
    $this->internal->kind = "GridLayout";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\GridLayoutKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\GridLayoutItemKind>> $items
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
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\GridLayoutItemKind> $item
     */
    public function item(\Grafana\Foundation\Cog\Builder $item): static
    {
        $itemResource = $item->build();
        $this->internal->spec->items[] = $itemResource;
    
        return $this;
    }

}
