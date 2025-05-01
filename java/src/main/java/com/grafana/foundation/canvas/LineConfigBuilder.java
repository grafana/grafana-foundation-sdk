// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.grafana.foundation.common.ColorDimensionConfig;

public class LineConfigBuilder implements com.grafana.foundation.cog.Builder<LineConfig> {
    protected final LineConfig internal;
    
    public LineConfigBuilder() {
        this.internal = new LineConfig();
    }
    public LineConfigBuilder color(com.grafana.foundation.cog.Builder<ColorDimensionConfig> color) {
    ColorDimensionConfig colorResource = color.build();
        this.internal.color = colorResource;
        return this;
    }
    
    public LineConfigBuilder width(Double width) {
        this.internal.width = width;
        return this;
    }
    
    public LineConfigBuilder radius(Double radius) {
        this.internal.radius = radius;
        return this;
    }
    public LineConfig build() {
        return this.internal;
    }
}
