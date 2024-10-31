<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeReduce>
 */
class TypeReduceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\TypeReduce $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\TypeReduce();
    $this->internal->type = "reduce";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Expr\TypeReduce
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The datasource
     */
    public function datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource): static
    {
        $this->internal->datasource = $datasource;
    
        return $this;
    }
    /**
     * Reference to single query result
     */
    public function expression(string $expression): static
    {
        if (!(strlen($expression) >= 1)) {
            throw new \ValueError('strlen($expression) must be >= 1');
        }
        $this->internal->expression = $expression;
    
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
     * The reducer
     * Possible enum values:
     *  - `"sum"` 
     *  - `"mean"` 
     *  - `"min"` 
     *  - `"max"` 
     *  - `"count"` 
     *  - `"last"` 
     *  - `"median"` 
     */
    public function reducer(\Grafana\Foundation\Expr\TypeReduceReducer $reducer): static
    {
        $this->internal->reducer = $reducer;
    
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
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeReduceResultAssertions> $resultAssertions
     */
    public function resultAssertions(\Grafana\Foundation\Cog\Builder $resultAssertions): static
    {
        $resultAssertionsResource = $resultAssertions->build();
        $this->internal->resultAssertions = $resultAssertionsResource;
    
        return $this;
    }
    /**
     * Reducer Options
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeReduceSettings> $settings
     */
    public function settings(\Grafana\Foundation\Cog\Builder $settings): static
    {
        $settingsResource = $settings->build();
        $this->internal->settings = $settingsResource;
    
        return $this;
    }
    /**
     * TimeRange represents the query range
     * NOTE: unlike generic /ds/query, we can now send explicit time values in each query
     * NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeReduceTimeRange> $timeRange
     */
    public function timeRange(\Grafana\Foundation\Cog\Builder $timeRange): static
    {
        $timeRangeResource = $timeRange->build();
        $this->internal->timeRange = $timeRangeResource;
    
        return $this;
    }

}
