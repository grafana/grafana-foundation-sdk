<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Adhoc variable kind
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind>
 */
class AdhocVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind $internal;

    public function __construct(string $name)
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind();
    $this->internal->kind = "AdhocVariable";
    $this->internal->spec->name = $name;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind
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
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1AdhocVariableKindDatasource> $datasource
     */
    public function datasource(\Grafana\Foundation\Cog\Builder $datasource): static
    {
        $datasourceResource = $datasource->build();
        $this->internal->datasource = $datasourceResource;
    
        return $this;
    }

    public function spec(\Grafana\Foundation\Dashboardv2beta1\AdhocVariableSpec $spec): static
    {
        $this->internal->spec = $spec;
    
        return $this;
    }

    public function name(string $name): static
    {
        $this->internal->spec->name = $name;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabels>> $baseFilters
     */
    public function baseFilters(array $baseFilters): static
    {
            $baseFiltersResources = [];
            foreach ($baseFilters as $r1) {
                    $baseFiltersResources[] = $r1->build();
            }
        $this->internal->spec->baseFilters = $baseFiltersResources;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabels>> $filters
     */
    public function filters(array $filters): static
    {
            $filtersResources = [];
            foreach ($filters as $r1) {
                    $filtersResources[] = $r1->build();
            }
        $this->internal->spec->filters = $filtersResources;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\MetricFindValue>> $defaultKeys
     */
    public function defaultKeys(array $defaultKeys): static
    {
            $defaultKeysResources = [];
            foreach ($defaultKeys as $r1) {
                    $defaultKeysResources[] = $r1->build();
            }
        $this->internal->spec->defaultKeys = $defaultKeysResources;
    
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
