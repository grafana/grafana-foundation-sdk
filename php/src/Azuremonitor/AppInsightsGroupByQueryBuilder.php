<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AppInsightsGroupByQuery>
 */
class AppInsightsGroupByQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\AppInsightsGroupByQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\AppInsightsGroupByQuery();
    $this->internal->kind = "AppInsightsGroupByQuery";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\AppInsightsGroupByQuery
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
    public function metricName(string $metricName): static
    {
        $this->internal->metricName = $metricName;
    
        return $this;
    }

}
