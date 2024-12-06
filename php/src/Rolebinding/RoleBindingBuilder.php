<?php

namespace Grafana\Foundation\Rolebinding;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Rolebinding\RoleBinding>
 */
class RoleBindingBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Rolebinding\RoleBinding $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Rolebinding\RoleBinding();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Rolebinding\RoleBinding
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The role we are discussing
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Rolebinding\BuiltinRoleRef>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Rolebinding\CustomRoleRef> $role
     */
    public function role( $role): static
    {
        $roleResource = $role->build();
        $this->internal->role = $roleResource;
    
        return $this;
    }
    /**
     * The team or user that has the specified role
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Rolebinding\RoleBindingSubject> $subject
     */
    public function subject(\Grafana\Foundation\Cog\Builder $subject): static
    {
        $subjectResource = $subject->build();
        $this->internal->subject = $subjectResource;
    
        return $this;
    }

}
