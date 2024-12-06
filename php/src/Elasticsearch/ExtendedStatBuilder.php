<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ExtendedStat>
 */
class ExtendedStatBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ExtendedStat $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ExtendedStat();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ExtendedStat
     */
    public function build()
    {
        return $this->internal;
    }

    public function label(string $label): static
    {
        $this->internal->label = $label;
    
        return $this;
    }
    public function value(\Grafana\Foundation\Elasticsearch\ExtendedStatMetaType $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }

}
