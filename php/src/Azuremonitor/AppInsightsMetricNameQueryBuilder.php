<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery>
 */
class AppInsightsMetricNameQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery();
    $this->internal->kind = "AppInsightsMetricNameQuery";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery
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
