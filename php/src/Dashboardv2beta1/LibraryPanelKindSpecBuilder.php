<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKindSpec>
 */
class LibraryPanelKindSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\LibraryPanelKindSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\LibraryPanelKindSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\LibraryPanelKindSpec
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Panel ID for the library panel in the dashboard
     */
    public function id(float $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }

    /**
     * Title for the library panel in the dashboard
     */
    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\LibraryPanelRef> $libraryPanel
     */
    public function libraryPanel(\Grafana\Foundation\Cog\Builder $libraryPanel): static
    {
        $libraryPanelResource = $libraryPanel->build();
        $this->internal->libraryPanel = $libraryPanelResource;
    
        return $this;
    }

}
