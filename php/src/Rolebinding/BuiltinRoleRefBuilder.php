<?php

namespace Grafana\Foundation\Rolebinding;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Rolebinding\BuiltinRoleRef>
 */
class BuiltinRoleRefBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Rolebinding\BuiltinRoleRef $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Rolebinding\BuiltinRoleRef();
    $this->internal->kind = "BuiltinRole";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Rolebinding\BuiltinRoleRef
     */
    public function build()
    {
        return $this->internal;
    }

    public function name(\Grafana\Foundation\Rolebinding\BuiltinRoleRefName $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }

}
