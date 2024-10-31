<?php

namespace Grafana\Foundation\Accesspolicy;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Accesspolicy\AccessRule>
 */
class AccessRuleBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Accesspolicy\AccessRule $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Accesspolicy\AccessRule();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Accesspolicy\AccessRule
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The kind this rule applies to (dashboards, alert, etc)
     */
    public function kind(string $kind): static
    {
        $this->internal->kind = $kind;
    
        return $this;
    }
    /**
     * READ, WRITE, CREATE, DELETE, ...
     * should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
     * @param string $verb
     */
    public function verb( $verb): static
    {
        $this->internal->verb = $verb;
    
        return $this;
    }
    /**
     * Specific sub-elements like "alert.rules" or "dashboard.permissions"????
     */
    public function target(string $target): static
    {
        $this->internal->target = $target;
    
        return $this;
    }

}
