// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class StackableFieldConfigBuilder implements com.grafana.foundation.cog.Builder<StackableFieldConfig> {
    protected final StackableFieldConfig internal;
    
    public StackableFieldConfigBuilder() {
        this.internal = new StackableFieldConfig();
    }
    public StackableFieldConfigBuilder stacking(com.grafana.foundation.cog.Builder<StackingConfig> stacking) {
        this.internal.stacking = stacking.build();
        return this;
    }
    public StackableFieldConfig build() {
        return this.internal;
    }
}
