<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardDashboardTemplating>
 */
class DashboardDashboardTemplatingBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\DashboardDashboardTemplating $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\DashboardDashboardTemplating();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\DashboardDashboardTemplating
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * List of configured template variables with their saved values along with some other metadata
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\VariableModel>> $list
     */
    public function list(array $list): static
    {
            $listResources = [];
            foreach ($list as $r1) {
                    $listResources[] = $r1->build();
            }
        $this->internal->list = $listResources;
    
        return $this;
    }

}
