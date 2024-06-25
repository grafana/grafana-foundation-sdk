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
     * @return \Grafana\Foundation\Alerting\Query
     */
    public function build()
    {
        return $this->internal;
    }

    public function datasourceUid(string $datasourceUid): static
    {
        $this->internal->datasourceUid = $datasourceUid;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cog\Dataquery> $model
     */
    public function model(\Grafana\Foundation\Cog\Builder $model): static
    {
        $modelResource = $model->build();
        $this->internal->model = $modelResource;
    
        return $this;
    }
    public function queryType(string $queryType): static
    {
        $this->internal->queryType = $queryType;
    
        return $this;
    }
    public function refId(string $refId): static
    {
        $this->internal->refId = $refId;
    
        return $this;
    }
    public function relativeTimeRange(\Grafana\Foundation\Alerting\RelativeTimeRange $relativeTimeRange): static
    {
        $this->internal->relativeTimeRange = $relativeTimeRange;
    
        return $this;
    }

}
