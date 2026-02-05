<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabSpec>
 */
class TabsLayoutTabSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabSpec
     */
    public function build()
    {
        return $this->internal;
    }

    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind> $layout
     */
    public function layout( $layout): static
    {
        $layoutResource = $layout->build();
        $this->internal->layout = $layoutResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind> $conditionalRendering
     */
    public function conditionalRendering(\Grafana\Foundation\Cog\Builder $conditionalRendering): static
    {
        $conditionalRenderingResource = $conditionalRendering->build();
        $this->internal->conditionalRendering = $conditionalRenderingResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabRepeatOptions> $repeat
     */
    public function repeat(\Grafana\Foundation\Cog\Builder $repeat): static
    {
        $repeatResource = $repeat->build();
        $this->internal->repeat = $repeatResource;
    
        return $this;
    }

}
