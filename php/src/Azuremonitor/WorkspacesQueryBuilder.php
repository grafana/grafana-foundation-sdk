<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\WorkspacesQuery>
 */
class WorkspacesQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\WorkspacesQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\WorkspacesQuery();
    $this->internal->kind = "WorkspacesQuery";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\WorkspacesQuery
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
