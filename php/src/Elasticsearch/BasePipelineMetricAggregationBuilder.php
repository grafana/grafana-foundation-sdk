<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\BasePipelineMetricAggregation>
 */
class BasePipelineMetricAggregationBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\BasePipelineMetricAggregation $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\BasePipelineMetricAggregation();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\BasePipelineMetricAggregation
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
    public function type(string $type): static
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
