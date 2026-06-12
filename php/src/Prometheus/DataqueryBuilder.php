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
     * Additional Ad-hoc filters that take precedence over Scope on conflict.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\AdhocFilters>> $adhocFilters
     */
    public function adhocFilters(array $adhocFilters): static
    {
            $adhocFiltersResources = [];
            foreach ($adhocFilters as $r1) {
                    $adhocFiltersResources[] = $r1->build();
            }
        $this->internal->adhocFilters = $adhocFiltersResources;
    
        return $this;
    }

    /**
     * The datasource
     */
    public function datasource(\Grafana\Foundation\Common\DataSourceRef $datasource): static
    {
        $this->internal->datasource = $datasource;
    
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
        $this->internal->editorMode = $editorMode;
    
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
     * The actual expression/query that will be evaluated by Prometheus
     */
    public function expr(string $expr): static
    {
        $this->internal->expr = $expr;
    
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
        $this->internal->format = $format;
    
        return $this;
    }

    /**
     * Group By parameters to apply to aggregate expressions in the query
     * @param array<string> $groupByKeys
     */
    public function groupByKeys(array $groupByKeys): static
    {
        $this->internal->groupByKeys = $groupByKeys;
    
        return $this;
    }

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * NOTE: this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
     */
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
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
     * Used to specify how many times to divide max data points by. We use max data points under query options
     * See https://github.com/grafana/grafana/issues/48081
     * Deprecated: use interval
     */
    public function intervalFactor(int $intervalFactor): static
    {
        $this->internal->intervalFactor = $intervalFactor;
    
        return $this;
    }

    /**
     * Interval is the suggested duration between time points in a time series query.
     * NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
     * from the interval required to fill a pixels in the visualization
     */
    public function intervalMs(float $intervalMs): static
    {
        $this->internal->intervalMs = $intervalMs;
    
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
     * MaxDataPoints is the maximum number of data points that should be returned from a time series query.
     * NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
     * from the number of pixels visible in a visualization
     */
    public function maxDataPoints(int $maxDataPoints): static
    {
        $this->internal->maxDataPoints = $maxDataPoints;
    
        return $this;
    }

    /**
     * QueryType is an optional identifier for the type of query.
     * It can be used to distinguish different types of queries.
     */
    public function queryType(string $queryType): static
    {
        $this->internal->queryType = $queryType;
    
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
     * RefID is the unique identifier of the query, set by the frontend call.
     */
    public function refId(string $refId): static
    {
        $this->internal->refId = $refId;
    
        return $this;
    }

    /**
     * Optionally define expected query result behavior
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\ResultAssertions> $resultAssertions
     */
    public function resultAssertions(\Grafana\Foundation\Cog\Builder $resultAssertions): static
    {
        $resultAssertionsResource = $resultAssertions->build();
        $this->internal->resultAssertions = $resultAssertionsResource;
    
        return $this;
    }

    /**
     * A set of filters applied to apply to the query
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\Scopes>> $scopes
     */
    public function scopes(array $scopes): static
    {
            $scopesResources = [];
            foreach ($scopes as $r1) {
                    $scopesResources[] = $r1->build();
            }
        $this->internal->scopes = $scopesResources;
    
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
        $timeRangeResource = $timeRange->build();
        $this->internal->timeRange = $timeRangeResource;
    
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
