<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\SelectableValue>
 */
class SelectableValueBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\SelectableValue $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\SelectableValue();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\SelectableValue
     */
    public function build()
    {
        return $this->internal;
    }

    public function label(string $label): static
    {
        $this->internal->label = $label;
    
        return $this;
    }

    public function value(string $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }

}
