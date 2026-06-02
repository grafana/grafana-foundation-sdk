<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\ResourceGraphQuery>
 */
class ResourceGraphQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\ResourceGraphQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\ResourceGraphQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\ResourceGraphQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Azure Resource Graph KQL query to be executed.
     */
    public function query(string $query): static
    {
        $this->internal->query = $query;
    
        return $this;
    }

    /**
     * Specifies the format results should be returned as. Defaults to table.
     */
    public function resultFormat(string $resultFormat): static
    {
        $this->internal->resultFormat = $resultFormat;
    
        return $this;
    }

}
