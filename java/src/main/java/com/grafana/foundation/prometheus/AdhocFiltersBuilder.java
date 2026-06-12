// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import java.util.List;

public class AdhocFiltersBuilder implements com.grafana.foundation.cog.Builder<AdhocFilters> {
    protected final AdhocFilters internal;
    
    public AdhocFiltersBuilder() {
        this.internal = new AdhocFilters();
    }
    public AdhocFiltersBuilder key(String key) {
        this.internal.key = key;
        return this;
    }
    
    public AdhocFiltersBuilder operator(String operator) {
        this.internal.operator = operator;
        return this;
    }
    
    public AdhocFiltersBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    
    public AdhocFiltersBuilder values(List<String> values) {
        this.internal.values = values;
        return this;
    }
    public AdhocFilters build() {
        return this.internal;
    }
}
