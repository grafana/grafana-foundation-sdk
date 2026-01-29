<?php

namespace Grafana\Foundation\Loki;

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
    $this->internal->group = "loki";
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

    /**
     * The LogQL query.
     */
    public function expr(string $expr): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Loki\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Loki\Dataquery);
        $this->internal->spec->expr = $expr;
    
        return $this;
    }

    /**
     * Used to override the name of the series.
     */
    public function legendFormat(string $legendFormat): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Loki\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Loki\Dataquery);
        $this->internal->spec->legendFormat = $legendFormat;
    
        return $this;
    }

    /**
     * Used to limit the number of log rows returned.
     */
    public function maxLines(int $maxLines): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Loki\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Loki\Dataquery);
        $this->internal->spec->maxLines = $maxLines;
    
        return $this;
    }

    /**
     * @deprecated, now use step.
     */
    public function resolution(int $resolution): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Loki\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Loki\Dataquery);
        $this->internal->spec->resolution = $resolution;
    
        return $this;
    }

    public function editorMode(\Grafana\Foundation\Loki\QueryEditorMode $editorMode): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Loki\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Loki\Dataquery);
        $this->internal->spec->editorMode = $editorMode;
    
        return $this;
    }

    /**
     * @deprecated, now use queryType.
     */
    public function range(bool $range): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Loki\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Loki\Dataquery);
        $this->internal->spec->range = $range;
    
        return $this;
    }

    /**
     * @deprecated, now use queryType.
     */
    public function instant(bool $instant): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Loki\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Loki\Dataquery);
        $this->internal->spec->instant = $instant;
    
        return $this;
    }

    /**
     * Used to set step value for range queries.
     */
    public function step(string $step): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Loki\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Loki\Dataquery);
        $this->internal->spec->step = $step;
    
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
            $this->internal->spec = new \Grafana\Foundation\Loki\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Loki\Dataquery);
        $this->internal->spec->refId = $refId;
    
        return $this;
    }

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Loki\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Loki\Dataquery);
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
            $this->internal->spec = new \Grafana\Foundation\Loki\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Loki\Dataquery);
        $this->internal->spec->queryType = $queryType;
    
        return $this;
    }

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public function oldDatasource(\Grafana\Foundation\Common\DataSourceRef $datasource): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Loki\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Loki\Dataquery);
        $this->internal->spec->datasource = $datasource;
    
        return $this;
    }

    public function direction(\Grafana\Foundation\Loki\LokiQueryDirection $direction): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Loki\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Loki\Dataquery);
        $this->internal->spec->direction = $direction;
    
        return $this;
    }

}
