<?php

namespace Grafana\Foundation\Prometheus;

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
    $this->internal->group = "prometheus";
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
     * The actual expression/query that will be evaluated by Prometheus
     */
    public function expr(string $expr): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->expr = $expr;
    
        return $this;
    }

    /**
     * Returns only the latest value that Prometheus has scraped for the requested time series
     */
    public function instant(): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->instant = true;
    
        return $this;
    }

    /**
     * Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
     */
    public function range(): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->range = true;
    
        return $this;
    }

    /**
     * Execute an additional query to identify interesting raw samples relevant for the given expr
     */
    public function exemplar(bool $exemplar): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->exemplar = $exemplar;
    
        return $this;
    }

    /**
     * Specifies which editor is being used to prepare the query. It can be "code" or "builder"
     */
    public function editorMode(\Grafana\Foundation\Prometheus\QueryEditorMode $editorMode): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->editorMode = $editorMode;
    
        return $this;
    }

    /**
     * Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
     */
    public function format(\Grafana\Foundation\Prometheus\PromQueryFormat $format): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->format = $format;
    
        return $this;
    }

    /**
     * Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
     */
    public function legendFormat(string $legendFormat): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->legendFormat = $legendFormat;
    
        return $this;
    }

    /**
     * @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
     * See https://github.com/grafana/grafana/issues/48081
     */
    public function intervalFactor(float $intervalFactor): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->intervalFactor = $intervalFactor;
    
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
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->refId = $refId;
    
        return $this;
    }

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
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
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->queryType = $queryType;
    
        return $this;
    }

    /**
     * An additional lower limit for the step parameter of the Prometheus query and for the
     * `$__interval` and `$__rate_interval` variables.
     */
    public function interval(string $interval): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->interval = $interval;
    
        return $this;
    }

    public function rangeAndInstant(): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->range = true;
        $this->internal->spec->instant = true;
    
        return $this;
    }

}
