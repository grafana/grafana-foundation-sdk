<?php

namespace Grafana\Foundation\Bigquery;

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
    $this->internal->group = "grafana-bigquery-datasource";
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

    public function dataset(string $dataset): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->dataset = $dataset;
    
        return $this;
    }

    public function table(string $table): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->table = $table;
    
        return $this;
    }

    public function project(string $project): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->project = $project;
    
        return $this;
    }

    public function format(\Grafana\Foundation\Bigquery\QueryFormat $format): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->format = $format;
    
        return $this;
    }

    public function rawQuery(bool $rawQuery): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->rawQuery = $rawQuery;
    
        return $this;
    }

    public function rawSql(string $rawSql): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->rawSql = $rawSql;
    
        return $this;
    }

    public function location(string $location): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->location = $location;
    
        return $this;
    }

    public function partitioned(bool $partitioned): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->partitioned = $partitioned;
    
        return $this;
    }

    public function partitionedField(string $partitionedField): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->partitionedField = $partitionedField;
    
        return $this;
    }

    public function convertToUTC(bool $convertToUTC): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->convertToUTC = $convertToUTC;
    
        return $this;
    }

    public function sharded(bool $sharded): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->sharded = $sharded;
    
        return $this;
    }

    public function queryPriority(\Grafana\Foundation\Bigquery\QueryPriority $queryPriority): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->queryPriority = $queryPriority;
    
        return $this;
    }

    public function timeShift(string $timeShift): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->timeShift = $timeShift;
    
        return $this;
    }

    public function editorMode(\Grafana\Foundation\Bigquery\EditorMode $editorMode): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->editorMode = $editorMode;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\SQLExpression> $sql
     */
    public function sql(\Grafana\Foundation\Cog\Builder $sql): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $sqlResource = $sql->build();
        $this->internal->spec->sql = $sqlResource;
    
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
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->refId = $refId;
    
        return $this;
    }

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
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
            $this->internal->spec = new \Grafana\Foundation\Bigquery\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $this->internal->spec->queryType = $queryType;
    
        return $this;
    }

}
