<?php

namespace Grafana\Foundation\Datasource;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Datasource\Dataquery>
 */
class DataqueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Datasource\Dataquery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Datasource\Dataquery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Datasource\Dataquery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Panel ID from wich the queries will be reused.
     */
    public function panelId(int $panelId): static
    {
        $this->internal->panelId = $panelId;
    
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

    public function withTransforms(bool $withTransforms): static
    {
        $this->internal->withTransforms = $withTransforms;
    
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
