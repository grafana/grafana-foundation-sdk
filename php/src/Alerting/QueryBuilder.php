<?php

namespace Grafana\Foundation\Alerting;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Query>
 */
class QueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\Query $internal;

    public function __construct(string $refId)
    {
    	$this->internal = new \Grafana\Foundation\Alerting\Query();
    $this->internal->refId = $refId;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alerting\Query
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
     */
    public function datasourceUid(string $datasourceUid): static
    {
        $this->internal->datasourceUid = $datasourceUid;
    
        return $this;
    }
    /**
     * JSON is the raw JSON query and includes the above properties as well as custom properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cog\Dataquery> $model
     */
    public function model(\Grafana\Foundation\Cog\Builder $model): static
    {
        $modelResource = $model->build();
        $this->internal->model = $modelResource;
    
        return $this;
    }
    /**
     * QueryType is an optional identifier for the type of query.
     * It can be used to distinguish different types of queries.
     */
    public function queryType(string $queryType): static
    {
        $this->internal->queryType = $queryType;
    
        return $this;
    }
    /**
     * RefID is the unique identifier of the query, set by the frontend call.
     */
    public function refId(string $refId): static
    {
        $this->internal->refId = $refId;
    
        return $this;
    }
    /**
     * RelativeTimeRange is the per query start and end time
     * for requests.
     */
    public function relativeTimeRange(\Grafana\Foundation\Alerting\RelativeTimeRange $relativeTimeRange): static
    {
        $this->internal->relativeTimeRange = $relativeTimeRange;
    
        return $this;
    }

}
