<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\DataQueryKind>
 */
class QueryV2Builder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\DataQueryKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\DataQueryKind();
    $this->internal->kind = "DataQuery";
    $this->internal->group = "cloudwatch";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\DataQueryKind
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
     * @param array<string, string> $labels
     */
    public function labels(array $labels): static
    {
        $this->internal->labels = $labels;
    
        return $this;
    }

    /**
     * New type for datasource reference
     * Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\Dashboardv2DataQueryKindDatasource> $datasource
     */
    public function datasource(\Grafana\Foundation\Cog\Builder $datasource): static
    {
        $datasourceResource = $datasource->build();
        $this->internal->datasource = $datasourceResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\MetricsQuery> $metricsQuery
     */
    public function metricsQuery(\Grafana\Foundation\Cog\Builder $metricsQuery): static
    {
        $metricsQueryResource = $metricsQuery->build();
        $this->internal->spec = $metricsQueryResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\LogsQuery> $logsQuery
     */
    public function logsQuery(\Grafana\Foundation\Cog\Builder $logsQuery): static
    {
        $logsQueryResource = $logsQuery->build();
        $this->internal->spec = $logsQueryResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\AnnotationQuery> $annotationQuery
     */
    public function annotationQuery(\Grafana\Foundation\Cog\Builder $annotationQuery): static
    {
        $annotationQueryResource = $annotationQuery->build();
        $this->internal->spec = $annotationQueryResource;
    
        return $this;
    }

}
