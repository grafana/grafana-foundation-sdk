<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchMinSettings>
 */
class ElasticsearchMinSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchMinSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchMinSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchMinSettings
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

}
