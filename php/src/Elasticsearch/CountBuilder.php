<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Count>
 */
class CountBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\Count $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\Count();
    $this->internal->type = "count";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\Count
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
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }

}
