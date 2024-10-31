<?php

namespace Grafana\Foundation\Common;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDimensionConfig>
 */
class ScaleDimensionConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\ScaleDimensionConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\ScaleDimensionConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\ScaleDimensionConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function min(float $min): static
    {
        $this->internal->min = $min;
    
        return $this;
    }
    public function max(float $max): static
    {
        $this->internal->max = $max;
    
        return $this;
    }
    public function fixed(float $fixed): static
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
    /**
     * | *"linear"
     */
    public function mode(\Grafana\Foundation\Common\ScaleDimensionMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }

}
