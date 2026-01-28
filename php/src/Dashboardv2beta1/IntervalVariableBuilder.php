<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Interval variable kind
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind>
 */
class IntervalVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind $internal;

    public function __construct(string $name)
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind();
    $this->internal->kind = "IntervalVariable";
    $this->internal->spec->refresh = "onTimeRangeChanged";
    $this->internal->spec->name = $name;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function spec(\Grafana\Foundation\Dashboardv2beta1\IntervalVariableSpec $spec): static
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

    /**
     * @param array<\Grafana\Foundation\Dashboardv2beta1\VariableOption> $options
     */
    public function options(array $options): static
    {
        $this->internal->spec->options = $options;
    
        return $this;
    }

    public function auto(bool $auto): static
    {
        $this->internal->spec->auto = $auto;
    
        return $this;
    }

    public function autoMin(string $autoMin): static
    {
        $this->internal->spec->autoMin = $autoMin;
    
        return $this;
    }

    public function autoCount(int $autoCount): static
    {
        $this->internal->spec->autoCount = $autoCount;
    
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
