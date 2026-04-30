<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ControlSourceRef>
 */
class ControlSourceRefBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\ControlSourceRef $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\ControlSourceRef();
    $this->internal->type = "datasource";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\ControlSourceRef
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The plugin type-id
     */
    public function group(string $group): static
    {
        $this->internal->group = $group;
    
        return $this;
    }

}
