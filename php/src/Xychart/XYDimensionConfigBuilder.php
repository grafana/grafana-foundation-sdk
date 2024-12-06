<?php

namespace Grafana\Foundation\Xychart;

/**
 * Configuration for the Table/Auto mode
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XYDimensionConfig>
 */
class XYDimensionConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Xychart\XYDimensionConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Xychart\XYDimensionConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Xychart\XYDimensionConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function frame(int $frame): static
    {
        if (!($frame >= 0)) {
            throw new \ValueError('$frame must be >= 0');
        }
        $this->internal->frame = $frame;
    
        return $this;
    }
    public function x(string $x): static
    {
        $this->internal->x = $x;
    
        return $this;
    }
    /**
     * @param array<string> $exclude
     */
    public function exclude(array $exclude): static
    {
        $this->internal->exclude = $exclude;
    
        return $this;
    }

}
