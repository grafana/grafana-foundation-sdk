<?php

namespace Grafana\Foundation\Nodegraph;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Nodegraph\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Nodegraph\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Nodegraph\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Nodegraph\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Nodegraph\NodeOptions> $nodes
     */
    public function nodes(\Grafana\Foundation\Cog\Builder $nodes): static
    {
        $nodesResource = $nodes->build();
        $this->internal->nodes = $nodesResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Nodegraph\EdgeOptions> $edges
     */
    public function edges(\Grafana\Foundation\Cog\Builder $edges): static
    {
        $edgesResource = $edges->build();
        $this->internal->edges = $edgesResource;
    
        return $this;
    }

}
