// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import java.util.List;

public class ScopesFiltersBuilder implements com.grafana.foundation.cog.Builder<ScopesFilters> {
    protected final ScopesFilters internal;
    
    public ScopesFiltersBuilder() {
        this.internal = new ScopesFilters();
    }
    public ScopesFiltersBuilder key(String key) {
        this.internal.key = key;
        return this;
    }
    
    public ScopesFiltersBuilder operator(String operator) {
        this.internal.operator = operator;
        return this;
    }
    
    public ScopesFiltersBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    
    public ScopesFiltersBuilder values(List<String> values) {
        this.internal.values = values;
        return this;
    }
    public ScopesFilters build() {
        return this.internal;
    }
}
