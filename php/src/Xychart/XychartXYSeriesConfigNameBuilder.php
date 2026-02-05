<?php

namespace Grafana\Foundation\Xychart;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Xychart\XychartXYSeriesConfigName>
 */
class XychartXYSeriesConfigNameBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Xychart\XychartXYSeriesConfigName $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Xychart\XychartXYSeriesConfigName();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Xychart\XychartXYSeriesConfigName
     */
    public function build()
    {
        return $this->internal;
    }

    public function fixed(string $fixed): static
    {
        $this->internal->fixed = $fixed;
    
        return $this;
    }

}
