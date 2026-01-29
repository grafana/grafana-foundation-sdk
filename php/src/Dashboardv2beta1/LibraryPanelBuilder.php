<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind>
 */
class LibraryPanelBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind();
    $this->internal->kind = "LibraryPanel";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKindSpec> $spec
     */
    public function spec(\Grafana\Foundation\Cog\Builder $spec): static
    {
        $specResource = $spec->build();
        $this->internal->spec = $specResource;
    
        return $this;
    }

    /**
     * Panel ID for the library panel in the dashboard
     */
    public function id(float $id): static
    {
        $this->internal->spec->id = $id;
    
        return $this;
    }

    /**
     * Title for the library panel in the dashboard
     */
    public function title(string $title): static
    {
        $this->internal->spec->title = $title;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\LibraryPanelRef> $libraryPanel
     */
    public function libraryPanel(\Grafana\Foundation\Cog\Builder $libraryPanel): static
    {
        $libraryPanelResource = $libraryPanel->build();
        $this->internal->spec->libraryPanel = $libraryPanelResource;
    
        return $this;
    }

}
