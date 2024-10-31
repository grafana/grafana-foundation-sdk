<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettings>
 */
class ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function missing(string $missing): static
    {
        $this->internal->missing = $missing;
    
        return $this;
    }

}
