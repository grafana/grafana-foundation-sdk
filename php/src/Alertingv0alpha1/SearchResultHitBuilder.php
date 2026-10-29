<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * #SearchResultHit is a single match: its identity, an optional relevance
 * score (present only when the query included free text), and the requested
 * fields.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\SearchResultHit>
 */
class SearchResultHitBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv0alpha1\SearchResultHit $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv0alpha1\SearchResultHit();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv0alpha1\SearchResultHit
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\SearchResultResource> $resource
     */
    public function resource(\Grafana\Foundation\Cog\Builder $resource): static
    {
        $resourceResource = $resource->build();
        $this->internal->resource = $resourceResource;
    
        return $this;
    }

    public function score(float $score): static
    {
        $this->internal->score = $score;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFields> $fields
     */
    public function fields(\Grafana\Foundation\Cog\Builder $fields): static
    {
        $fieldsResource = $fields->build();
        $this->internal->fields = $fieldsResource;
    
        return $this;
    }

}
