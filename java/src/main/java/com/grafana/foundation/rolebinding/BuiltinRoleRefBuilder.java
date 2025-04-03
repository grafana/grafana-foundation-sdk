// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;


public class BuiltinRoleRefBuilder implements com.grafana.foundation.cog.Builder<BuiltinRoleRef> {
    protected final BuiltinRoleRef internal;
    
    public BuiltinRoleRefBuilder() {
        this.internal = new BuiltinRoleRef();
        this.internal.kind = "BuiltinRole";
    }
    public BuiltinRoleRefBuilder name(BuiltinRoleRefName name) {
        this.internal.name = name;
        return this;
    }
    public BuiltinRoleRef build() {
        return this.internal;
    }
}
