<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchDerivativeSettings>
 */
class ElasticsearchDerivativeSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchDerivativeSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchDerivativeSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchDerivativeSettings
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

}
