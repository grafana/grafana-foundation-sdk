<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\Dashboardv2ActionStyle>
 */
class Dashboardv2ActionStyleBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\Dashboardv2ActionStyle $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\Dashboardv2ActionStyle();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\Dashboardv2ActionStyle
     */
    public function build()
    {
        return $this->internal;
    }

    public function backgroundColor(string $backgroundColor): static
    {
        $this->internal->backgroundColor = $backgroundColor;
    
        return $this;
    }

}
