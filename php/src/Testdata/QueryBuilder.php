<?php

namespace Grafana\Foundation\Testdata;

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
    $this->internal->group = "testdata";
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

    public function alias(string $alias): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->alias = $alias;
    
        return $this;
    }

    /**
     * Used for live query
     */
    public function channel(string $channel): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->channel = $channel;
    
        return $this;
    }

    public function csvContent(string $csvContent): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->csvContent = $csvContent;
    
        return $this;
    }

    public function csvFileName(string $csvFileName): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->csvFileName = $csvFileName;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\CSVWave>> $csvWave
     */
    public function csvWave(array $csvWave): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
            $csvWaveResources = [];
            foreach ($csvWave as $r1) {
                    $csvWaveResources[] = $r1->build();
            }
        $this->internal->spec->csvWave = $csvWaveResources;
    
        return $this;
    }

    /**
     * The datasource
     */
    public function oldDatasource(\Grafana\Foundation\Common\DataSourceRef $datasource): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->datasource = $datasource;
    
        return $this;
    }

    /**
     * Drop percentage (the chance we will lose a point 0-100)
     */
    public function dropPercent(float $dropPercent): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->dropPercent = $dropPercent;
    
        return $this;
    }

    /**
     * Possible enum values:
     *  - `"plugin"` 
     *  - `"downstream"` 
     */
    public function errorSource(\Grafana\Foundation\Testdata\DataqueryErrorSource $errorSource): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->errorSource = $errorSource;
    
        return $this;
    }

    /**
     * Possible enum values:
     *  - `"frontend_exception"` 
     *  - `"frontend_observable"` 
     *  - `"server_panic"` 
     */
    public function errorType(\Grafana\Foundation\Testdata\DataqueryErrorType $errorType): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->errorType = $errorType;
    
        return $this;
    }

    public function flamegraphDiff(bool $flamegraphDiff): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->flamegraphDiff = $flamegraphDiff;
    
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
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->hide = $hide;
    
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
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->intervalMs = $intervalMs;
    
        return $this;
    }

    public function labels(string $labels): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->labels = $labels;
    
        return $this;
    }

    public function levelColumn(bool $levelColumn): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->levelColumn = $levelColumn;
    
        return $this;
    }

    public function lines(int $lines): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->lines = $lines;
    
        return $this;
    }

    public function max(float $max): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->max = $max;
    
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
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->maxDataPoints = $maxDataPoints;
    
        return $this;
    }

    public function min(float $min): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->min = $min;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\NodesQuery> $nodes
     */
    public function nodes(\Grafana\Foundation\Cog\Builder $nodes): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $nodesResource = $nodes->build();
        $this->internal->spec->nodes = $nodesResource;
    
        return $this;
    }

    public function noise(float $noise): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->noise = $noise;
    
        return $this;
    }

    /**
     * @param array<array<mixed>> $points
     */
    public function points(array $points): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->points = $points;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\PulseWaveQuery> $pulseWave
     */
    public function pulseWave(\Grafana\Foundation\Cog\Builder $pulseWave): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $pulseWaveResource = $pulseWave->build();
        $this->internal->spec->pulseWave = $pulseWaveResource;
    
        return $this;
    }

    /**
     * QueryType is an optional identifier for the type of query.
     * It can be used to distinguish different types of queries.
     */
    public function queryType(string $queryType): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->queryType = $queryType;
    
        return $this;
    }

    public function rawFrameContent(string $rawFrameContent): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->rawFrameContent = $rawFrameContent;
    
        return $this;
    }

    /**
     * RefID is the unique identifier of the query, set by the frontend call.
     */
    public function refId(string $refId): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->refId = $refId;
    
        return $this;
    }

    /**
     * Optionally define expected query result behavior
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\ResultAssertions> $resultAssertions
     */
    public function resultAssertions(\Grafana\Foundation\Cog\Builder $resultAssertions): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $resultAssertionsResource = $resultAssertions->build();
        $this->internal->spec->resultAssertions = $resultAssertionsResource;
    
        return $this;
    }

    /**
     * Possible enum values:
     *  - `"annotations"` 
     *  - `"arrow"` 
     *  - `"csv_content"` 
     *  - `"csv_file"` 
     *  - `"csv_metric_values"` 
     *  - `"datapoints_outside_range"` 
     *  - `"error_with_source"` 
     *  - `"exponential_heatmap_bucket_data"` 
     *  - `"flame_graph"` 
     *  - `"grafana_api"` 
     *  - `"linear_heatmap_bucket_data"` 
     *  - `"live"` 
     *  - `"logs"` 
     *  - `"manual_entry"` 
     *  - `"no_data_points"` 
     *  - `"node_graph"` 
     *  - `"predictable_csv_wave"` 
     *  - `"predictable_pulse"` 
     *  - `"query_meta"` 
     *  - `"random_walk"` 
     *  - `"random_walk_table"` 
     *  - `"random_walk_with_error"` 
     *  - `"raw_frame"` 
     *  - `"server_error_500"` 
     *  - `"steps"` 
     *  - `"simulation"` 
     *  - `"slow_query"` 
     *  - `"streaming_client"` 
     *  - `"table_static"` 
     *  - `"trace"` 
     *  - `"usa"` 
     *  - `"variables-query"` 
     */
    public function scenarioId(\Grafana\Foundation\Testdata\DataqueryScenarioId $scenarioId): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->scenarioId = $scenarioId;
    
        return $this;
    }

    public function seriesCount(int $seriesCount): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->seriesCount = $seriesCount;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\SimulationQuery> $sim
     */
    public function sim(\Grafana\Foundation\Cog\Builder $sim): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $simResource = $sim->build();
        $this->internal->spec->sim = $simResource;
    
        return $this;
    }

    public function spanCount(int $spanCount): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->spanCount = $spanCount;
    
        return $this;
    }

    public function spread(float $spread): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->spread = $spread;
    
        return $this;
    }

    public function startValue(float $startValue): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->startValue = $startValue;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\StreamingQuery> $stream
     */
    public function stream(\Grafana\Foundation\Cog\Builder $stream): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $streamResource = $stream->build();
        $this->internal->spec->stream = $streamResource;
    
        return $this;
    }

    /**
     * common parameter used by many query types
     */
    public function stringInput(string $stringInput): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->stringInput = $stringInput;
    
        return $this;
    }

    /**
     * TimeRange represents the query range
     * NOTE: unlike generic /ds/query, we can now send explicit time values in each query
     * NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\TimeRange> $timeRange
     */
    public function timeRange(\Grafana\Foundation\Cog\Builder $timeRange): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $timeRangeResource = $timeRange->build();
        $this->internal->spec->timeRange = $timeRangeResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\USAQuery> $usa
     */
    public function usa(\Grafana\Foundation\Cog\Builder $usa): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $usaResource = $usa->build();
        $this->internal->spec->usa = $usaResource;
    
        return $this;
    }

    public function withNil(bool $withNil): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Testdata\Dataquery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Testdata\Dataquery);
        $this->internal->spec->withNil = $withNil;
    
        return $this;
    }

}
