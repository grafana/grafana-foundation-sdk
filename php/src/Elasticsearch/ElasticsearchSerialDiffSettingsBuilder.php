<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchSerialDiffSettings>
 */
class ElasticsearchSerialDiffSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchSerialDiffSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchSerialDiffSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchSerialDiffSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function lag(string $lag): static
    {
        $this->internal->lag = $lag;
    
        return $this;
    }

}
