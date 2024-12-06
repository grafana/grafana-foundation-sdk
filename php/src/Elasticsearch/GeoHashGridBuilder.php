<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\GeoHashGrid>
 */
class GeoHashGridBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\GeoHashGrid $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\GeoHashGrid();
    $this->internal->type = "geohash_grid";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\GeoHashGrid
     */
    public function build()
    {
        return $this->internal;
    }

    public function field(string $field): static
    {
        $this->internal->field = $field;
    
        return $this;
    }
    public function id(string $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchGeoHashGridSettings> $settings
     */
    public function settings(\Grafana\Foundation\Cog\Builder $settings): static
    {
        $settingsResource = $settings->build();
        $this->internal->settings = $settingsResource;
    
        return $this;
    }

}
