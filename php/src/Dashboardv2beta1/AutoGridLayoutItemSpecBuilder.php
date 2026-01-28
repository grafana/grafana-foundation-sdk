<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemSpec>
 */
class AutoGridLayoutItemSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemSpec
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ElementReference> $element
     */
    public function element(\Grafana\Foundation\Cog\Builder $element): static
    {
        $elementResource = $element->build();
        $this->internal->element = $elementResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptions> $repeat
     */
    public function repeat(\Grafana\Foundation\Cog\Builder $repeat): static
    {
        $repeatResource = $repeat->build();
        $this->internal->repeat = $repeatResource;
    
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

}
