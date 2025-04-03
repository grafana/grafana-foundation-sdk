// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.accesspolicy;


public class AccessRuleBuilder implements com.grafana.foundation.cog.Builder<AccessRule> {
    protected final AccessRule internal;
    
    public AccessRuleBuilder() {
        this.internal = new AccessRule();
    }
    public AccessRuleBuilder kind(String kind) {
        this.internal.kind = kind;
        return this;
    }
    
    public AccessRuleBuilder verb(String verb) {
        this.internal.verb = verb;
        return this;
    }
    
    public AccessRuleBuilder target(String target) {
        this.internal.target = target;
        return this;
    }
    public AccessRule build() {
        return this.internal;
    }
}
