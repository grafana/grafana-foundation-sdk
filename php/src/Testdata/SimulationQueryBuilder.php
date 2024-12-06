<?php

namespace Grafana\Foundation\Testdata;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\SimulationQuery>
 */
class SimulationQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Testdata\SimulationQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Testdata\SimulationQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Testdata\SimulationQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param mixed $config
     */
    public function config( $config): static
    {
        $this->internal->config = $config;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\Key> $key
     */
    public function key(\Grafana\Foundation\Cog\Builder $key): static
    {
        $keyResource = $key->build();
        $this->internal->key = $keyResource;
    
        return $this;
    }
    public function last(bool $last): static
    {
        $this->internal->last = $last;
    
        return $this;
    }
    public function stream(bool $stream): static
    {
        $this->internal->stream = $stream;
    
        return $this;
    }

}
