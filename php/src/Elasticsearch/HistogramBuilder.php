<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Histogram>
 */
class HistogramBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\Histogram $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\Histogram();
    $this->internal->type = "histogram";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\Histogram
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
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchHistogramSettings> $settings
     */
    public function settings(\Grafana\Foundation\Cog\Builder $settings): static
    {
        $settingsResource = $settings->build();
        $this->internal->settings = $settingsResource;
    
        return $this;
    }

}
