<?php

namespace Grafana\Foundation\Debug;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Debug\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Debug\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Debug\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Debug\Options
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(\Grafana\Foundation\Debug\DebugMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Debug\UpdateConfig> $counters
     */
    public function counters(\Grafana\Foundation\Cog\Builder $counters): static
    {
        $countersResource = $counters->build();
        $this->internal->counters = $countersResource;
    
        return $this;
    }

}
