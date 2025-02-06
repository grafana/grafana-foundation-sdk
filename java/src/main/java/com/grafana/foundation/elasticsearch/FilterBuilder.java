// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class FilterBuilder implements com.grafana.foundation.cog.Builder<Filter> {
    protected final Filter internal;
    
    public FilterBuilder() {
        this.internal = new Filter();
    }
    public FilterBuilder query(String query) {
        this.internal.query = query;
        return this;
    }
    
    public FilterBuilder label(String label) {
        this.internal.label = label;
        return this;
    }
    public Filter build() {
        return this.internal;
    }
}
