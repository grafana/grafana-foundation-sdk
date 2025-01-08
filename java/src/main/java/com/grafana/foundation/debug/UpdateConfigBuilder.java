// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.debug;


public class UpdateConfigBuilder implements com.grafana.foundation.cog.Builder<UpdateConfig> {
    protected final UpdateConfig internal;
    
    public UpdateConfigBuilder() {
        this.internal = new UpdateConfig();
    }
    public UpdateConfigBuilder render(Boolean render) {
    this.internal.render = render;
        return this;
    }
    
    public UpdateConfigBuilder dataChanged(Boolean dataChanged) {
    this.internal.dataChanged = dataChanged;
        return this;
    }
    
    public UpdateConfigBuilder schemaChanged(Boolean schemaChanged) {
    this.internal.schemaChanged = schemaChanged;
        return this;
    }
    public UpdateConfig build() {
        return this.internal;
    }
}
