<?php

namespace Grafana\Foundation\Xychart;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XychartFieldConfigPointSize>
 */
class XychartFieldConfigPointSizeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Xychart\XychartFieldConfigPointSize $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Xychart\XychartFieldConfigPointSize();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Xychart\XychartFieldConfigPointSize
     */
    public function build()
    {
        return $this->internal;
    }

    public function fixed(int $fixed): static
    {
        if (!($fixed >= 0)) {
            throw new \ValueError('$fixed must be >= 0');
        }
        $this->internal->fixed = $fixed;
    
        return $this;
    }

    public function min(int $min): static
    {
        if (!($min >= 0)) {
            throw new \ValueError('$min must be >= 0');
        }
        $this->internal->min = $min;
    
        return $this;
    }

    public function max(int $max): static
    {
        if (!($max >= 0)) {
            throw new \ValueError('$max must be >= 0');
        }
        $this->internal->max = $max;
    
        return $this;
    }

}
