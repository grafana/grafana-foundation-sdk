<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\BaseGrafanaTemplateVariableQuery>
 */
class BaseGrafanaTemplateVariableQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\BaseGrafanaTemplateVariableQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\BaseGrafanaTemplateVariableQuery();
    }

    /**
     * @return \Grafana\Foundation\Azuremonitor\BaseGrafanaTemplateVariableQuery
     */
    public function build()
    {
        return $this->internal;
    }

    public function rawQuery(string $rawQuery): static
    {
        $this->internal->rawQuery = $rawQuery;
    
        return $this;
    }

}
