<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchRawDocumentSettings>
 */
class ElasticsearchRawDocumentSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchRawDocumentSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchRawDocumentSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchRawDocumentSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function size(string $size): static
    {
        $this->internal->size = $size;
    
        return $this;
    }

}
