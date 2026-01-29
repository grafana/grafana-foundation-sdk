<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DataQueryKind>
 */
class QueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\DataQueryKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\DataQueryKind();
    $this->internal->kind = "DataQuery";
    $this->internal->group = "cloud-monitoring";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\DataQueryKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function version(string $version): static
    {
        $this->internal->version = $version;
    
        return $this;
    }

    /**
     * New type for datasource reference
     * Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource> $datasource
     */
    public function datasource(\Grafana\Foundation\Cog\Builder $datasource): static
    {
        $datasourceResource = $datasource->build();
        $this->internal->datasource = $datasourceResource;
    
        return $this;
    }

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public function refId(string $refId): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery);
        $this->internal->spec->refId = $refId;
    
        return $this;
    }

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery);
        $this->internal->spec->hide = $hide;
    
        return $this;
    }

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public function queryType(string $queryType): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery);
        $this->internal->spec->queryType = $queryType;
    
        return $this;
    }

    /**
     * Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
     */
    public function aliasBy(string $aliasBy): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery);
        $this->internal->spec->aliasBy = $aliasBy;
    
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
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery);
        $timeSeriesListResource = $timeSeriesList->build();
        $this->internal->spec->timeSeriesList = $timeSeriesListResource;
    
        return $this;
    }

    /**
     * Time Series sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQuery> $timeSeriesQuery
     */
    public function timeSeriesQuery(\Grafana\Foundation\Cog\Builder $timeSeriesQuery): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery);
        $timeSeriesQueryResource = $timeSeriesQuery->build();
        $this->internal->spec->timeSeriesQuery = $timeSeriesQueryResource;
    
        return $this;
    }

    /**
     * SLO sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\SLOQuery> $sloQuery
     */
    public function sloQuery(\Grafana\Foundation\Cog\Builder $sloQuery): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery);
        $sloQueryResource = $sloQuery->build();
        $this->internal->spec->sloQuery = $sloQueryResource;
    
        return $this;
    }

    /**
     * PromQL sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\PromQLQuery> $promQLQuery
     */
    public function promQLQuery(\Grafana\Foundation\Cog\Builder $promQLQuery): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery);
        $promQLQueryResource = $promQLQuery->build();
        $this->internal->spec->promQLQuery = $promQLQueryResource;
    
        return $this;
    }

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public function oldDatasource(\Grafana\Foundation\Common\DataSourceRef $datasource): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery);
        $this->internal->spec->datasource = $datasource;
    
        return $this;
    }

    /**
     * Time interval in milliseconds.
     */
    public function intervalMs(float $intervalMs): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery);
        $this->internal->spec->intervalMs = $intervalMs;
    
        return $this;
    }

}
