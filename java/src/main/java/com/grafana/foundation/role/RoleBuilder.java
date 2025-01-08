// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.role;


public class RoleBuilder implements com.grafana.foundation.cog.Builder<Role> {
    protected final Role internal;
    
    public RoleBuilder() {
        this.internal = new Role();
    }
    public RoleBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public RoleBuilder displayName(String displayName) {
    this.internal.displayName = displayName;
        return this;
    }
    
    public RoleBuilder groupName(String groupName) {
    this.internal.groupName = groupName;
        return this;
    }
    
    public RoleBuilder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public RoleBuilder hidden(Boolean hidden) {
    this.internal.hidden = hidden;
        return this;
    }
    public Role build() {
        return this.internal;
    }
}
