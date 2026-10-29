// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;


/**
 * #SearchResultHit is a single match: its identity, an optional relevance
 * score (present only when the query included free text), and the requested
 * fields.
 */
public class SearchResultHitBuilder implements com.grafana.foundation.cog.Builder<SearchResultHit> {
    protected final SearchResultHit internal;
    
    public SearchResultHitBuilder() {
        this.internal = new SearchResultHit();
    }
    public SearchResultHitBuilder resource(com.grafana.foundation.cog.Builder<SearchResultResource> resource) {
    SearchResultResource resourceResource = resource.build();
        this.internal.resource = resourceResource;
        return this;
    }
    
    public SearchResultHitBuilder score(Double score) {
        this.internal.score = score;
        return this;
    }
    
    public SearchResultHitBuilder fields(com.grafana.foundation.cog.Builder<RuleSearchHitFields> fields) {
    RuleSearchHitFields fieldsResource = fields.build();
        this.internal.fields = fieldsResource;
        return this;
    }
    public SearchResultHit build() {
        return this.internal;
    }
}
