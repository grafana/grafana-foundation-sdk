<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\PipelineMetricAggregationWithMultipleBucketPaths>
 */
class PipelineMetricAggregationWithMultipleBucketPathsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\PipelineMetricAggregationWithMultipleBucketPaths $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\PipelineMetricAggregationWithMultipleBucketPaths();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\PipelineMetricAggregationWithMultipleBucketPaths
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\PipelineVariable>> $pipelineVariables
     */
    public function pipelineVariables(array $pipelineVariables): static
    {
            $pipelineVariablesResources = [];
            foreach ($pipelineVariables as $r1) {
                    $pipelineVariablesResources[] = $r1->build();
            }
        $this->internal->pipelineVariables = $pipelineVariablesResources;
    
        return $this;
    }
    /**
     * @param string|\Grafana\Foundation\Elasticsearch\PipelineMetricAggregationType $type
     */
    public function type( $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    public function id(string $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }

}
