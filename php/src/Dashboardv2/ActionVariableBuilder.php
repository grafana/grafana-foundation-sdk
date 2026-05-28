<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ActionVariable>
 */
class ActionVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\ActionVariable $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\ActionVariable();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\ActionVariable
     */
    public function build()
    {
        return $this->internal;
    }

    public function key(string $key): static
    {
        $this->internal->key = $key;
    
        return $this;
    }

    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }

}
