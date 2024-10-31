<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * Time Series sub-query properties.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQuery>
 */
class TimeSeriesQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * GCP project to execute the query against.
     */
    public function projectName(string $projectName): static
    {
        $this->internal->projectName = $projectName;
    
        return $this;
    }
    /**
     * MQL query to be executed.
     */
    public function query(string $query): static
    {
        $this->internal->query = $query;
    
        return $this;
    }
    /**
     * To disable the graphPeriod, it should explictly be set to 'disabled'.
     */
    public function graphPeriod(string $graphPeriod): static
    {
        $this->internal->graphPeriod = $graphPeriod;
    
        return $this;
    }

}
