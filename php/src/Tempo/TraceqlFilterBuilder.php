<?php

namespace Grafana\Foundation\Tempo;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Tempo\TraceqlFilter>
 */
class TraceqlFilterBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Tempo\TraceqlFilter $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Tempo\TraceqlFilter();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Tempo\TraceqlFilter
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Uniquely identify the filter, will not be used in the query generation
     */
    public function id(string $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    /**
     * The tag for the search filter, for example: .http.status_code, .service.name, status
     */
    public function tag(string $tag): static
    {
        $this->internal->tag = $tag;
    
        return $this;
    }
    /**
     * The operator that connects the tag to the value, for example: =, >, !=, =~
     */
    public function operator(string $operator): static
    {
        $this->internal->operator = $operator;
    
        return $this;
    }
    /**
     * The value for the search filter
     * @param string|array<string> $value
     */
    public function value( $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }
    /**
     * The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query
     */
    public function valueType(string $valueType): static
    {
        $this->internal->valueType = $valueType;
    
        return $this;
    }
    /**
     * The scope of the filter, can either be unscoped/all scopes, resource or span
     */
    public function scope(\Grafana\Foundation\Tempo\TraceqlSearchScope $scope): static
    {
        $this->internal->scope = $scope;
    
        return $this;
    }

}
