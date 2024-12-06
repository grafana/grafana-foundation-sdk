<?php

namespace Grafana\Foundation\Accesspolicy;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Accesspolicy\RoleRef>
 */
class RoleRefBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Accesspolicy\RoleRef $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Accesspolicy\RoleRef();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Accesspolicy\RoleRef
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Policies can apply to roles, teams, or users
     * Applying policies to individual users is supported, but discouraged
     */
    public function kind(\Grafana\Foundation\Accesspolicy\RoleRefKind $kind): static
    {
        $this->internal->kind = $kind;
    
        return $this;
    }
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    public function xname(string $xname): static
    {
        $this->internal->xname = $xname;
    
        return $this;
    }

}
