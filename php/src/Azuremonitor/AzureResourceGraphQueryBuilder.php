<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureResourceGraphQuery>
 */
class AzureResourceGraphQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\AzureResourceGraphQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\AzureResourceGraphQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\AzureResourceGraphQuery
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
