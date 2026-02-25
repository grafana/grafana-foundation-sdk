<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind>
 */
class RowBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind();
    $this->internal->kind = "RowsLayoutRow";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function title(string $title): static
    {
        $this->internal->spec->title = $title;
    
        return $this;
    }

    public function collapse(bool $collapse): static
    {
        $this->internal->spec->collapse = $collapse;
    
        return $this;
    }

    public function hideHeader(bool $hideHeader): static
    {
        $this->internal->spec->hideHeader = $hideHeader;
    
        return $this;
    }

    public function fillScreen(bool $fillScreen): static
    {
        $this->internal->spec->fillScreen = $fillScreen;
    
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

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowRepeatOptions> $repeat
     */
    public function repeat(\Grafana\Foundation\Cog\Builder $repeat): static
    {
        $repeatResource = $repeat->build();
        $this->internal->spec->repeat = $repeatResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind> $layout
     */
    public function layout( $layout): static
    {
        $layoutResource = $layout->build();
        $this->internal->spec->layout = $layoutResource;
    
        return $this;
    }

}
