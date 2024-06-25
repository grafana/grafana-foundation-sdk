<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\BucketAggregationWithField>
 */
class BucketAggregationWithFieldBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\BucketAggregationWithField $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\BucketAggregationWithField();
    }

    /**
     * @return \Grafana\Foundation\Elasticsearch\BucketAggregationWithField
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
    public function id(string $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    public function type(\Grafana\Foundation\Elasticsearch\BucketAggregationType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    /**
     * @param mixed $settings
     */
    public function settings( $settings): static
    {
        $this->internal->settings = $settings;
    
        return $this;
    }

}
