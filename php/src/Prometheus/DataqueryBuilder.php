<?php

namespace Grafana\Foundation\Prometheus;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\Dataquery>
 */
class DataqueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Prometheus\Dataquery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Prometheus\Dataquery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Prometheus\Dataquery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The actual expression/query that will be evaluated by Prometheus
     */
    public function expr(string $expr): static
    {
        $this->internal->expr = $expr;
    
        return $this;
    }
    /**
     * Returns only the latest value that Prometheus has scraped for the requested time series
     */
    public function instant(): static
    {
        $this->internal->instant = true;
        $this->internal->range = false;
    
        return $this;
    }
    /**
     * Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
     */
    public function range(): static
    {
        $this->internal->range = true;
        $this->internal->instant = false;
    
        return $this;
    }
    /**
     * Execute an additional query to identify interesting raw samples relevant for the given expr
     */
    public function exemplar(bool $exemplar): static
    {
        $this->internal->exemplar = $exemplar;
    
        return $this;
    }
    /**
     * Specifies which editor is being used to prepare the query. It can be "code" or "builder"
     */
    public function editorMode(\Grafana\Foundation\Prometheus\QueryEditorMode $editorMode): static
    {
        $this->internal->editorMode = $editorMode;
    
        return $this;
    }
    /**
     * Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
     */
    public function format(\Grafana\Foundation\Prometheus\PromQueryFormat $format): static
    {
        $this->internal->format = $format;
    
        return $this;
    }
    /**
     * Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
     */
    public function legendFormat(string $legendFormat): static
    {
        $this->internal->legendFormat = $legendFormat;
    
        return $this;
    }
    /**
     * @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
     * See https://github.com/grafana/grafana/issues/48081
     */
    public function intervalFactor(float $intervalFactor): static
    {
        $this->internal->intervalFactor = $intervalFactor;
    
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
    /**
     * An additional lower limit for the step parameter of the Prometheus query and for the
     * `$__interval` and `$__rate_interval` variables.
     */
    public function interval(string $interval): static
    {
        $this->internal->interval = $interval;
    
        return $this;
    }
    public function rangeAndInstant(): static
    {
        $this->internal->range = true;
        $this->internal->instant = true;
    
        return $this;
    }

}
