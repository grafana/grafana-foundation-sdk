<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time
 * settings, and a map of elements (panels and cells) referenced by the layout.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\NotebookSpec>
 */
class NotebookSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\NotebookSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\NotebookSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\NotebookSpec
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

    public function description(string $description): static
    {
        $this->internal->description = $description;
    
        return $this;
    }

    /**
     * @param array<string> $tags
     */
    public function tags(array $tags): static
    {
        $this->internal->tags = $tags;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec> $timeSettings
     */
    public function timeSettings(\Grafana\Foundation\Cog\Builder $timeSettings): static
    {
        $timeSettingsResource = $timeSettings->build();
        $this->internal->timeSettings = $timeSettingsResource;
    
        return $this;
    }

    /**
     * @param array<string, \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\CellKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\V2PanelKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind>> $elements
     */
    public function elements(array $elements): static
    {
            $elementsResources = [];
            foreach ($elements as $key1 => $val1) {
                    $elementsResources[$key1] = $val1->build();
            }
        $this->internal->elements = $elementsResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutKind> $layout
     */
    public function layout(\Grafana\Foundation\Cog\Builder $layout): static
    {
        $layoutResource = $layout->build();
        $this->internal->layout = $layoutResource;
    
        return $this;
    }

}
