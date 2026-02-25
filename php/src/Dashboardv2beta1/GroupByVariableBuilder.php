<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Group variable kind
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind>
 */
class GroupByVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind $internal;

    public function __construct(string $name)
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind();
    $this->internal->kind = "GroupByVariable";
    $this->internal->spec->name = $name;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function group(string $group): static
    {
        $this->internal->group = $group;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1GroupByVariableKindDatasource> $datasource
     */
    public function datasource(\Grafana\Foundation\Cog\Builder $datasource): static
    {
        $datasourceResource = $datasource->build();
        $this->internal->datasource = $datasourceResource;
    
        return $this;
    }

    public function name(string $name): static
    {
        $this->internal->spec->name = $name;
    
        return $this;
    }

    public function defaultValue(\Grafana\Foundation\Dashboardv2beta1\VariableOption $defaultValue): static
    {
        $this->internal->spec->defaultValue = $defaultValue;
    
        return $this;
    }

    public function current(\Grafana\Foundation\Dashboardv2beta1\VariableOption $current): static
    {
        $this->internal->spec->current = $current;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Dashboardv2beta1\VariableOption> $options
     */
    public function options(array $options): static
    {
        $this->internal->spec->options = $options;
    
        return $this;
    }

    public function multi(bool $multi): static
    {
        $this->internal->spec->multi = $multi;
    
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
