<?php

namespace Grafana\Foundation\Common;

/**
 * Links to a resource (image/svg path)
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ResourceDimensionConfig>
 */
class ResourceDimensionConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\ResourceDimensionConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\ResourceDimensionConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\ResourceDimensionConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(\Grafana\Foundation\Common\ResourceDimensionMode $mode): static
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
