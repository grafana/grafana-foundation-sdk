<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource>
 */
class Dashboardv2beta1DataQueryKindDatasourceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource
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
