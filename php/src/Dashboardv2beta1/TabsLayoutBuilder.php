<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind>
 */
class TabsLayoutBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind();
    $this->internal->kind = "TabsLayout";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabKind>> $tabs
     */
    public function tabs(array $tabs): static
    {
            $tabsResources = [];
            foreach ($tabs as $r1) {
                    $tabsResources[] = $r1->build();
            }
        $this->internal->spec->tabs = $tabsResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabKind> $tab
     */
    public function tab(\Grafana\Foundation\Cog\Builder $tab): static
    {
        $tabResource = $tab->build();
        $this->internal->spec->tabs[] = $tabResource;
    
        return $this;
    }

}
