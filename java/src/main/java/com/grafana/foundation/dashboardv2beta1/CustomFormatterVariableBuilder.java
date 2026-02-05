// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class CustomFormatterVariableBuilder implements com.grafana.foundation.cog.Builder<CustomFormatterVariable> {
    protected final CustomFormatterVariable internal;
    
    public CustomFormatterVariableBuilder() {
        this.internal = new CustomFormatterVariable();
    }
    public CustomFormatterVariableBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public CustomFormatterVariableBuilder type(VariableType type) {
        this.internal.type = type;
        return this;
    }
    
    public CustomFormatterVariableBuilder multi(Boolean multi) {
        this.internal.multi = multi;
        return this;
    }
    
    public CustomFormatterVariableBuilder includeAll(Boolean includeAll) {
        this.internal.includeAll = includeAll;
        return this;
    }
    public CustomFormatterVariable build() {
        return this.internal;
    }
}
