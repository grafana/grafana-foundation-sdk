<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\NamedRoutingTree>
 */
class NamedRoutingTreeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv0alpha1\NamedRoutingTree $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv0alpha1\NamedRoutingTree();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv0alpha1\NamedRoutingTree
     */
    public function build()
    {
        return $this->internal;
    }

    public function routingTree(string $routingTree): static
    {
        $this->internal->routingTree = $routingTree;
    
        return $this;
    }

}
