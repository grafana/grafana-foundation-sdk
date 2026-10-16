// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


/**
 * The notebook's own panel chain. It is a copy of the dashboard one down to the transformation,
 * which follows the dashboard v2 shape rather than the v2beta1 shape in this package. The chain has
 * to be forked rather than shared because PanelKind reaches TransformationKind through
 * QueryGroupKind, and those three are what Dashboard v2beta1 serves. Everything the chain does not
 * change (DataLink, VizConfigKind, PanelQueryKind, QueryOptionsSpec) stays shared.
 */
public class V2PanelKindBuilder implements com.grafana.foundation.cog.Builder<V2PanelKind> {
    protected final V2PanelKind internal;
    
    public V2PanelKindBuilder() {
        this.internal = new V2PanelKind();
        this.internal.kind = "Panel";
    }
    public V2PanelKindBuilder spec(com.grafana.foundation.cog.Builder<V2PanelSpec> spec) {
    V2PanelSpec specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    public V2PanelKind build() {
        return this.internal;
    }
}
