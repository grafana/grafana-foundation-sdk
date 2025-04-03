<?php

namespace Grafana\Foundation\Rolebinding;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Rolebinding\CustomRoleRef>
 */
class CustomRoleRefBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Rolebinding\CustomRoleRef $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Rolebinding\CustomRoleRef();
    $this->internal->kind = "Role";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Rolebinding\CustomRoleRef
     */
    public function build()
    {
        return $this->internal;
    }

    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }

}
