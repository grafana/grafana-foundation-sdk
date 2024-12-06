<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\BaseMetricAggregation>
 */
class BaseMetricAggregationBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\BaseMetricAggregation $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\BaseMetricAggregation();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\BaseMetricAggregation
     */
    public function build()
    {
        return $this->internal;
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
