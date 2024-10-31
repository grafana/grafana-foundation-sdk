<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\GeoHashGridSettings>
 */
class GeoHashGridSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\GeoHashGridSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\GeoHashGridSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\GeoHashGridSettings
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
