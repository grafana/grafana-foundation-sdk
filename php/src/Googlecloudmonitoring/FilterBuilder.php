<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * Query filter representation.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\Filter>
 */
class FilterBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Googlecloudmonitoring\Filter $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Googlecloudmonitoring\Filter();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Googlecloudmonitoring\Filter
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Filter key.
     */
    public function key(string $key): static
    {
        $this->internal->key = $key;
    
        return $this;
    }
    /**
     * Filter operator.
     */
    public function operator(string $operator): static
    {
        $this->internal->operator = $operator;
    
        return $this;
    }
    /**
     * Filter value.
     */
    public function value(string $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }
    /**
     * Filter condition.
     */
    public function condition(string $condition): static
    {
        $this->internal->condition = $condition;
    
        return $this;
    }

}
