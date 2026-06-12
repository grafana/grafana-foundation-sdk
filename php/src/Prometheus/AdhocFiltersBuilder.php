<?php

namespace Grafana\Foundation\Prometheus;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\AdhocFilters>
 */
class AdhocFiltersBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Prometheus\AdhocFilters $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Prometheus\AdhocFilters();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Prometheus\AdhocFilters
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
