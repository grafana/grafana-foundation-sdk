// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.accesspolicy;


public class ResourceRefBuilder implements com.grafana.foundation.cog.Builder<ResourceRef> {
    protected final ResourceRef internal;
    
    public ResourceRefBuilder() {
        this.internal = new ResourceRef();
    }
    public ResourceRefBuilder kind(String kind) {
    this.internal.kind = kind;
        return this;
    }
    
    public ResourceRefBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    public ResourceRef build() {
        return this.internal;
    }
}
