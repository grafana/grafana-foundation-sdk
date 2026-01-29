// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class KindBuilder implements com.grafana.foundation.cog.Builder<Kind> {
    protected final Kind internal;
    
    public KindBuilder() {
        this.internal = new Kind();
    }
    public KindBuilder kind(String kind) {
        this.internal.kind = kind;
        return this;
    }
    
    public KindBuilder spec(Object spec) {
        this.internal.spec = spec;
        return this;
    }
    
    public KindBuilder metadata(Object metadata) {
        this.internal.metadata = metadata;
        return this;
    }
    public Kind build() {
        return this.internal;
    }
}
