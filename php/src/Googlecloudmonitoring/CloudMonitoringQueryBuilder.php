<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery>
 */
class CloudMonitoringQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public function refId(string $refId): static
    {
        $this->internal->refId = $refId;
    
        return $this;
    }
    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }
    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public function queryType(string $queryType): static
    {
        $this->internal->queryType = $queryType;
    
        return $this;
    }
    /**
     * Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
     */
    public function aliasBy(string $aliasBy): static
    {
        $this->internal->aliasBy = $aliasBy;
    
        return $this;
    }
    /**
     * GCM query type.
     * queryType: #QueryType
     * Time Series List sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesList> $timeSeriesList
     */
    public function timeSeriesList(\Grafana\Foundation\Cog\Builder $timeSeriesList): static
    {
        $timeSeriesListResource = $timeSeriesList->build();
        $this->internal->timeSeriesList = $timeSeriesListResource;
    
        return $this;
    }
    /**
     * Time Series sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQuery> $timeSeriesQuery
     */
    public function timeSeriesQuery(\Grafana\Foundation\Cog\Builder $timeSeriesQuery): static
    {
        $timeSeriesQueryResource = $timeSeriesQuery->build();
        $this->internal->timeSeriesQuery = $timeSeriesQueryResource;
    
        return $this;
    }
    /**
     * SLO sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\SLOQuery> $sloQuery
     */
    public function sloQuery(\Grafana\Foundation\Cog\Builder $sloQuery): static
    {
        $sloQueryResource = $sloQuery->build();
        $this->internal->sloQuery = $sloQueryResource;
    
        return $this;
    }
    /**
     * PromQL sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\PromQLQuery> $promQLQuery
     */
    public function promQLQuery(\Grafana\Foundation\Cog\Builder $promQLQuery): static
    {
        $promQLQueryResource = $promQLQuery->build();
        $this->internal->promQLQuery = $promQLQueryResource;
    
        return $this;
    }
    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public function datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource): static
    {
        $this->internal->datasource = $datasource;
    
        return $this;
    }
    /**
     * Time interval in milliseconds.
     */
    public function intervalMs(float $intervalMs): static
    {
        $this->internal->intervalMs = $intervalMs;
    
        return $this;
    }

}
