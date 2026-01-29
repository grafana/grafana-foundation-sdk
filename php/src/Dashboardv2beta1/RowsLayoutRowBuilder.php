<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind>
 */
class RowsLayoutRowBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind $internal;

    public function __construct(string $title)
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind();
    $this->internal->kind = "RowsLayoutRow";
    $this->internal->spec->title = $title;
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
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GridLayoutKind> $gridLayoutKind
     */
    public function gridLayout(\Grafana\Foundation\Cog\Builder $gridLayoutKind): static
    {
        $gridLayoutKindResource = $gridLayoutKind->build();
        $this->internal->spec->layout = $gridLayoutKindResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind> $autoGridLayoutKind
     */
    public function autoGridLayout(\Grafana\Foundation\Cog\Builder $autoGridLayoutKind): static
    {
        $autoGridLayoutKindResource = $autoGridLayoutKind->build();
        $this->internal->spec->layout = $autoGridLayoutKindResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind> $tabsLayoutKind
     */
    public function tabsLayout(\Grafana\Foundation\Cog\Builder $tabsLayoutKind): static
    {
        $tabsLayoutKindResource = $tabsLayoutKind->build();
        $this->internal->spec->layout = $tabsLayoutKindResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind> $rowsLayoutKind
     */
    public function rowsLayout(\Grafana\Foundation\Cog\Builder $rowsLayoutKind): static
    {
        $rowsLayoutKindResource = $rowsLayoutKind->build();
        $this->internal->spec->layout = $rowsLayoutKindResource;
    
        return $this;
    }

}
