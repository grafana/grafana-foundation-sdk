<?php

namespace Grafana\Foundation\Starsv1alpha1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Starsv1alpha1\Resource>
 */
class ResourceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Starsv1alpha1\Resource $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Starsv1alpha1\Resource();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Starsv1alpha1\Resource
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

    public function kind(string $kind): static
    {
        $this->internal->kind = $kind;
    
        return $this;
    }

    /**
     * The set of resources
     * +listType=set
     * @param array<string> $names
     */
    public function names(array $names): static
    {
        $this->internal->names = $names;
    
        return $this;
    }

}
