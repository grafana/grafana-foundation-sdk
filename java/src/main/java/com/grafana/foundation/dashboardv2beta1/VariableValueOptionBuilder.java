// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class VariableValueOptionBuilder implements com.grafana.foundation.cog.Builder<VariableValueOption> {
    protected final VariableValueOption internal;
    
    public VariableValueOptionBuilder() {
        this.internal = new VariableValueOption();
    }
    public VariableValueOptionBuilder label(String label) {
        this.internal.label = label;
        return this;
    }
    
    public VariableValueOptionBuilder value(com.grafana.foundation.cog.Builder<VariableValueSingle> value) {
    VariableValueSingle valueResource = value.build();
        this.internal.value = valueResource;
        return this;
    }
    
    public VariableValueOptionBuilder group(String group) {
        this.internal.group = group;
        return this;
    }
    public VariableValueOption build() {
        return this.internal;
    }
}
