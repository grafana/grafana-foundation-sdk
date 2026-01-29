<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * A library panel is a reusable panel that you can use in any dashboard.
 * When you make a change to a library panel, that change propagates to all instances of where the panel is used.
 * Library panels streamline reuse of panels across multiple dashboards.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\LibraryPanelRef>
 */
class LibraryPanelRefBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\LibraryPanelRef $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\LibraryPanelRef();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\LibraryPanelRef
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Library panel name
     */
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }

    /**
     * Library panel uid
     */
    public function uid(string $uid): static
    {
        $this->internal->uid = $uid;
    
        return $this;
    }

}
