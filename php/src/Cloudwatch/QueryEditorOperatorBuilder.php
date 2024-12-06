<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cloudwatch\QueryEditorOperator>
 */
class QueryEditorOperatorBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Cloudwatch\QueryEditorOperator $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Cloudwatch\QueryEditorOperator();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Cloudwatch\QueryEditorOperator
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
    /**
     * @param string|bool|int|array<string|bool|int> $value
     */
    public function value( $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }

}
