<?php

namespace Grafana\Foundation\Testdata;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\NodesQuery>
 */
class NodesQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Testdata\NodesQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Testdata\NodesQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Testdata\NodesQuery
     */
    public function build()
    {
        return $this->internal;
    }

    public function type(\Grafana\Foundation\Testdata\NodesQueryType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    public function count(int $count): static
    {
        $this->internal->count = $count;
    
        return $this;
    }
    public function seed(int $seed): static
    {
        $this->internal->seed = $seed;
    
        return $this;
    }

}
