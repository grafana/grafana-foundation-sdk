<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemKind>
 */
class AutoGridLayoutItemBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemKind $internal;

    public function __construct(string $name)
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemKind();
    $this->internal->kind = "AutoGridLayoutItem";
    $this->internal->spec->element->kind = "ElementReference";
    $this->internal->spec->element->name = $name;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemKind
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
        $this->internal->spec->element = $elementResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptions> $repeat
     */
    public function repeat(\Grafana\Foundation\Cog\Builder $repeat): static
    {
        $repeatResource = $repeat->build();
        $this->internal->spec->repeat = $repeatResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind> $conditionalRendering
     */
    public function conditionalRendering(\Grafana\Foundation\Cog\Builder $conditionalRendering): static
    {
        $conditionalRenderingResource = $conditionalRendering->build();
        $this->internal->spec->conditionalRendering = $conditionalRenderingResource;
    
        return $this;
    }

    public function name(string $name): static
    {
        $this->internal->spec->element->name = $name;
    
        return $this;
    }

}
