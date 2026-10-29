// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;


/**
 * #FacetValue is a single value/count pair in a facet breakdown.
 */
public class FacetValueBuilder implements com.grafana.foundation.cog.Builder<FacetValue> {
    protected final FacetValue internal;
    
    public FacetValueBuilder() {
        this.internal = new FacetValue();
    }
    public FacetValueBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    
    public FacetValueBuilder count(Long count) {
        this.internal.count = count;
        return this;
    }
    public FacetValue build() {
        return this.internal;
    }
}
