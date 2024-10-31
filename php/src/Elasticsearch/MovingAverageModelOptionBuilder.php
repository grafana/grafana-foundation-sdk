<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\MovingAverageModelOption>
 */
class MovingAverageModelOptionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\MovingAverageModelOption $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\MovingAverageModelOption();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\MovingAverageModelOption
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
    public function value(\Grafana\Foundation\Elasticsearch\MovingAverageModel $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }

}
