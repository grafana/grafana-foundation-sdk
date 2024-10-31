<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\SubscriptionsQuery>
 */
class SubscriptionsQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\SubscriptionsQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\SubscriptionsQuery();
    $this->internal->kind = "SubscriptionsQuery";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\SubscriptionsQuery
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
