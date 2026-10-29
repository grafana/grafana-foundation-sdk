// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;


/**
 * #SearchResultResource is the full identity of a hit.
 */
public class SearchResultResourceBuilder implements com.grafana.foundation.cog.Builder<SearchResultResource> {
    protected final SearchResultResource internal;
    
    public SearchResultResourceBuilder() {
        this.internal = new SearchResultResource();
    }
    public SearchResultResourceBuilder group(String group) {
        this.internal.group = group;
        return this;
    }
    
    public SearchResultResourceBuilder resource(String resource) {
        this.internal.resource = resource;
        return this;
    }
    
    public SearchResultResourceBuilder kind(String kind) {
        this.internal.kind = kind;
        return this;
    }
    
    public SearchResultResourceBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    public SearchResultResource build() {
        return this.internal;
    }
}
