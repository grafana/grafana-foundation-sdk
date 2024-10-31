<?php

namespace Grafana\Foundation\Rolebinding;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Rolebinding\RoleBindingSubject>
 */
class RoleBindingSubjectBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Rolebinding\RoleBindingSubject $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Rolebinding\RoleBindingSubject();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Rolebinding\RoleBindingSubject
     */
    public function build()
    {
        return $this->internal;
    }

    public function kind(\Grafana\Foundation\Rolebinding\RoleBindingSubjectKind $kind): static
    {
        $this->internal->kind = $kind;
    
        return $this;
    }
    /**
     * The team/user identifier name
     */
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }

}
