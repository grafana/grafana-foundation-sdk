// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.ScaleDimensionConfig;
import java.util.List;

public class CanvasConnectionBuilder implements com.grafana.foundation.cog.Builder<CanvasConnection> {
    protected final CanvasConnection internal;
    
    public CanvasConnectionBuilder() {
        this.internal = new CanvasConnection();
    }
    public CanvasConnectionBuilder source(com.grafana.foundation.cog.Builder<ConnectionCoordinates> source) {
    this.internal.source = source.build();
        return this;
    }
    
    public CanvasConnectionBuilder target(com.grafana.foundation.cog.Builder<ConnectionCoordinates> target) {
    this.internal.target = target.build();
        return this;
    }
    
    public CanvasConnectionBuilder targetName(String targetName) {
    this.internal.targetName = targetName;
        return this;
    }
    
    public CanvasConnectionBuilder path(ConnectionPath path) {
    this.internal.path = path;
        return this;
    }
    
    public CanvasConnectionBuilder color(com.grafana.foundation.cog.Builder<ColorDimensionConfig> color) {
    this.internal.color = color.build();
        return this;
    }
    
    public CanvasConnectionBuilder size(com.grafana.foundation.cog.Builder<ScaleDimensionConfig> size) {
    this.internal.size = size.build();
        return this;
    }
    
    public CanvasConnectionBuilder vertices(com.grafana.foundation.cog.Builder<List<ConnectionCoordinates>> vertices) {
    this.internal.vertices = vertices.build();
        return this;
    }
    
    public CanvasConnectionBuilder sourceOriginal(com.grafana.foundation.cog.Builder<ConnectionCoordinates> sourceOriginal) {
    this.internal.sourceOriginal = sourceOriginal.build();
        return this;
    }
    
    public CanvasConnectionBuilder targetOriginal(com.grafana.foundation.cog.Builder<ConnectionCoordinates> targetOriginal) {
    this.internal.targetOriginal = targetOriginal.build();
        return this;
    }
    public CanvasConnection build() {
        return this.internal;
    }
}
