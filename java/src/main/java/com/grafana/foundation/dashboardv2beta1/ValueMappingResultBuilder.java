// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class ValueMappingResultBuilder implements com.grafana.foundation.cog.Builder<ValueMappingResult> {
    protected final ValueMappingResult internal;
    
    public ValueMappingResultBuilder() {
        this.internal = new ValueMappingResult();
    }
    public ValueMappingResultBuilder text(String text) {
        this.internal.text = text;
        return this;
    }
    
    public ValueMappingResultBuilder color(String color) {
        this.internal.color = color;
        return this;
    }
    
    public ValueMappingResultBuilder icon(String icon) {
        this.internal.icon = icon;
        return this;
    }
    
    public ValueMappingResultBuilder index(Integer index) {
        this.internal.index = index;
        return this;
    }
    public ValueMappingResult build() {
        return this.internal;
    }
}
