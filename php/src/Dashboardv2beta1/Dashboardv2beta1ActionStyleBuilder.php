<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1ActionStyle>
 */
class Dashboardv2beta1ActionStyleBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1ActionStyle $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1ActionStyle();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1ActionStyle
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
