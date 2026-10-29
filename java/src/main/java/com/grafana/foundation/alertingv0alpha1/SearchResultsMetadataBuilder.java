// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;


/**
 * #SearchResultsMetadata carries the paging token and total authorised match
 * count.
 */
public class SearchResultsMetadataBuilder implements com.grafana.foundation.cog.Builder<SearchResultsMetadata> {
    protected final SearchResultsMetadata internal;
    
    public SearchResultsMetadataBuilder() {
        this.internal = new SearchResultsMetadata();
    }
    public SearchResultsMetadataBuilder continueArg(String continueArg) {
        this.internal.continueArg = continueArg;
        return this;
    }
    
    public SearchResultsMetadataBuilder totalHits(Long totalHits) {
        this.internal.totalHits = totalHits;
        return this;
    }
    
    public SearchResultsMetadataBuilder totalHitsRelation(TotalHitsRelation totalHitsRelation) {
        this.internal.totalHitsRelation = totalHitsRelation;
        return this;
    }
    public SearchResultsMetadata build() {
        return this.internal;
    }
}
