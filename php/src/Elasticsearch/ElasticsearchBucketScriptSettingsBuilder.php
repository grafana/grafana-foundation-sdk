<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchBucketScriptSettings>
 */
class ElasticsearchBucketScriptSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchBucketScriptSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchBucketScriptSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchBucketScriptSettings
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

}
