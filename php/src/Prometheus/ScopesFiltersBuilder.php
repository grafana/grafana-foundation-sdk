<?php

namespace Grafana\Foundation\Prometheus;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\ScopesFilters>
 */
class ScopesFiltersBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Prometheus\ScopesFilters $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Prometheus\ScopesFilters();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Prometheus\ScopesFilters
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

}
