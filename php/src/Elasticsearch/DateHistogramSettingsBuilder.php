<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\DateHistogramSettings>
 */
class DateHistogramSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\DateHistogramSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\DateHistogramSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\DateHistogramSettings
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
    public function trimEdges(string $trimEdges): static
    {
        $this->internal->trimEdges = $trimEdges;
    
        return $this;
    }
    public function offset(string $offset): static
    {
        $this->internal->offset = $offset;
    
        return $this;
    }
    public function timeZone(string $timeZone): static
    {
        $this->internal->timeZone = $timeZone;
    
        return $this;
    }

}
