// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.accesspolicy;


public class RoleRefBuilder implements com.grafana.foundation.cog.Builder<RoleRef> {
    protected final RoleRef internal;
    
    public RoleRefBuilder() {
        this.internal = new RoleRef();
    }
    public RoleRefBuilder kind(RoleRefKind kind) {
    this.internal.kind = kind;
        return this;
    }
    
    public RoleRefBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public RoleRefBuilder xname(String xname) {
    this.internal.xname = xname;
        return this;
    }
    public RoleRef build() {
        return this.internal;
    }
}
