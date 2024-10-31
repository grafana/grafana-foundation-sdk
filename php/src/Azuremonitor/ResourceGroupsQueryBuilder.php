<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\ResourceGroupsQuery>
 */
class ResourceGroupsQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\ResourceGroupsQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\ResourceGroupsQuery();
    $this->internal->kind = "ResourceGroupsQuery";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\ResourceGroupsQuery
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
    public function subscription(string $subscription): static
    {
        $this->internal->subscription = $subscription;
    
        return $this;
    }

}
