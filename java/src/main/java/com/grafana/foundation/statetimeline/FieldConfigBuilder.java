// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.statetimeline;

import com.grafana.foundation.common.HideSeriesConfig;

public class FieldConfigBuilder implements com.grafana.foundation.cog.Builder<FieldConfig> {
    protected final FieldConfig internal;
    
    public FieldConfigBuilder() {
        this.internal = new FieldConfig();
    }
    public FieldConfigBuilder lineWidth(Integer lineWidth) {
        if (!(lineWidth <= 10)) {
            throw new IllegalArgumentException("lineWidth must be <= 10");
        }
        this.internal.lineWidth = lineWidth;
        return this;
    }
    
    public FieldConfigBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
    HideSeriesConfig hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
        return this;
    }
    
    public FieldConfigBuilder fillOpacity(Integer fillOpacity) {
        if (!(fillOpacity <= 100)) {
            throw new IllegalArgumentException("fillOpacity must be <= 100");
        }
        this.internal.fillOpacity = fillOpacity;
        return this;
    }
    public FieldConfig build() {
        return this.internal;
    }
}
