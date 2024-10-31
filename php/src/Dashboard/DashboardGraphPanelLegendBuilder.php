<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardGraphPanelLegend>
 */
class DashboardGraphPanelLegendBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\DashboardGraphPanelLegend $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\DashboardGraphPanelLegend();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\DashboardGraphPanelLegend
     */
    public function build()
    {
        return $this->internal;
    }

    public function show(bool $show): static
    {
        $this->internal->show = $show;
    
        return $this;
    }
    public function sort(string $sort): static
    {
        $this->internal->sort = $sort;
    
        return $this;
    }
    public function sortDesc(bool $sortDesc): static
    {
        $this->internal->sortDesc = $sortDesc;
    
        return $this;
    }

}
