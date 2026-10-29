// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;


public class RuleRefBuilder implements com.grafana.foundation.cog.Builder<RuleRef> {
    protected final RuleRef internal;
    
    public RuleRefBuilder() {
        this.internal = new RuleRef();
    }
    public RuleRefBuilder name(String name) {
        if (!name.matches("^[a-zA-Z0-9_-]+$")) {
            throw new IllegalArgumentException("name must match regex ^[a-zA-Z0-9_-]+$");
        }
        this.internal.name = name;
        return this;
    }
    public RuleRef build() {
        return this.internal;
    }
}
