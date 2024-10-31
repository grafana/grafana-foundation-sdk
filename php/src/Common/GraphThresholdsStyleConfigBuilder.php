<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\GraphThresholdsStyleConfig>
 */
class GraphThresholdsStyleConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\GraphThresholdsStyleConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\GraphThresholdsStyleConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\GraphThresholdsStyleConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(\Grafana\Foundation\Common\GraphTresholdsStyleMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }

}
