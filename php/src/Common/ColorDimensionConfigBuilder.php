<?php

namespace Grafana\Foundation\Common;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ColorDimensionConfig>
 */
class ColorDimensionConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\ColorDimensionConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\ColorDimensionConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\ColorDimensionConfig
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * color value
     */
    public function fixed(string $fixed): static
    {
        $this->internal->fixed = $fixed;
    
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

}
