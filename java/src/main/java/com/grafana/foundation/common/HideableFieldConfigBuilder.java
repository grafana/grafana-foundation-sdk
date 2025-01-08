// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class HideableFieldConfigBuilder implements com.grafana.foundation.cog.Builder<HideableFieldConfig> {
    protected final HideableFieldConfig internal;
    
    public HideableFieldConfigBuilder() {
        this.internal = new HideableFieldConfig();
    }
    public HideableFieldConfigBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
    this.internal.hideFrom = hideFrom.build();
        return this;
    }
    public HideableFieldConfig build() {
        return this.internal;
    }
}
