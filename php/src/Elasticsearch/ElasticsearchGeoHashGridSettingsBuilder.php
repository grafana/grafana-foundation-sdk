<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchGeoHashGridSettings>
 */
class ElasticsearchGeoHashGridSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchGeoHashGridSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchGeoHashGridSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchGeoHashGridSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function precision(string $precision): static
    {
        $this->internal->precision = $precision;
    
        return $this;
    }

}
