<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Dataquery>
 */
class DataqueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\Dataquery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\Dataquery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\Dataquery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Alias pattern
     */
    public function alias(string $alias): static
    {
        $this->internal->alias = $alias;
    
        return $this;
    }
    /**
     * Lucene query
     */
    public function query(string $query): static
    {
        $this->internal->query = $query;
    
        return $this;
    }
    /**
     * Name of time field
     */
    public function timeField(string $timeField): static
    {
        $this->internal->timeField = $timeField;
    
        return $this;
    }
    /**
     * List of bucket aggregations
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\DateHistogram>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Histogram>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Terms>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Filters>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\GeoHashGrid>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Nested>> $bucketAggs
     */
    public function bucketAggs(array $bucketAggs): static
    {
            $bucketAggsResources = [];
            foreach ($bucketAggs as $r1) {
                    $bucketAggsResources[] = $r1->build();
            }
        $this->internal->bucketAggs = $bucketAggsResources;
    
        return $this;
    }
    /**
     * List of metric aggregations
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Count>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\MovingAverage>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Derivative>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\CumulativeSum>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\BucketScript>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\SerialDiff>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\RawData>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\RawDocument>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\UniqueCount>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Percentiles>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ExtendedStats>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Min>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Max>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Sum>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Average>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\MovingFunction>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Logs>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Rate>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\TopMetrics>> $metrics
     */
    public function metrics(array $metrics): static
    {
            $metricsResources = [];
            foreach ($metrics as $r1) {
                    $metricsResources[] = $r1->build();
            }
        $this->internal->metrics = $metricsResources;
    
        return $this;
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

}
