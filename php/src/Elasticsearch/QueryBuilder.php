<?php

namespace Grafana\Foundation\Elasticsearch;

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
    $this->internal->group = "elasticsearch";
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
     * Alias pattern
     */
    public function alias(string $alias): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Elasticsearch\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery);
        $this->internal->spec->alias = $alias;
    
        return $this;
    }

    /**
     * Lucene query
     */
    public function query(string $query): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Elasticsearch\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery);
        $this->internal->spec->query = $query;
    
        return $this;
    }

    /**
     * Name of time field
     */
    public function timeField(string $timeField): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Elasticsearch\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery);
        $this->internal->spec->timeField = $timeField;
    
        return $this;
    }

    /**
     * List of bucket aggregations
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\DateHistogram>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Histogram>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Terms>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Filters>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\GeoHashGrid>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Nested>> $bucketAggs
     */
    public function bucketAggs(array $bucketAggs): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Elasticsearch\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery);
            $bucketAggsResources = [];
            foreach ($bucketAggs as $r1) {
                    $bucketAggsResources[] = $r1->build();
            }
        $this->internal->spec->bucketAggs = $bucketAggsResources;
    
        return $this;
    }

    /**
     * List of metric aggregations
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Count>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\MovingAverage>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Derivative>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\CumulativeSum>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\BucketScript>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\SerialDiff>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\RawData>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\RawDocument>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\UniqueCount>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Percentiles>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ExtendedStats>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Min>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Max>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Sum>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Average>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\MovingFunction>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Logs>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Rate>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\TopMetrics>> $metrics
     */
    public function metrics(array $metrics): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Elasticsearch\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery);
            $metricsResources = [];
            foreach ($metrics as $r1) {
                    $metricsResources[] = $r1->build();
            }
        $this->internal->spec->metrics = $metricsResources;
    
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
            $this->internal->spec = new \Grafana\Foundation\Elasticsearch\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery);
        $this->internal->spec->refId = $refId;
    
        return $this;
    }

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Elasticsearch\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery);
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
            $this->internal->spec = new \Grafana\Foundation\Elasticsearch\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery);
        $this->internal->spec->queryType = $queryType;
    
        return $this;
    }

}
