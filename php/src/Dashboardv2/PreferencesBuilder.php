<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * Dashboard specific preferences (applied per dashboard = all users using the dashboard)
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\Preferences>
 */
class PreferencesBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\Preferences $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\Preferences();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\Preferences
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * default layout template to be used when new containers are created
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\AutoGridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\GridLayoutKind> $layout
     */
    public function layout( $layout): static
    {
        $layoutResource = $layout->build();
        $this->internal->layout = $layoutResource;
    
        return $this;
    }

}
