// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class SelectableValueBuilder implements com.grafana.foundation.cog.Builder<SelectableValue> {
    protected final SelectableValue internal;
    
    public SelectableValueBuilder() {
        this.internal = new SelectableValue();
    }
    public SelectableValueBuilder label(String label) {
        this.internal.label = label;
        return this;
    }
    
    public SelectableValueBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    public SelectableValue build() {
        return this.internal;
    }
}
