// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import java.util.List;
import java.util.LinkedList;
import java.util.Map;
import java.util.HashMap;

/**
 * #SearchResults is the response envelope, mirroring
 * search.grafana.app SearchResults.
 */
public class SearchResultsBuilder implements com.grafana.foundation.cog.Builder<SearchResults> {
    protected final SearchResults internal;
    
    public SearchResultsBuilder() {
        this.internal = new SearchResults();
    }
    public SearchResultsBuilder metadata(com.grafana.foundation.cog.Builder<SearchResultsMetadata> metadata) {
    SearchResultsMetadata metadataResource = metadata.build();
        this.internal.metadata = metadataResource;
        return this;
    }
    
    public SearchResultsBuilder items(List<com.grafana.foundation.cog.Builder<SearchResultHit>> items) {
        List<SearchResultHit> itemsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<SearchResultHit> r1 : items) {
                SearchResultHit itemsDepth1 = r1.build();
                itemsResources.add(itemsDepth1); 
        }
        this.internal.items = itemsResources;
        return this;
    }
    
    public SearchResultsBuilder facets(Map<String, List<com.grafana.foundation.cog.Builder<FacetValue>>> facets) {
        Map<String, List<FacetValue>> facetsResources = new HashMap<>();
        for (var entry1 : facets.entrySet()) {
                List<FacetValue> facetsDepth1 = new LinkedList<>();
                var value1 = entry1.getValue();
        for (com.grafana.foundation.cog.Builder<FacetValue> r2 : value1) {
                FacetValue facetsDepth2 = r2.build();
                facetsDepth1.add(facetsDepth2); 
        }
                facetsResources.put(entry1.getKey(), facetsDepth1);           
        }
        this.internal.facets = facetsResources;
        return this;
    }
    public SearchResults build() {
        return this.internal;
    }
}
