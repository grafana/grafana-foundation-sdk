<?php

namespace Grafana\Foundation\Bigquery;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Bigquery\QueryEditorProperty>
 */
class QueryEditorPropertyBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Bigquery\QueryEditorProperty $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Bigquery\QueryEditorProperty();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Bigquery\QueryEditorProperty
     */
    public function build()
    {
        return $this->internal;
    }

    public function type(\Grafana\Foundation\Bigquery\QueryEditorPropertyType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }

    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }

}
