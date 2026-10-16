<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentSpec>
 */
class MarkdownCellContentSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentSpec
     */
    public function build()
    {
        return $this->internal;
    }

    public function text(string $text): static
    {
        $this->internal->text = $text;
    
        return $this;
    }

}
