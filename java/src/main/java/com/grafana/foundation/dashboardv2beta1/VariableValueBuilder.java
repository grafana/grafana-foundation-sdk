// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class VariableValueBuilder implements com.grafana.foundation.cog.Builder<VariableValue> {
    protected final VariableValue internal;
    
    public VariableValueBuilder() {
        this.internal = new VariableValue();
    }
    public VariableValueBuilder string(String string) {
        this.internal.string = string;
        return this;
    }
    
    public VariableValueBuilder bool(Boolean bool) {
        this.internal.bool = bool;
        return this;
    }
    
    public VariableValueBuilder float64(Double float64) {
        this.internal.float64 = float64;
        return this;
    }
    
    public VariableValueBuilder customVariableValue(CustomVariableValue customVariableValue) {
        this.internal.customVariableValue = customVariableValue;
        return this;
    }
    
    public VariableValueBuilder arrayOfVariableValueSingle(List<com.grafana.foundation.cog.Builder<VariableValueSingle>> arrayOfVariableValueSingle) {
        List<VariableValueSingle> arrayOfVariableValueSingleResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<VariableValueSingle> r1 : arrayOfVariableValueSingle) {
                VariableValueSingle arrayOfVariableValueSingleDepth1 = r1.build();
                arrayOfVariableValueSingleResources.add(arrayOfVariableValueSingleDepth1); 
        }
        this.internal.arrayOfVariableValueSingle = arrayOfVariableValueSingleResources;
        return this;
    }
    public VariableValue build() {
        return this.internal;
    }
}
