<?php

namespace Grafana\Foundation\Testdata;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\USAQuery>
 */
class USAQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Testdata\USAQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Testdata\USAQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Testdata\USAQuery
     */
    public function build()
    {
        return $this->internal;
    }

    public function mode(string $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }
    public function period(string $period): static
    {
        $this->internal->period = $period;
    
        return $this;
    }
    /**
     * @param array<string> $fields
     */
    public function fields(array $fields): static
    {
        $this->internal->fields = $fields;
    
        return $this;
    }
    /**
     * @param array<string> $states
     */
    public function states(array $states): static
    {
        $this->internal->states = $states;
    
        return $this;
    }

}
