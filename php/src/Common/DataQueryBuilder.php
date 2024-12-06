<?php

namespace Grafana\Foundation\Common;

/**
 * These are the common properties available to all queries in all datasources.
 * Specific implementations will *extend* this interface, adding the required
 * properties for the given context.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\DataQuery>
 */
class DataQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\DataQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\DataQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\DataQuery
     */
    public function build()
    {
        return $this->internal;
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
     * @param mixed $datasource
     */
    public function datasource( $datasource): static
    {
        $this->internal->datasource = $datasource;
    
        return $this;
    }

}
