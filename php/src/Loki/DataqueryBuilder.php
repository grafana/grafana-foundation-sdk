<?php

namespace Grafana\Foundation\Loki;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Loki\Dataquery>
 */
class DataqueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Loki\Dataquery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Loki\Dataquery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Loki\Dataquery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The LogQL query.
     */
    public function expr(string $expr): static
    {
        $this->internal->expr = $expr;
    
        return $this;
    }
    /**
     * Used to override the name of the series.
     */
    public function legendFormat(string $legendFormat): static
    {
        $this->internal->legendFormat = $legendFormat;
    
        return $this;
    }
    /**
     * Used to limit the number of log rows returned.
     */
    public function maxLines(int $maxLines): static
    {
        $this->internal->maxLines = $maxLines;
    
        return $this;
    }
    /**
     * @deprecated, now use step.
     */
    public function resolution(int $resolution): static
    {
        $this->internal->resolution = $resolution;
    
        return $this;
    }
    public function editorMode(\Grafana\Foundation\Loki\QueryEditorMode $editorMode): static
    {
        $this->internal->editorMode = $editorMode;
    
        return $this;
    }
    /**
     * @deprecated, now use queryType.
     */
    public function range(bool $range): static
    {
        $this->internal->range = $range;
    
        return $this;
    }
    /**
     * @deprecated, now use queryType.
     */
    public function instant(bool $instant): static
    {
        $this->internal->instant = $instant;
    
        return $this;
    }
    /**
     * Used to set step value for range queries.
     */
    public function step(string $step): static
    {
        $this->internal->step = $step;
    
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
    public function datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource): static
    {
        $this->internal->datasource = $datasource;
    
        return $this;
    }

}
