<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\MovingAverageSimpleModelSettings>
 */
class MovingAverageSimpleModelSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\MovingAverageSimpleModelSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\MovingAverageSimpleModelSettings();
    $this->internal->model = "simple";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\MovingAverageSimpleModelSettings
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
