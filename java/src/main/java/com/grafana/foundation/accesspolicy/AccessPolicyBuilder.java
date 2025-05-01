// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.accesspolicy;

import java.util.LinkedList;

public class AccessPolicyBuilder implements com.grafana.foundation.cog.Builder<AccessPolicy> {
    protected final AccessPolicy internal;
    
    public AccessPolicyBuilder() {
        this.internal = new AccessPolicy();
    }
    public AccessPolicyBuilder scope(com.grafana.foundation.cog.Builder<ResourceRef> scope) {
    ResourceRef scopeResource = scope.build();
        this.internal.scope = scopeResource;
        return this;
    }
    
    public AccessPolicyBuilder role(com.grafana.foundation.cog.Builder<RoleRef> role) {
    RoleRef roleResource = role.build();
        this.internal.role = roleResource;
        return this;
    }
    
    public AccessPolicyBuilder rules(com.grafana.foundation.cog.Builder<AccessRule> rule) {
		if (this.internal.rules == null) {
			this.internal.rules = new LinkedList<>();
		}
    AccessRule ruleResource = rule.build();
        this.internal.rules.add(ruleResource);
        return this;
    }
    public AccessPolicy build() {
        return this.internal;
    }
}
