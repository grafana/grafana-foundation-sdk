<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\CumulativeSum>
 */
class CumulativeSumBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\CumulativeSum $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\CumulativeSum();
    $this->internal->type = "cumulative_sum";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\CumulativeSum
     */
    public function build()
    {
        return $this->internal;
    }

    public function pipelineAgg(string $pipelineAgg): static
    {
        $this->internal->pipelineAgg = $pipelineAgg;
    
        return $this;
    }
    public function field(string $field): static
    {
        $this->internal->field = $field;
    
        return $this;
    }
    public function id(string $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchCumulativeSumSettings> $settings
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
