<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchUniqueCountSettings>
 */
class ElasticsearchUniqueCountSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchUniqueCountSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchUniqueCountSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchUniqueCountSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function precisionThreshold(string $precisionThreshold): static
    {
        $this->internal->precisionThreshold = $precisionThreshold;
    
        return $this;
    }
    public function missing(string $missing): static
    {
        $this->internal->missing = $missing;
    
        return $this;
    }

}
