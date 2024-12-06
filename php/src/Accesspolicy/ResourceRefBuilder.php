<?php

namespace Grafana\Foundation\Accesspolicy;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Accesspolicy\ResourceRef>
 */
class ResourceRefBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Accesspolicy\ResourceRef $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Accesspolicy\ResourceRef();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Accesspolicy\ResourceRef
     */
    public function build()
    {
        return $this->internal;
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
