<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchTopMetricsSettings>
 */
class ElasticsearchTopMetricsSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchTopMetricsSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchTopMetricsSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchTopMetricsSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function order(string $order): static
    {
        $this->internal->order = $order;
    
        return $this;
    }
    public function orderBy(string $orderBy): static
    {
        $this->internal->orderBy = $orderBy;
    
        return $this;
    }
    /**
     * @param array<string> $metrics
     */
    public function metrics(array $metrics): static
    {
        $this->internal->metrics = $metrics;
    
        return $this;
    }

}
