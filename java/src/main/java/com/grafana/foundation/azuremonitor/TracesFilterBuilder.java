// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class TracesFilterBuilder implements com.grafana.foundation.cog.Builder<TracesFilter> {
    protected final TracesFilter internal;
    
    public TracesFilterBuilder() {
        this.internal = new TracesFilter();
    }
    public TracesFilterBuilder property(String property) {
        this.internal.property = property;
        return this;
    }
    
    public TracesFilterBuilder operation(String operation) {
        this.internal.operation = operation;
        return this;
    }
    
    public TracesFilterBuilder filters(List<String> filters) {
        this.internal.filters = filters;
        return this;
    }
    public TracesFilter build() {
        return this.internal;
    }
}
