<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltModelSettingsSettings>
 */
class ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltModelSettingsSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltModelSettingsSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltModelSettingsSettings
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
    public function beta(string $beta): static
    {
        $this->internal->beta = $beta;
    
        return $this;
    }

}
