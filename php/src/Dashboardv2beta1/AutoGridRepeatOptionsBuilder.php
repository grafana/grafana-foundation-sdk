<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptions>
 */
class AutoGridRepeatOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptions
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

}
