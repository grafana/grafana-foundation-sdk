// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;


public class CustomRoleRefBuilder implements com.grafana.foundation.cog.Builder<CustomRoleRef> {
    protected final CustomRoleRef internal;
    
    public CustomRoleRefBuilder() {
        this.internal = new CustomRoleRef();
        this.internal.kind = "Role";
    }
    public CustomRoleRefBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    public CustomRoleRef build() {
        return this.internal;
    }
}
