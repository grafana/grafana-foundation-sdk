<?php

namespace Grafana\Foundation\Role;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Role\Role>
 */
class RoleBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Role\Role $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Role\Role();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Role\Role
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The role identifier `managed:builtins:editor:permissions`
     */
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    /**
     * Optional display
     */
    public function displayName(string $displayName): static
    {
        $this->internal->displayName = $displayName;
    
        return $this;
    }
    /**
     * Name of the team.
     */
    public function groupName(string $groupName): static
    {
        $this->internal->groupName = $groupName;
    
        return $this;
    }
    /**
     * Role description
     */
    public function description(string $description): static
    {
        $this->internal->description = $description;
    
        return $this;
    }
    /**
     * Do not show this role
     */
    public function hidden(bool $hidden): static
    {
        $this->internal->hidden = $hidden;
    
        return $this;
    }

}
