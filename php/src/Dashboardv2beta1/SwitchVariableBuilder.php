<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind>
 */
class SwitchVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind $internal;

    public function __construct(string $name)
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind();
    $this->internal->kind = "SwitchVariable";
    $this->internal->spec->name = $name;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function name(string $name): static
    {
        $this->internal->spec->name = $name;
    
        return $this;
    }

    public function current(string $current): static
    {
        $this->internal->spec->current = $current;
    
        return $this;
    }

    public function enabledValue(string $enabledValue): static
    {
        $this->internal->spec->enabledValue = $enabledValue;
    
        return $this;
    }

    public function disabledValue(string $disabledValue): static
    {
        $this->internal->spec->disabledValue = $disabledValue;
    
        return $this;
    }

    public function label(string $label): static
    {
        $this->internal->spec->label = $label;
    
        return $this;
    }

    public function hide(\Grafana\Foundation\Dashboardv2beta1\VariableHide $hide): static
    {
        $this->internal->spec->hide = $hide;
    
        return $this;
    }

    public function skipUrlSync(bool $skipUrlSync): static
    {
        $this->internal->spec->skipUrlSync = $skipUrlSync;
    
        return $this;
    }

    public function description(string $description): static
    {
        $this->internal->spec->description = $description;
    
        return $this;
    }

}
