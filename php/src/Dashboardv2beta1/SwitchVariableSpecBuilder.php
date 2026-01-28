<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\SwitchVariableSpec>
 */
class SwitchVariableSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\SwitchVariableSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\SwitchVariableSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\SwitchVariableSpec
     */
    public function build()
    {
        return $this->internal;
    }

    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }

    public function current(string $current): static
    {
        $this->internal->current = $current;
    
        return $this;
    }

    public function enabledValue(string $enabledValue): static
    {
        $this->internal->enabledValue = $enabledValue;
    
        return $this;
    }

    public function disabledValue(string $disabledValue): static
    {
        $this->internal->disabledValue = $disabledValue;
    
        return $this;
    }

    public function label(string $label): static
    {
        $this->internal->label = $label;
    
        return $this;
    }

    public function hide(\Grafana\Foundation\Dashboardv2beta1\VariableHide $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }

    public function skipUrlSync(bool $skipUrlSync): static
    {
        $this->internal->skipUrlSync = $skipUrlSync;
    
        return $this;
    }

    public function description(string $description): static
    {
        $this->internal->description = $description;
    
        return $this;
    }

}
