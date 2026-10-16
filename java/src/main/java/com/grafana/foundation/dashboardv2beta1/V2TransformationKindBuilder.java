// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


/**
 * Dashboard v2 shape: the transformation ID moved from `kind` to `group`.
 */
public class V2TransformationKindBuilder implements com.grafana.foundation.cog.Builder<V2TransformationKind> {
    protected final V2TransformationKind internal;
    
    public V2TransformationKindBuilder() {
        this.internal = new V2TransformationKind();
        this.internal.kind = "Transformation";
    }
    public V2TransformationKindBuilder group(String group) {
        this.internal.group = group;
        return this;
    }
    
    public V2TransformationKindBuilder spec(com.grafana.foundation.cog.Builder<V2TransformationSpec> spec) {
    V2TransformationSpec specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    public V2TransformationKind build() {
        return this.internal;
    }
}
