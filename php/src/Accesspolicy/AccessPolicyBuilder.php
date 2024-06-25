<?php

namespace Grafana\Foundation\Accesspolicy;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Accesspolicy\AccessPolicy>
 */
class AccessPolicyBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Accesspolicy\AccessPolicy $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Accesspolicy\AccessPolicy();
    }

    /**
     * @return \Grafana\Foundation\Accesspolicy\AccessPolicy
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The scope where these policies should apply
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Accesspolicy\ResourceRef> $scope
     */
    public function scope(\Grafana\Foundation\Cog\Builder $scope): static
    {
        $scopeResource = $scope->build();
        $this->internal->scope = $scopeResource;
    
        return $this;
    }
    /**
     * The role that must apply this policy
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Accesspolicy\RoleRef> $role
     */
    public function role(\Grafana\Foundation\Cog\Builder $role): static
    {
        $roleResource = $role->build();
        $this->internal->role = $roleResource;
    
        return $this;
    }
    /**
     * The set of rules to apply.  Note that * is required to modify
     * access policy rules, and that "none" will reject all actions
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Accesspolicy\AccessRule> $rules
     */
    public function rules(\Grafana\Foundation\Cog\Builder $rules): static
    {
        $rulesResource = $rules->build();
        $this->internal->rules[] = $rulesResource;
    
        return $this;
    }

}
