<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * #MovingAverage's settings are overridden in types.ts
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\MovingAverage>
 */
class MovingAverageBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\MovingAverage $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\MovingAverage();
    $this->internal->type = "moving_avg";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\MovingAverage
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
     * @param array<string, mixed> $settings
     */
    public function settings(array $settings): static
    {
        $this->internal->settings = $settings;
    
        return $this;
    }
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }

}
