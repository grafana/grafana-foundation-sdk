// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class V2QueryGroupKindBuilder implements com.grafana.foundation.cog.Builder<V2QueryGroupKind> {
    protected final V2QueryGroupKind internal;
    
    public V2QueryGroupKindBuilder() {
        this.internal = new V2QueryGroupKind();
        this.internal.kind = "QueryGroup";
    }
    public V2QueryGroupKindBuilder spec(com.grafana.foundation.cog.Builder<V2QueryGroupSpec> spec) {
    V2QueryGroupSpec specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    public V2QueryGroupKind build() {
        return this.internal;
    }
}
