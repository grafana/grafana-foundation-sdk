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
    public function scenarioId(\Grafana\Foundation\Testdata\TestDataQueryType $scenarioId): static
    {
        $this->internal->scenarioId = $scenarioId;
    
        return $this;
    }
    public function stringInput(string $stringInput): static
    {
        $this->internal->stringInput = $stringInput;
    
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
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\PulseWaveQuery> $pulseWave
     */
    public function pulseWave(\Grafana\Foundation\Cog\Builder $pulseWave): static
    {
        $pulseWaveResource = $pulseWave->build();
        $this->internal->pulseWave = $pulseWaveResource;
    
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
    public function labels(string $labels): static
    {
        $this->internal->labels = $labels;
    
        return $this;
    }
    public function lines(int $lines): static
    {
        $this->internal->lines = $lines;
    
        return $this;
    }
    public function levelColumn(bool $levelColumn): static
    {
        $this->internal->levelColumn = $levelColumn;
    
        return $this;
    }
    public function channel(string $channel): static
    {
        $this->internal->channel = $channel;
    
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
    public function csvFileName(string $csvFileName): static
    {
        $this->internal->csvFileName = $csvFileName;
    
        return $this;
    }
    public function csvContent(string $csvContent): static
    {
        $this->internal->csvContent = $csvContent;
    
        return $this;
    }
    public function rawFrameContent(string $rawFrameContent): static
    {
        $this->internal->rawFrameContent = $rawFrameContent;
    
        return $this;
    }
    public function seriesCount(int $seriesCount): static
    {
        $this->internal->seriesCount = $seriesCount;
    
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
    public function errorType(\Grafana\Foundation\Testdata\DataqueryErrorType $errorType): static
    {
        $this->internal->errorType = $errorType;
    
        return $this;
    }
    public function spanCount(int $spanCount): static
    {
        $this->internal->spanCount = $spanCount;
    
        return $this;
    }
    /**
     * @param array<array<string|int>> $points
     */
    public function points(array $points): static
    {
        $this->internal->points = $points;
    
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
