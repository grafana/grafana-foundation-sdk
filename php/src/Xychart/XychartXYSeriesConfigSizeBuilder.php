<?php

namespace Grafana\Foundation\Xychart;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XychartXYSeriesConfigSize>
 */
class XychartXYSeriesConfigSizeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Xychart\XychartXYSeriesConfigSize $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Xychart\XychartXYSeriesConfigSize();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Xychart\XychartXYSeriesConfigSize
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\MatcherConfig> $matcher
     */
    public function matcher(\Grafana\Foundation\Cog\Builder $matcher): static
    {
        $matcherResource = $matcher->build();
        $this->internal->matcher = $matcherResource;
    
        return $this;
    }

}
