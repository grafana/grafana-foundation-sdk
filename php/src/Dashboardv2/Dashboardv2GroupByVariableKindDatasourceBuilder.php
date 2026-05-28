<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\Dashboardv2GroupByVariableKindDatasource>
 */
class Dashboardv2GroupByVariableKindDatasourceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\Dashboardv2GroupByVariableKindDatasource $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\Dashboardv2GroupByVariableKindDatasource();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\Dashboardv2GroupByVariableKindDatasource
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
