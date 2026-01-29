<?php

namespace Grafana\Foundation\Cloudwatch;

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
    $this->internal->group = "cloudwatch";
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
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\CloudWatchMetricsQuery> $cloudWatchMetricsQuery
     */
    public function cloudWatchMetricsQuery(\Grafana\Foundation\Cog\Builder $cloudWatchMetricsQuery): static
    {
        $cloudWatchMetricsQueryResource = $cloudWatchMetricsQuery->build();
        $this->internal->spec = $cloudWatchMetricsQueryResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\CloudWatchLogsQuery> $cloudWatchLogsQuery
     */
    public function cloudWatchLogsQuery(\Grafana\Foundation\Cog\Builder $cloudWatchLogsQuery): static
    {
        $cloudWatchLogsQueryResource = $cloudWatchLogsQuery->build();
        $this->internal->spec = $cloudWatchLogsQueryResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\CloudWatchAnnotationQuery> $cloudWatchAnnotationQuery
     */
    public function cloudWatchAnnotationQuery(\Grafana\Foundation\Cog\Builder $cloudWatchAnnotationQuery): static
    {
        $cloudWatchAnnotationQueryResource = $cloudWatchAnnotationQuery->build();
        $this->internal->spec = $cloudWatchAnnotationQueryResource;
    
        return $this;
    }

}
