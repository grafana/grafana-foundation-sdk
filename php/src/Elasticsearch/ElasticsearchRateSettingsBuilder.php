<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchRateSettings>
 */
class ElasticsearchRateSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchRateSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchRateSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchRateSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function unit(string $unit): static
    {
        $this->internal->unit = $unit;
    
        return $this;
    }
    public function mode(string $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }

}
