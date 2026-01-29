// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class VariableValueSingleBuilder implements com.grafana.foundation.cog.Builder<VariableValueSingle> {
    protected final VariableValueSingle internal;
    
    public VariableValueSingleBuilder() {
        this.internal = new VariableValueSingle();
    }
    public VariableValueSingleBuilder string(String string) {
        this.internal.string = string;
        return this;
    }
    
    public VariableValueSingleBuilder bool(Boolean bool) {
        this.internal.bool = bool;
        return this;
    }
    
    public VariableValueSingleBuilder float64(Double float64) {
        this.internal.float64 = float64;
        return this;
    }
    
    public VariableValueSingleBuilder customVariableValue(CustomVariableValue customVariableValue) {
        this.internal.customVariableValue = customVariableValue;
        return this;
    }
    public VariableValueSingle build() {
        return this.internal;
    }
}
