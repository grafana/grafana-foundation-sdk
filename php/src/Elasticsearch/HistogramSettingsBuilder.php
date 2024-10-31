<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\HistogramSettings>
 */
class HistogramSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\HistogramSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\HistogramSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\HistogramSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function interval(string $interval): static
    {
        $this->internal->interval = $interval;
    
        return $this;
    }
    public function minDocCount(string $minDocCount): static
    {
        $this->internal->minDocCount = $minDocCount;
    
        return $this;
    }

}
