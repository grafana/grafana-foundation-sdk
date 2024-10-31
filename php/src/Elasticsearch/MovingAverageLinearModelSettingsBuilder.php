<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\MovingAverageLinearModelSettings>
 */
class MovingAverageLinearModelSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\MovingAverageLinearModelSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\MovingAverageLinearModelSettings();
    $this->internal->model = "linear";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\MovingAverageLinearModelSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function window(string $window): static
    {
        $this->internal->window = $window;
    
        return $this;
    }
    public function predict(string $predict): static
    {
        $this->internal->predict = $predict;
    
        return $this;
    }

}
