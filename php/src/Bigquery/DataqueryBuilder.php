<?php

namespace Grafana\Foundation\Bigquery;

/**
 * Manually converted from https://github.com/grafana/google-bigquery-datasource/blob/18680e42ba557791d109c7c540c2c3f2647592f0/src/types.ts#L75
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\Dataquery>
 */
class DataqueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Bigquery\Dataquery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Bigquery\Dataquery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Bigquery\Dataquery
     */
    public function build()
    {
        return $this->internal;
    }

    public function dataset(string $dataset): static
    {
        $this->internal->dataset = $dataset;
    
        return $this;
    }

    public function table(string $table): static
    {
        $this->internal->table = $table;
    
        return $this;
    }

    public function project(string $project): static
    {
        $this->internal->project = $project;
    
        return $this;
    }

    public function format(\Grafana\Foundation\Bigquery\QueryFormat $format): static
    {
        $this->internal->format = $format;
    
        return $this;
    }

    public function rawQuery(bool $rawQuery): static
    {
        $this->internal->rawQuery = $rawQuery;
    
        return $this;
    }

    public function rawSql(string $rawSql): static
    {
        $this->internal->rawSql = $rawSql;
    
        return $this;
    }

    public function location(string $location): static
    {
        $this->internal->location = $location;
    
        return $this;
    }

    public function partitioned(bool $partitioned): static
    {
        $this->internal->partitioned = $partitioned;
    
        return $this;
    }

    public function partitionedField(string $partitionedField): static
    {
        $this->internal->partitionedField = $partitionedField;
    
        return $this;
    }

    public function convertToUTC(bool $convertToUTC): static
    {
        $this->internal->convertToUTC = $convertToUTC;
    
        return $this;
    }

    public function sharded(bool $sharded): static
    {
        $this->internal->sharded = $sharded;
    
        return $this;
    }

    public function queryPriority(\Grafana\Foundation\Bigquery\QueryPriority $queryPriority): static
    {
        $this->internal->queryPriority = $queryPriority;
    
        return $this;
    }

    public function timeShift(string $timeShift): static
    {
        $this->internal->timeShift = $timeShift;
    
        return $this;
    }

    public function editorMode(\Grafana\Foundation\Bigquery\EditorMode $editorMode): static
    {
        $this->internal->editorMode = $editorMode;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\SQLExpression> $sql
     */
    public function sql(\Grafana\Foundation\Cog\Builder $sql): static
    {
        $sqlResource = $sql->build();
        $this->internal->sql = $sqlResource;
    
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
     * true if query is disabled (ie should not be returned to the dashboard)
     * Note this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
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
    public function datasource(\Grafana\Foundation\Common\DataSourceRef $datasource): static
    {
        $this->internal->datasource = $datasource;
    
        return $this;
    }

}
