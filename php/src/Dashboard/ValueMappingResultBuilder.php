<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Result used as replacement with text and color when the value matches
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\ValueMappingResult>
 */
class ValueMappingResultBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\ValueMappingResult $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\ValueMappingResult();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\ValueMappingResult
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Text to display when the value matches
     */
    public function text(string $text): static
    {
        $this->internal->text = $text;
    
        return $this;
    }
    /**
     * Text to use when the value matches
     */
    public function color(string $color): static
    {
        $this->internal->color = $color;
    
        return $this;
    }
    /**
     * Icon to display when the value matches. Only specific visualizations.
     */
    public function icon(string $icon): static
    {
        $this->internal->icon = $icon;
    
        return $this;
    }
    /**
     * Position in the mapping array. Only used internally.
     */
    public function index(int $index): static
    {
        $this->internal->index = $index;
    
        return $this;
    }

}
