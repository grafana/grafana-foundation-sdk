<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchLogsSettings>
 */
class ElasticsearchLogsSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchLogsSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchLogsSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchLogsSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function limit(string $limit): static
    {
        $this->internal->limit = $limit;
    
        return $this;
    }

}
