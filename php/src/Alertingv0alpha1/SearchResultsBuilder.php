<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * #SearchResults is the response envelope, mirroring
 * search.grafana.app SearchResults.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\SearchResults>
 */
class SearchResultsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv0alpha1\SearchResults $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv0alpha1\SearchResults();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv0alpha1\SearchResults
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadata> $metadata
     */
    public function metadata(\Grafana\Foundation\Cog\Builder $metadata): static
    {
        $metadataResource = $metadata->build();
        $this->internal->metadata = $metadataResource;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\SearchResultHit>> $items
     */
    public function items(array $items): static
    {
            $itemsResources = [];
            foreach ($items as $r1) {
                    $itemsResources[] = $r1->build();
            }
        $this->internal->items = $itemsResources;
    
        return $this;
    }

    /**
     * @param array<string, array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\FacetValue>>> $facets
     */
    public function facets(array $facets): static
    {
            $facetsResources = [];
            foreach ($facets as $key1 => $val1) {
                    $facetsDepth1 = [];
            foreach ($val1 as $r2) {
                    $facetsDepth1[] = $r2->build();
            }
    
                    $facetsResources[$key1] = $facetsDepth1;
            }
        $this->internal->facets = $facetsResources;
    
        return $this;
    }

}
