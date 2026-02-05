<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutSpec>
 */
class TabsLayoutSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\TabsLayoutSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\TabsLayoutSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\TabsLayoutSpec
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
        $this->internal->tabs = $tabsResources;
    
        return $this;
    }

}
