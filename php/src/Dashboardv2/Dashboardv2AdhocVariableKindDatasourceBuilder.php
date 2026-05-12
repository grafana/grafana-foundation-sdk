<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\Dashboardv2AdhocVariableKindDatasource>
 */
class Dashboardv2AdhocVariableKindDatasourceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\Dashboardv2AdhocVariableKindDatasource $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\Dashboardv2AdhocVariableKindDatasource();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\Dashboardv2AdhocVariableKindDatasource
     */
    public function build()
    {
        return $this->internal;
    }

    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }

}
