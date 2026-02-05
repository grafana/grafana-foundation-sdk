<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Constant variable kind
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind>
 */
class ConstantVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind $internal;

    public function __construct(string $name)
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind();
    $this->internal->kind = "ConstantVariable";
    $this->internal->spec->name = $name;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function spec(\Grafana\Foundation\Dashboardv2beta1\ConstantVariableSpec $spec): static
    {
        $this->internal->spec = $spec;
    
        return $this;
    }

    public function name(string $name): static
    {
        $this->internal->spec->name = $name;
    
        return $this;
    }

    public function query(string $query): static
    {
        $this->internal->spec->query = $query;
    
        return $this;
    }

    public function current(\Grafana\Foundation\Dashboardv2beta1\VariableOption $current): static
    {
        $this->internal->spec->current = $current;
    
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
