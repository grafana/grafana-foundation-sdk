<?php

namespace Grafana\Foundation\Starsv1alpha1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Starsv1alpha1\Stars>
 */
class StarsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Starsv1alpha1\Stars $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Starsv1alpha1\Stars();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Starsv1alpha1\Stars
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Starsv1alpha1\Resource>> $resource
     */
    public function resources(array $resource): static
    {
            $resourceResources = [];
            foreach ($resource as $r1) {
                    $resourceResources[] = $r1->build();
            }
        $this->internal->resource = $resourceResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Starsv1alpha1\Resource> $resource
     */
    public function resource(\Grafana\Foundation\Cog\Builder $resource): static
    {
        $resourceResource = $resource->build();
        $this->internal->resource[] = $resourceResource;
    
        return $this;
    }

}
