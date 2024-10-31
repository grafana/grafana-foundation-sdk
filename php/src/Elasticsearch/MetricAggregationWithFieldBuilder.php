<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\MetricAggregationWithField>
 */
class MetricAggregationWithFieldBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\MetricAggregationWithField $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\MetricAggregationWithField();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\MetricAggregationWithField
     */
    public function build()
    {
        return $this->internal;
    }

    public function field(string $field): static
    {
        $this->internal->field = $field;
    
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
