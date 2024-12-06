<?php

namespace Grafana\Foundation\Nodegraph;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Nodegraph\EdgeOptions>
 */
class EdgeOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Nodegraph\EdgeOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Nodegraph\EdgeOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Nodegraph\EdgeOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Unit for the main stat to override what ever is set in the data frame.
     */
    public function mainStatUnit(string $mainStatUnit): static
    {
        $this->internal->mainStatUnit = $mainStatUnit;
    
        return $this;
    }
    /**
     * Unit for the secondary stat to override what ever is set in the data frame.
     */
    public function secondaryStatUnit(string $secondaryStatUnit): static
    {
        $this->internal->secondaryStatUnit = $secondaryStatUnit;
    
        return $this;
    }

}
