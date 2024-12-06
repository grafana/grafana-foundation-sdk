<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\PipelineVariable>
 */
class PipelineVariableBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\PipelineVariable $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\PipelineVariable();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\PipelineVariable
     */
    public function build()
    {
        return $this->internal;
    }

    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    public function pipelineAgg(string $pipelineAgg): static
    {
        $this->internal->pipelineAgg = $pipelineAgg;
    
        return $this;
    }

}
