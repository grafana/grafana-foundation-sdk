// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class ActionVariableBuilder implements com.grafana.foundation.cog.Builder<ActionVariable> {
    protected final ActionVariable internal;
    
    public ActionVariableBuilder() {
        this.internal = new ActionVariable();
    }
    public ActionVariableBuilder key(String key) {
        this.internal.key = key;
        return this;
    }
    
    public ActionVariableBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    public ActionVariable build() {
        return this.internal;
    }
}
