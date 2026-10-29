<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * #FacetValue is a single value/count pair in a facet breakdown.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\FacetValue>
 */
class FacetValueBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv0alpha1\FacetValue $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv0alpha1\FacetValue();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv0alpha1\FacetValue
     */
    public function build()
    {
        return $this->internal;
    }

    public function value(string $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }

    public function count(int $count): static
    {
        $this->internal->count = $count;
    
        return $this;
    }

}
