<?php

namespace Grafana\Foundation\Testdata;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\Dataquery>
 */
class DataqueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Testdata\Dataquery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Testdata\Dataquery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Testdata\Dataquery
     */
    public function build()
    {
        return $this->internal;
    }

    public function alias(string $alias): static
    {
        $this->internal->alias = $alias;
    
        return $this;
    }
    /**
     * Used for live query
     */
    public function channel(string $channel): static
    {
        $this->internal->channel = $channel;
    
        return $this;
    }
    public function csvContent(string $csvContent): static
    {
        $this->internal->csvContent = $csvContent;
    
        return $this;
    }
    public function csvFileName(string $csvFileName): static
    {
        $this->internal->csvFileName = $csvFileName;
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\CSVWave>> $csvWave
     */
    public function csvWave(array $csvWave): static
    {
            $csvWaveResources = [];
            foreach ($csvWave as $r1) {
                    $csvWaveResources[] = $r1->build();
            }
        $this->internal->csvWave = $csvWaveResources;
    
        return $this;
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
     * Drop percentage (the chance we will lose a point 0-100)
     */
    public function dropPercent(float $dropPercent): static
    {
        $this->internal->dropPercent = $dropPercent;
    
        return $this;
    }
    /**
     * Possible enum values:
     *  - `"plugin"` 
     *  - `"downstream"` 
     */
    public function errorSource(\Grafana\Foundation\Testdata\DataqueryErrorSource $errorSource): static
    {
        $this->internal->errorSource = $errorSource;
    
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
        $this->internal->errorType = $errorType;
    
        return $this;
    }
    public function flamegraphDiff(bool $flamegraphDiff): static
    {
        $this->internal->flamegraphDiff = $flamegraphDiff;
    
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
    public function labels(string $labels): static
    {
        $this->internal->labels = $labels;
    
        return $this;
    }
    public function levelColumn(bool $levelColumn): static
    {
        $this->internal->levelColumn = $levelColumn;
    
        return $this;
    }
    public function lines(int $lines): static
    {
        $this->internal->lines = $lines;
    
        return $this;
    }
    public function max(float $max): static
    {
        $this->internal->max = $max;
    
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
    public function min(float $min): static
    {
        $this->internal->min = $min;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\NodesQuery> $nodes
     */
    public function nodes(\Grafana\Foundation\Cog\Builder $nodes): static
    {
        $nodesResource = $nodes->build();
        $this->internal->nodes = $nodesResource;
    
        return $this;
    }
    public function noise(float $noise): static
    {
        $this->internal->noise = $noise;
    
        return $this;
    }
    /**
     * @param array<array<mixed>> $points
     */
    public function points(array $points): static
    {
        $this->internal->points = $points;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\PulseWaveQuery> $pulseWave
     */
    public function pulseWave(\Grafana\Foundation\Cog\Builder $pulseWave): static
    {
        $pulseWaveResource = $pulseWave->build();
        $this->internal->pulseWave = $pulseWaveResource;
    
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
    public function rawFrameContent(string $rawFrameContent): static
    {
        $this->internal->rawFrameContent = $rawFrameContent;
    
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
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\ResultAssertions> $resultAssertions
     */
    public function resultAssertions(\Grafana\Foundation\Cog\Builder $resultAssertions): static
    {
        $resultAssertionsResource = $resultAssertions->build();
        $this->internal->resultAssertions = $resultAssertionsResource;
    
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
     *  - `"random_walk"` 
     *  - `"random_walk_table"` 
     *  - `"random_walk_with_error"` 
     *  - `"raw_frame"` 
     *  - `"server_error_500"` 
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
        $this->internal->scenarioId = $scenarioId;
    
        return $this;
    }
    public function seriesCount(int $seriesCount): static
    {
        $this->internal->seriesCount = $seriesCount;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\SimulationQuery> $sim
     */
    public function sim(\Grafana\Foundation\Cog\Builder $sim): static
    {
        $simResource = $sim->build();
        $this->internal->sim = $simResource;
    
        return $this;
    }
    public function spanCount(int $spanCount): static
    {
        $this->internal->spanCount = $spanCount;
    
        return $this;
    }
    public function spread(float $spread): static
    {
        $this->internal->spread = $spread;
    
        return $this;
    }
    public function startValue(float $startValue): static
    {
        $this->internal->startValue = $startValue;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\StreamingQuery> $stream
     */
    public function stream(\Grafana\Foundation\Cog\Builder $stream): static
    {
        $streamResource = $stream->build();
        $this->internal->stream = $streamResource;
    
        return $this;
    }
    /**
     * common parameter used by many query types
     */
    public function stringInput(string $stringInput): static
    {
        $this->internal->stringInput = $stringInput;
    
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
        $timeRangeResource = $timeRange->build();
        $this->internal->timeRange = $timeRangeResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\USAQuery> $usa
     */
    public function usa(\Grafana\Foundation\Cog\Builder $usa): static
    {
        $usaResource = $usa->build();
        $this->internal->usa = $usaResource;
    
        return $this;
    }
    public function withNil(bool $withNil): static
    {
        $this->internal->withNil = $withNil;
    
        return $this;
    }

}
