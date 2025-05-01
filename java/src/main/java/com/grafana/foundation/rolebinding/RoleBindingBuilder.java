// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;


public class RoleBindingBuilder implements com.grafana.foundation.cog.Builder<RoleBinding> {
    protected final RoleBinding internal;
    
    public RoleBindingBuilder() {
        this.internal = new RoleBinding();
    }
    public RoleBindingBuilder role(BuiltinRoleRefOrCustomRoleRef role) {
        this.internal.role = role;
        return this;
    }
    
    public RoleBindingBuilder subject(com.grafana.foundation.cog.Builder<RoleBindingSubject> subject) {
    RoleBindingSubject subjectResource = subject.build();
        this.internal.subject = subjectResource;
        return this;
    }
    public RoleBinding build() {
        return this.internal;
    }
}
