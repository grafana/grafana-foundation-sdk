// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.grafana.foundation.common.ColorDimensionConfig;

public class LineConfigBuilder implements com.grafana.foundation.cog.Builder<LineConfig> {
    protected final LineConfig internal;
    
    public LineConfigBuilder() {
        this.internal = new LineConfig();
    }
    public LineConfigBuilder color(com.grafana.foundation.cog.Builder<ColorDimensionConfig> color) {
    this.internal.color = color.build();
        return this;
    }
    
    public LineConfigBuilder width(Double width) {
    this.internal.width = width;
        return this;
    }
    public LineConfig build() {
        return this.internal;
    }
}
