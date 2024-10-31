<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\TermsSettings>
 */
class TermsSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\TermsSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\TermsSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\TermsSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function order(\Grafana\Foundation\Elasticsearch\TermsOrder $order): static
    {
        $this->internal->order = $order;
    
        return $this;
    }
    public function size(string $size): static
    {
        $this->internal->size = $size;
    
        return $this;
    }
    public function minDocCount(string $minDocCount): static
    {
        $this->internal->minDocCount = $minDocCount;
    
        return $this;
    }
    public function orderBy(string $orderBy): static
    {
        $this->internal->orderBy = $orderBy;
    
        return $this;
    }
    public function missing(string $missing): static
    {
        $this->internal->missing = $missing;
    
        return $this;
    }

}
