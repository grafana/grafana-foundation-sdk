<?php

namespace Grafana\Foundation\Canvas;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\CanvasOptionsRoot>
 */
class CanvasOptionsRootBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Canvas\CanvasOptionsRoot $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Canvas\CanvasOptionsRoot();
    $this->internal->type = "frame";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Canvas\CanvasOptionsRoot
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Name of the root element
     */
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    /**
     * The list of canvas elements attached to the root element
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\CanvasElementOptions>> $elements
     */
    public function elements(array $elements): static
    {
            $elementsResources = [];
            foreach ($elements as $r1) {
                    $elementsResources[] = $r1->build();
            }
        $this->internal->elements = $elementsResources;
    
        return $this;
    }

}
