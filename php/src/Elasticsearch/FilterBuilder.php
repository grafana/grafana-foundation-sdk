<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Filter>
 */
class FilterBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\Filter $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\Filter();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\Filter
     */
    public function build()
    {
        return $this->internal;
    }

    public function query(string $query): static
    {
        $this->internal->query = $query;
    
        return $this;
    }
    public function label(string $label): static
    {
        $this->internal->label = $label;
    
        return $this;
    }

}
