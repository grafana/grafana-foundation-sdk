<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript>
 */
class ElasticsearchInlineScriptBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript
     */
    public function build()
    {
        return $this->internal;
    }

    public function inline(string $inline): static
    {
        $this->internal->inline = $inline;
    
        return $this;
    }

}
