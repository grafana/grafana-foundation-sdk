// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


/**
 * A cell holds non-panel narrative content (markdown text, code) in a notebook layout.
 * Panel cells are not represented here — they reuse V2PanelKind.
 */
public class CellKindBuilder implements com.grafana.foundation.cog.Builder<CellKind> {
    protected final CellKind internal;
    
    public CellKindBuilder() {
        this.internal = new CellKind();
        this.internal.kind = "Cell";
    }
    public CellKindBuilder spec(com.grafana.foundation.cog.Builder<CellSpec> spec) {
    CellSpec specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    public CellKind build() {
        return this.internal;
    }
}
