<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchExtendedStatsSettings>
 */
class ElasticsearchExtendedStatsSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchExtendedStatsSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchExtendedStatsSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchExtendedStatsSettings
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param string|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript> $script
     */
    public function script( $script): static
    {
        /** @var string|\Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript $scriptResource */
        $scriptResource = $script instanceof \Grafana\Foundation\Cog\Builder ? $script->build() : $script;
        $this->internal->script = $scriptResource;
    
        return $this;
    }
    public function missing(string $missing): static
    {
        $this->internal->missing = $missing;
    
        return $this;
    }
    public function sigma(string $sigma): static
    {
        $this->internal->sigma = $sigma;
    
        return $this;
    }

}
