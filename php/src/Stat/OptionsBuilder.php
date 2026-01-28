<?php

namespace Grafana\Foundation\Stat;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Stat\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Stat\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Stat\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Stat\Options
     */
    public function build()
    {
        return $this->internal;
    }

    public function graphMode(\Grafana\Foundation\Common\BigValueGraphMode $graphMode): static
    {
        $this->internal->graphMode = $graphMode;
    
        return $this;
    }

    public function colorMode(\Grafana\Foundation\Common\BigValueColorMode $colorMode): static
    {
        $this->internal->colorMode = $colorMode;
    
        return $this;
    }

    public function justifyMode(\Grafana\Foundation\Common\BigValueJustifyMode $justifyMode): static
    {
        $this->internal->justifyMode = $justifyMode;
    
        return $this;
    }

    public function textMode(\Grafana\Foundation\Common\BigValueTextMode $textMode): static
    {
        $this->internal->textMode = $textMode;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ReduceDataOptions> $reduceOptions
     */
    public function reduceOptions(\Grafana\Foundation\Cog\Builder $reduceOptions): static
    {
        $reduceOptionsResource = $reduceOptions->build();
        $this->internal->reduceOptions = $reduceOptionsResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizTextDisplayOptions> $text
     */
    public function text(\Grafana\Foundation\Cog\Builder $text): static
    {
        $textResource = $text->build();
        $this->internal->text = $textResource;
    
        return $this;
    }

    public function wideLayout(bool $wideLayout): static
    {
        $this->internal->wideLayout = $wideLayout;
    
        return $this;
    }

    public function orientation(\Grafana\Foundation\Common\VizOrientation $orientation): static
    {
        $this->internal->orientation = $orientation;
    
        return $this;
    }

}
