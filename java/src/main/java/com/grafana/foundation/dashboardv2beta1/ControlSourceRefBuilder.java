// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class ControlSourceRefBuilder implements com.grafana.foundation.cog.Builder<ControlSourceRef> {
    protected final ControlSourceRef internal;
    
    public ControlSourceRefBuilder() {
        this.internal = new ControlSourceRef();
        this.internal.type = "datasource";
    }
    public ControlSourceRefBuilder group(String group) {
        this.internal.group = group;
        return this;
    }
    public ControlSourceRef build() {
        return this.internal;
    }
}
