<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Query variable kind
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\QueryVariableKind>
 */
class QueryVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\QueryVariableKind $internal;

    public function __construct(string $name)
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\QueryVariableKind();
    $this->internal->kind = "QueryVariable";
    $this->internal->spec->name = $name;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\QueryVariableKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function spec(\Grafana\Foundation\Dashboardv2beta1\QueryVariableSpec $spec): static
    {
        $this->internal->spec = $spec;
    
        return $this;
    }

    public function name(string $name): static
    {
        $this->internal->spec->name = $name;
    
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

    public function refresh(\Grafana\Foundation\Dashboardv2beta1\VariableRefresh $refresh): static
    {
        $this->internal->spec->refresh = $refresh;
    
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

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DataQueryKind> $query
     */
    public function query(\Grafana\Foundation\Cog\Builder $query): static
    {
        $queryResource = $query->build();
        $this->internal->spec->query = $queryResource;
    
        return $this;
    }

    public function regex(string $regex): static
    {
        $this->internal->spec->regex = $regex;
    
        return $this;
    }

    public function regexApplyTo(\Grafana\Foundation\Dashboardv2beta1\VariableRegexApplyTo $regexApplyTo): static
    {
        $this->internal->spec->regexApplyTo = $regexApplyTo;
    
        return $this;
    }

    public function sort(\Grafana\Foundation\Dashboardv2beta1\VariableSort $sort): static
    {
        $this->internal->spec->sort = $sort;
    
        return $this;
    }

    public function definition(string $definition): static
    {
        $this->internal->spec->definition = $definition;
    
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

    public function includeAll(bool $includeAll): static
    {
        $this->internal->spec->includeAll = $includeAll;
    
        return $this;
    }

    public function allValue(string $allValue): static
    {
        $this->internal->spec->allValue = $allValue;
    
        return $this;
    }

    public function placeholder(string $placeholder): static
    {
        $this->internal->spec->placeholder = $placeholder;
    
        return $this;
    }

    public function allowCustomValue(bool $allowCustomValue): static
    {
        $this->internal->spec->allowCustomValue = $allowCustomValue;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Dashboardv2beta1\VariableOption> $staticOptions
     */
    public function staticOptions(array $staticOptions): static
    {
        $this->internal->spec->staticOptions = $staticOptions;
    
        return $this;
    }

    public function staticOptionsOrder(\Grafana\Foundation\Dashboardv2beta1\QueryVariableSpecStaticOptionsOrder $staticOptionsOrder): static
    {
        $this->internal->spec->staticOptionsOrder = $staticOptionsOrder;
    
        return $this;
    }

}
