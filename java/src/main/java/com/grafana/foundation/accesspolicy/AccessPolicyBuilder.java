// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.accesspolicy;

import java.util.LinkedList;

public class AccessPolicyBuilder implements com.grafana.foundation.cog.Builder<AccessPolicy> {
    protected final AccessPolicy internal;
    
    public AccessPolicyBuilder() {
        this.internal = new AccessPolicy();
    }
    public AccessPolicyBuilder scope(com.grafana.foundation.cog.Builder<ResourceRef> scope) {
    this.internal.scope = scope.build();
        return this;
    }
    
    public AccessPolicyBuilder role(com.grafana.foundation.cog.Builder<RoleRef> role) {
    this.internal.role = role.build();
        return this;
    }
    
    public AccessPolicyBuilder rules(com.grafana.foundation.cog.Builder<AccessRule> rule) {
		if (this.internal.rules == null) {
			this.internal.rules = new LinkedList<>();
		}
    this.internal.rules.add(rule.build());
        return this;
    }
    public AccessPolicy build() {
        return this.internal;
    }
}
