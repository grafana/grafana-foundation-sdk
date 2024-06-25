<?php

namespace Grafana\Foundation\Grafanapyroscope;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Grafanapyroscope\Dataquery>
 */
class DataqueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Grafanapyroscope\Dataquery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Grafanapyroscope\Dataquery();
    }

    /**
     * @return \Grafana\Foundation\Grafanapyroscope\Dataquery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Specifies the query label selectors.
     */
    public function labelSelector(string $labelSelector): static
    {
        $this->internal->labelSelector = $labelSelector;
    
        return $this;
    }
    /**
     * Specifies the type of profile to query.
     */
    public function profileTypeId(string $profileTypeId): static
    {
        $this->internal->profileTypeId = $profileTypeId;
    
        return $this;
    }
    /**
     * Allows to group the results.
     * @param array<string> $groupBy
     */
    public function groupBy(array $groupBy): static
    {
        $this->internal->groupBy = $groupBy;
    
        return $this;
    }
    /**
     * Sets the maximum number of nodes in the flamegraph.
     */
    public function maxNodes(int $maxNodes): static
    {
        $this->internal->maxNodes = $maxNodes;
    
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
     * @param mixed $datasource
     */
    public function datasource( $datasource): static
    {
        $this->internal->datasource = $datasource;
    
        return $this;
    }

}
