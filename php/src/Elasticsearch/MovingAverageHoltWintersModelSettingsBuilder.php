<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\MovingAverageHoltWintersModelSettings>
 */
class MovingAverageHoltWintersModelSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\MovingAverageHoltWintersModelSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\MovingAverageHoltWintersModelSettings();
    $this->internal->model = "holt_winters";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\MovingAverageHoltWintersModelSettings
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltWintersModelSettingsSettings> $settings
     */
    public function settings(\Grafana\Foundation\Cog\Builder $settings): static
    {
        $settingsResource = $settings->build();
        $this->internal->settings = $settingsResource;
    
        return $this;
    }
    public function window(string $window): static
    {
        $this->internal->window = $window;
    
        return $this;
    }
    public function minimize(bool $minimize): static
    {
        $this->internal->minimize = $minimize;
    
        return $this;
    }
    public function predict(string $predict): static
    {
        $this->internal->predict = $predict;
    
        return $this;
    }

}
