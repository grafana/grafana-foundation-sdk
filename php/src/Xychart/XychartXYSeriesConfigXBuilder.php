<?php

namespace Grafana\Foundation\Xychart;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XychartXYSeriesConfigX>
 */
class XychartXYSeriesConfigXBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Xychart\XychartXYSeriesConfigX $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Xychart\XychartXYSeriesConfigX();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Xychart\XychartXYSeriesConfigX
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
