// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class AzureTracesFilterBuilder implements com.grafana.foundation.cog.Builder<AzureTracesFilter> {
    protected final AzureTracesFilter internal;
    
    public AzureTracesFilterBuilder() {
        this.internal = new AzureTracesFilter();
    }
    public AzureTracesFilterBuilder property(String property) {
    this.internal.property = property;
        return this;
    }
    
    public AzureTracesFilterBuilder operation(String operation) {
    this.internal.operation = operation;
        return this;
    }
    
    public AzureTracesFilterBuilder filters(List<String> filters) {
    this.internal.filters = filters;
        return this;
    }
    public AzureTracesFilter build() {
        return this.internal;
    }
}
