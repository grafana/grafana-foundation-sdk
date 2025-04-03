<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageEWMAModelSettingsSettings>
 */
class ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageEWMAModelSettingsSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageEWMAModelSettingsSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageEWMAModelSettingsSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function alpha(string $alpha): static
    {
        $this->internal->alpha = $alpha;
    
        return $this;
    }

}
