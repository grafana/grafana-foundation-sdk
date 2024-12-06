<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\BaseMovingAverageModelSettings>
 */
class BaseMovingAverageModelSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\BaseMovingAverageModelSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\BaseMovingAverageModelSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\BaseMovingAverageModelSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function model(\Grafana\Foundation\Elasticsearch\MovingAverageModel $model): static
    {
        $this->internal->model = $model;
    
        return $this;
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
