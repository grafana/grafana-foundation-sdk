<?php

namespace Grafana\Foundation\Text;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Text\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Text\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Text\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Text\Options
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(\Grafana\Foundation\Text\TextMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Text\CodeOptions> $code
     */
    public function code(\Grafana\Foundation\Cog\Builder $code): static
    {
        $codeResource = $code->build();
        $this->internal->code = $codeResource;
    
        return $this;
    }

    public function content(string $content): static
    {
        $this->internal->content = $content;
    
        return $this;
    }

}
