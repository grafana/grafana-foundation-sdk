<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\ActionVariable>
 */
class ActionVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\ActionVariable $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\ActionVariable();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\ActionVariable
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
