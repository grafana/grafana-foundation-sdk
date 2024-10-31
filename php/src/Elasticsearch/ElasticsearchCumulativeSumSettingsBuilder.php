<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchCumulativeSumSettings>
 */
class ElasticsearchCumulativeSumSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchCumulativeSumSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchCumulativeSumSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchCumulativeSumSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function format(string $format): static
    {
        $this->internal->format = $format;
    
        return $this;
    }

}
