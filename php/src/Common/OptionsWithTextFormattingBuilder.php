<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\OptionsWithTextFormatting>
 */
class OptionsWithTextFormattingBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\OptionsWithTextFormatting $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\OptionsWithTextFormatting();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\OptionsWithTextFormatting
     */
    public function build()
    {
        return $this->internal;
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

}
