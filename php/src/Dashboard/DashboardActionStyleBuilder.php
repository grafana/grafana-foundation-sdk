<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardActionStyle>
 */
class DashboardActionStyleBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\DashboardActionStyle $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\DashboardActionStyle();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\DashboardActionStyle
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
