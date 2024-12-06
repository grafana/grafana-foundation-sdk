<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ReduceDataOptions>
 */
class ReduceDataOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\ReduceDataOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\ReduceDataOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\ReduceDataOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * If true show each row value
     */
    public function values(bool $values): static
    {
        $this->internal->values = $values;
    
        return $this;
    }
    /**
     * if showing all values limit
     */
    public function limit(float $limit): static
    {
        $this->internal->limit = $limit;
    
        return $this;
    }
    /**
     * When !values, pick one value for the whole field
     * @param array<string> $calcs
     */
    public function calcs(array $calcs): static
    {
        $this->internal->calcs = $calcs;
    
        return $this;
    }
    /**
     * Which fields to show.  By default this is only numeric fields
     */
    public function fields(string $fields): static
    {
        $this->internal->fields = $fields;
    
        return $this;
    }

}
