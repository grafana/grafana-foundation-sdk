<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardDashboardTime>
 */
class DashboardDashboardTimeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\DashboardDashboardTime $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\DashboardDashboardTime();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\DashboardDashboardTime
     */
    public function build()
    {
        return $this->internal;
    }

    public function from(string $from): static
    {
        $this->internal->from = $from;
    
        return $this;
    }
    public function to(string $to): static
    {
        $this->internal->to = $to;
    
        return $this;
    }

}
