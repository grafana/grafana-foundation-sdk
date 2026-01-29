<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Define the AdHocFilterWithLabels type
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabels>
 */
class AdHocFilterWithLabelsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabels $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabels();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabels
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

    public function operator(string $operator): static
    {
        $this->internal->operator = $operator;
    
        return $this;
    }

    public function value(string $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }

    /**
     * @param array<string> $values
     */
    public function values(array $values): static
    {
        $this->internal->values = $values;
    
        return $this;
    }

    public function keyLabel(string $keyLabel): static
    {
        $this->internal->keyLabel = $keyLabel;
    
        return $this;
    }

    /**
     * @param array<string> $valueLabels
     */
    public function valueLabels(array $valueLabels): static
    {
        $this->internal->valueLabels = $valueLabels;
    
        return $this;
    }

    public function forceEdit(bool $forceEdit): static
    {
        $this->internal->forceEdit = $forceEdit;
    
        return $this;
    }

    public function origin(string $origin): static
    {
        $this->internal->origin = $origin;
    
        return $this;
    }

    /**
     * @deprecated
     */
    public function condition(string $condition): static
    {
        $this->internal->condition = $condition;
    
        return $this;
    }

}
