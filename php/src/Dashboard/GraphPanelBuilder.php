<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Support for legacy graph panel.
 * @deprecated this a deprecated panel type
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\GraphPanel>
 */
class GraphPanelBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\GraphPanel $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\GraphPanel();
    $this->internal->type = "graph";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\GraphPanel
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @deprecated this is part of deprecated graph panel
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardGraphPanelLegend> $legend
     */
    public function legend(\Grafana\Foundation\Cog\Builder $legend): static
    {
        $legendResource = $legend->build();
        $this->internal->legend = $legendResource;
    
        return $this;
    }

}
