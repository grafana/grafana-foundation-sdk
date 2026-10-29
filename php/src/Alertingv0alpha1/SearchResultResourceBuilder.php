<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * #SearchResultResource is the full identity of a hit.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\SearchResultResource>
 */
class SearchResultResourceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv0alpha1\SearchResultResource $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv0alpha1\SearchResultResource();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv0alpha1\SearchResultResource
     */
    public function build()
    {
        return $this->internal;
    }

    public function group(string $group): static
    {
        $this->internal->group = $group;
    
        return $this;
    }

    public function resource(string $resource): static
    {
        $this->internal->resource = $resource;
    
        return $this;
    }

    public function kind(string $kind): static
    {
        $this->internal->kind = $kind;
    
        return $this;
    }

    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }

}
