<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Map a field to a color.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\FieldColor>
 */
class FieldColorBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\FieldColor $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\FieldColor();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\FieldColor
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The main color scheme mode.
     */
    public function mode(\Grafana\Foundation\Dashboard\FieldColorModeId $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }
    /**
     * The fixed color value for fixed or shades color modes.
     */
    public function fixedColor(string $fixedColor): static
    {
        $this->internal->fixedColor = $fixedColor;
    
        return $this;
    }
    /**
     * Some visualizations need to know how to assign a series color from by value color schemes.
     */
    public function seriesBy(\Grafana\Foundation\Dashboard\FieldColorSeriesByMode $seriesBy): static
    {
        $this->internal->seriesBy = $seriesBy;
    
        return $this;
    }

}
