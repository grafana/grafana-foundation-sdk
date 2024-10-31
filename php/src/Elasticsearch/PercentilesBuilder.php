<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\Percentiles>
 */
class PercentilesBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\Percentiles $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\Percentiles();
    $this->internal->type = "percentiles";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\Percentiles
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
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchPercentilesSettings> $settings
     */
    public function settings(\Grafana\Foundation\Cog\Builder $settings): static
    {
        $settingsResource = $settings->build();
        $this->internal->settings = $settingsResource;
    
        return $this;
    }
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }

}
