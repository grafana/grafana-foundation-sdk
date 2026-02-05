<?php

namespace Grafana\Foundation\Athena;

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
    $this->internal->group = "grafana-athena-datasource";
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

    public function format(\Grafana\Foundation\Athena\FormatOptions $format): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Athena\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Athena\Dataquery);
        $this->internal->spec->format = $format;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Athena\ConnectionArgs> $connectionArgs
     */
    public function connectionArgs(\Grafana\Foundation\Cog\Builder $connectionArgs): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Athena\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Athena\Dataquery);
        $connectionArgsResource = $connectionArgs->build();
        $this->internal->spec->connectionArgs = $connectionArgsResource;
    
        return $this;
    }

    public function table(string $table): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Athena\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Athena\Dataquery);
        $this->internal->spec->table = $table;
    
        return $this;
    }

    public function column(string $column): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Athena\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Athena\Dataquery);
        $this->internal->spec->column = $column;
    
        return $this;
    }

    public function queryID(string $queryID): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Athena\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Athena\Dataquery);
        $this->internal->spec->queryID = $queryID;
    
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
            $this->internal->spec = new \Grafana\Foundation\Athena\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Athena\Dataquery);
        $this->internal->spec->refId = $refId;
    
        return $this;
    }

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Athena\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Athena\Dataquery);
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
            $this->internal->spec = new \Grafana\Foundation\Athena\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Athena\Dataquery);
        $this->internal->spec->queryType = $queryType;
    
        return $this;
    }

    public function rawSQL(string $rawSQL): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Athena\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Athena\Dataquery);
        $this->internal->spec->rawSQL = $rawSQL;
    
        return $this;
    }

}
