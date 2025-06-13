// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.ScaleDimensionConfig;

public class CanvasConnectionBuilder implements com.grafana.foundation.cog.Builder<CanvasConnection> {
    protected final CanvasConnection internal;
    
    public CanvasConnectionBuilder() {
        this.internal = new CanvasConnection();
    }
    public CanvasConnectionBuilder source(com.grafana.foundation.cog.Builder<ConnectionCoordinates> source) {
    ConnectionCoordinates sourceResource = source.build();
        this.internal.source = sourceResource;
        return this;
    }
    
    public CanvasConnectionBuilder target(com.grafana.foundation.cog.Builder<ConnectionCoordinates> target) {
    ConnectionCoordinates targetResource = target.build();
        this.internal.target = targetResource;
        return this;
    }
    
    public CanvasConnectionBuilder targetName(String targetName) {
        this.internal.targetName = targetName;
        return this;
    }
    
    public CanvasConnectionBuilder color(com.grafana.foundation.cog.Builder<ColorDimensionConfig> color) {
    ColorDimensionConfig colorResource = color.build();
        this.internal.color = colorResource;
        return this;
    }
    
    public CanvasConnectionBuilder size(com.grafana.foundation.cog.Builder<ScaleDimensionConfig> size) {
    ScaleDimensionConfig sizeResource = size.build();
        this.internal.size = sizeResource;
        return this;
    }
    public CanvasConnection build() {
        return this.internal;
    }
}
