<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Datasource variable kind
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind>
 */
class DatasourceVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind $internal;

    public function __construct(string $name)
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind();
    $this->internal->kind = "DatasourceVariable";
    $this->internal->spec->name = $name;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function spec(\Grafana\Foundation\Dashboardv2beta1\DatasourceVariableSpec $spec): static
    {
        $this->internal->spec = $spec;
    
        return $this;
    }

    public function name(string $name): static
    {
        $this->internal->spec->name = $name;
    
        return $this;
    }

    public function pluginId(string $pluginId): static
    {
        $this->internal->spec->pluginId = $pluginId;
    
        return $this;
    }

    public function refresh(\Grafana\Foundation\Dashboardv2beta1\VariableRefresh $refresh): static
    {
        $this->internal->spec->refresh = $refresh;
    
        return $this;
    }

    public function regex(string $regex): static
    {
        $this->internal->spec->regex = $regex;
    
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

    public function allowCustomValue(bool $allowCustomValue): static
    {
        $this->internal->spec->allowCustomValue = $allowCustomValue;
    
        return $this;
    }

}
