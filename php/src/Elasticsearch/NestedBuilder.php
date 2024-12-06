<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Nested>
 */
class NestedBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\Nested $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\Nested();
    $this->internal->type = "nested";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\Nested
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
    /**
     * @param mixed $settings
     */
    public function settings( $settings): static
    {
        $this->internal->settings = $settings;
    
        return $this;
    }

}
