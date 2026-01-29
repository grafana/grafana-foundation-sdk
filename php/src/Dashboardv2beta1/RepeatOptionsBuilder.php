<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RepeatOptions>
 */
class RepeatOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\RepeatOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\RepeatOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\RepeatOptions
     */
    public function build()
    {
        return $this->internal;
    }

    public function value(string $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }

    public function direction(\Grafana\Foundation\Dashboardv2beta1\RepeatOptionsDirection $direction): static
    {
        $this->internal->direction = $direction;
    
        return $this;
    }

    public function maxPerRow(int $maxPerRow): static
    {
        $this->internal->maxPerRow = $maxPerRow;
    
        return $this;
    }

}
