<?php

namespace Grafana\Foundation\Canvas;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Canvas\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Canvas\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Canvas\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Enable inline editing
     */
    public function inlineEditing(bool $inlineEditing): static
    {
        $this->internal->inlineEditing = $inlineEditing;
    
        return $this;
    }

    /**
     * Show all available element types
     */
    public function showAdvancedTypes(bool $showAdvancedTypes): static
    {
        $this->internal->showAdvancedTypes = $showAdvancedTypes;
    
        return $this;
    }

    /**
     * Enable pan and zoom
     */
    public function panZoom(bool $panZoom): static
    {
        $this->internal->panZoom = $panZoom;
    
        return $this;
    }

    /**
     * Enable infinite pan
     */
    public function infinitePan(bool $infinitePan): static
    {
        $this->internal->infinitePan = $infinitePan;
    
        return $this;
    }

    /**
     * The root element of canvas (frame), where all canvas elements are nested
     * TODO: Figure out how to define a default value for this
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\CanvasOptionsRoot> $root
     */
    public function root(\Grafana\Foundation\Cog\Builder $root): static
    {
        $rootResource = $root->build();
        $this->internal->root = $rootResource;
    
        return $this;
    }

}
