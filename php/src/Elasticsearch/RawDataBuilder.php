<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\RawData>
 */
class RawDataBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\RawData $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\RawData();
    $this->internal->type = "raw_data";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\RawData
     */
    public function build()
    {
        return $this->internal;
    }

    public function id(string $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchRawDataSettings> $settings
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
