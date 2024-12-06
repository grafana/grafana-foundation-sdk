<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorProperty>
 */
class QueryEditorPropertyBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Cloudwatch\QueryEditorProperty $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Cloudwatch\QueryEditorProperty();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Cloudwatch\QueryEditorProperty
     */
    public function build()
    {
        return $this->internal;
    }

    public function type(\Grafana\Foundation\Cloudwatch\QueryEditorPropertyType $type): static
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
