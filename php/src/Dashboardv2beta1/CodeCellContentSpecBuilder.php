<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\CodeCellContentSpec>
 */
class CodeCellContentSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\CodeCellContentSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\CodeCellContentSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\CodeCellContentSpec
     */
    public function build()
    {
        return $this->internal;
    }

    public function language(string $language): static
    {
        $this->internal->language = $language;
    
        return $this;
    }

    public function code(string $code): static
    {
        $this->internal->code = $code;
    
        return $this;
    }

    /**
     * @param array<int> $highlight
     */
    public function highlight(array $highlight): static
    {
        $this->internal->highlight = $highlight;
    
        return $this;
    }

    public function annotation(string $annotation): static
    {
        $this->internal->annotation = $annotation;
    
        return $this;
    }

}
