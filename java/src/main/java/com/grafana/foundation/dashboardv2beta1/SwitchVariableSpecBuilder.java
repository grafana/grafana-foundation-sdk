// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class SwitchVariableSpecBuilder implements com.grafana.foundation.cog.Builder<SwitchVariableSpec> {
    protected final SwitchVariableSpec internal;
    
    public SwitchVariableSpecBuilder() {
        this.internal = new SwitchVariableSpec();
    }
    public SwitchVariableSpecBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public SwitchVariableSpecBuilder current(String current) {
        this.internal.current = current;
        return this;
    }
    
    public SwitchVariableSpecBuilder enabledValue(String enabledValue) {
        this.internal.enabledValue = enabledValue;
        return this;
    }
    
    public SwitchVariableSpecBuilder disabledValue(String disabledValue) {
        this.internal.disabledValue = disabledValue;
        return this;
    }
    
    public SwitchVariableSpecBuilder label(String label) {
        this.internal.label = label;
        return this;
    }
    
    public SwitchVariableSpecBuilder hide(VariableHide hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public SwitchVariableSpecBuilder skipUrlSync(Boolean skipUrlSync) {
        this.internal.skipUrlSync = skipUrlSync;
        return this;
    }
    
    public SwitchVariableSpecBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    public SwitchVariableSpec build() {
        return this.internal;
    }
}
