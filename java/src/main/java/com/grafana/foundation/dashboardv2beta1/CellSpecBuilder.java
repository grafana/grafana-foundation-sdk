// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class CellSpecBuilder implements com.grafana.foundation.cog.Builder<CellSpec> {
    protected final CellSpec internal;
    
    public CellSpecBuilder() {
        this.internal = new CellSpec();
    }
    public CellSpecBuilder content(CellContentKind content) {
        this.internal.content = content;
        return this;
    }
    public CellSpec build() {
        return this.internal;
    }
}
