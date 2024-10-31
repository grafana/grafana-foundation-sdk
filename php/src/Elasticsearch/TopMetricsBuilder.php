<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\TopMetrics>
 */
class TopMetricsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\TopMetrics $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\TopMetrics();
    $this->internal->type = "top_metrics";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\TopMetrics
     */
    public function build()
    {
        return $this->internal;
    }

    public function id(string $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchTopMetricsSettings> $settings
     */
    public function settings(\Grafana\Foundation\Cog\Builder $settings): static
    {
        $settingsResource = $settings->build();
        $this->internal->settings = $settingsResource;
    
        return $this;
    }
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }

}
