<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\BucketScript>
 */
class BucketScriptBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\BucketScript $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\BucketScript();
    $this->internal->type = "bucket_script";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\BucketScript
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
    public function id(string $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchBucketScriptSettings> $settings
     */
    public function settings(\Grafana\Foundation\Cog\Builder $settings): static
    {
        $settingsResource = $settings->build();
        $this->internal->settings = $settingsResource;
    
        return $this;
    }
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }

}
