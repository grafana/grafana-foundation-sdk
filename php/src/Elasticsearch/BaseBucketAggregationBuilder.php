<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\BaseBucketAggregation>
 */
class BaseBucketAggregationBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\BaseBucketAggregation $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\BaseBucketAggregation();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\BaseBucketAggregation
     */
    public function build()
    {
        return $this->internal;
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
