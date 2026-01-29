<?php

namespace Grafana\Foundation\Athena;

/**
 * Manually converted from https://github.com/grafana/athena-datasource/blob/57ad707147b7a11e9a521a836d6bf9799877e0e3/src/types.ts
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Athena\Dataquery>
 */
class DataqueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Athena\Dataquery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Athena\Dataquery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Athena\Dataquery
     */
    public function build()
    {
        return $this->internal;
    }

    public function format(\Grafana\Foundation\Athena\FormatOptions $format): static
    {
        $this->internal->format = $format;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Athena\ConnectionArgs> $connectionArgs
     */
    public function connectionArgs(\Grafana\Foundation\Cog\Builder $connectionArgs): static
    {
        $connectionArgsResource = $connectionArgs->build();
        $this->internal->connectionArgs = $connectionArgsResource;
    
        return $this;
    }

    public function table(string $table): static
    {
        $this->internal->table = $table;
    
        return $this;
    }

    public function column(string $column): static
    {
        $this->internal->column = $column;
    
        return $this;
    }

    public function queryID(string $queryID): static
    {
        $this->internal->queryID = $queryID;
    
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

    public function rawSQL(string $rawSQL): static
    {
        $this->internal->rawSQL = $rawSQL;
    
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
