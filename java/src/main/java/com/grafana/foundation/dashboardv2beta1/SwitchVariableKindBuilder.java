// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class SwitchVariableKindBuilder implements com.grafana.foundation.cog.Builder<SwitchVariableKind> {
    protected final SwitchVariableKind internal;
    
    public SwitchVariableKindBuilder() {
        this.internal = new SwitchVariableKind();
        this.internal.kind = "SwitchVariable";
    }
    public SwitchVariableKindBuilder spec(com.grafana.foundation.cog.Builder<SwitchVariableSpec> spec) {
    SwitchVariableSpec specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    public SwitchVariableKind build() {
        return this.internal;
    }
}
