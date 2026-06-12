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
     * @param array<string, string> $labels
     */
    public function labels(array $labels): static
    {
        $this->internal->labels = $labels;
    
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
     * Additional Ad-hoc filters that take precedence over Scope on conflict.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\AdhocFilters>> $adhocFilters
     */
    public function adhocFilters(array $adhocFilters): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
            $adhocFiltersResources = [];
            foreach ($adhocFilters as $r1) {
                    $adhocFiltersResources[] = $r1->build();
            }
        $this->internal->spec->adhocFilters = $adhocFiltersResources;
    
        return $this;
    }

    /**
     * what we should show in the editor
     * Possible enum values:
     *  - `"builder"` 
     *  - `"code"` 
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
     * The response format
     * Possible enum values:
     *  - `"time_series"` 
     *  - `"table"` 
     *  - `"heatmap"` 
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
     * Group By parameters to apply to aggregate expressions in the query
     * @param array<string> $groupByKeys
     */
    public function groupByKeys(array $groupByKeys): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->groupByKeys = $groupByKeys;
    
        return $this;
    }

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * NOTE: this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
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
     * Returns only the latest value that Prometheus has scraped for the requested time series
     */
    public function instant(): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->instant = true;
        $this->internal->spec->range = false;
    
        return $this;
    }

    /**
     * Used to specify how many times to divide max data points by. We use max data points under query options
     * See https://github.com/grafana/grafana/issues/48081
     * Deprecated: use interval
     */
    public function intervalFactor(int $intervalFactor): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->intervalFactor = $intervalFactor;
    
        return $this;
    }

    /**
     * Interval is the suggested duration between time points in a time series query.
     * NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
     * from the interval required to fill a pixels in the visualization
     */
    public function intervalMs(float $intervalMs): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->intervalMs = $intervalMs;
    
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
     * MaxDataPoints is the maximum number of data points that should be returned from a time series query.
     * NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
     * from the number of pixels visible in a visualization
     */
    public function maxDataPoints(int $maxDataPoints): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->maxDataPoints = $maxDataPoints;
    
        return $this;
    }

    /**
     * QueryType is an optional identifier for the type of query.
     * It can be used to distinguish different types of queries.
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
     * Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
     */
    public function range(): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $this->internal->spec->range = true;
        $this->internal->spec->instant = false;
    
        return $this;
    }

    /**
     * RefID is the unique identifier of the query, set by the frontend call.
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
     * Optionally define expected query result behavior
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\ResultAssertions> $resultAssertions
     */
    public function resultAssertions(\Grafana\Foundation\Cog\Builder $resultAssertions): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $resultAssertionsResource = $resultAssertions->build();
        $this->internal->spec->resultAssertions = $resultAssertionsResource;
    
        return $this;
    }

    /**
     * A set of filters applied to apply to the query
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\Scopes>> $scopes
     */
    public function scopes(array $scopes): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
            $scopesResources = [];
            foreach ($scopes as $r1) {
                    $scopesResources[] = $r1->build();
            }
        $this->internal->spec->scopes = $scopesResources;
    
        return $this;
    }

    /**
     * TimeRange represents the query range
     * NOTE: unlike generic /ds/query, we can now send explicit time values in each query
     * NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\TimeRange> $timeRange
     */
    public function timeRange(\Grafana\Foundation\Cog\Builder $timeRange): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Prometheus\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Prometheus\Dataquery);
        $timeRangeResource = $timeRange->build();
        $this->internal->spec->timeRange = $timeRangeResource;
    
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
