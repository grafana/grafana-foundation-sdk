<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * #SearchResultsMetadata carries the paging token and total authorised match
 * count.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadata>
 */
class SearchResultsMetadataBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadata $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadata();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadata
     */
    public function build()
    {
        return $this->internal;
    }

    public function continue(string $continue): static
    {
        $this->internal->continue = $continue;
    
        return $this;
    }

    public function totalHits(int $totalHits): static
    {
        $this->internal->totalHits = $totalHits;
    
        return $this;
    }

    /**
     * Always read totalHits together with totalHitsRelation.
     */
    public function totalHitsRelation(\Grafana\Foundation\Alertingv0alpha1\TotalHitsRelation $totalHitsRelation): static
    {
        $this->internal->totalHitsRelation = $totalHitsRelation;
    
        return $this;
    }

}
