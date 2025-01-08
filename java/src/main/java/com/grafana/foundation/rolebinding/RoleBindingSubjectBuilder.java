// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;


public class RoleBindingSubjectBuilder implements com.grafana.foundation.cog.Builder<RoleBindingSubject> {
    protected final RoleBindingSubject internal;
    
    public RoleBindingSubjectBuilder() {
        this.internal = new RoleBindingSubject();
    }
    public RoleBindingSubjectBuilder kind(RoleBindingSubjectKind kind) {
    this.internal.kind = kind;
        return this;
    }
    
    public RoleBindingSubjectBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    public RoleBindingSubject build() {
        return this.internal;
    }
}
