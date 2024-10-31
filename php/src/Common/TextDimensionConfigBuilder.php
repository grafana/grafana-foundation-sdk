<?php

namespace Grafana\Foundation\Common;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TextDimensionConfig>
 */
class TextDimensionConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\TextDimensionConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\TextDimensionConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\TextDimensionConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(\Grafana\Foundation\Common\TextDimensionMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }
    /**
     * fixed: T -- will be added by each element
     */
    public function field(string $field): static
    {
        $this->internal->field = $field;
    
        return $this;
    }
    public function fixed(string $fixed): static
    {
        $this->internal->fixed = $fixed;
    
        return $this;
    }

}
