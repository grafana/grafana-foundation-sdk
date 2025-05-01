// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.ScaleDimensionConfig;
import java.util.List;
import java.util.LinkedList;

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
    
    public CanvasConnectionBuilder path(ConnectionPath path) {
        this.internal.path = path;
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
    
    public CanvasConnectionBuilder vertices(List<com.grafana.foundation.cog.Builder<ConnectionCoordinates>> vertices) {
        List<ConnectionCoordinates> verticesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<ConnectionCoordinates> r1 : vertices) {
                ConnectionCoordinates verticesDepth1 = r1.build();
                verticesResources.add(verticesDepth1); 
        }
        this.internal.vertices = verticesResources;
        return this;
    }
    
    public CanvasConnectionBuilder sourceOriginal(com.grafana.foundation.cog.Builder<ConnectionCoordinates> sourceOriginal) {
    ConnectionCoordinates sourceOriginalResource = sourceOriginal.build();
        this.internal.sourceOriginal = sourceOriginalResource;
        return this;
    }
    
    public CanvasConnectionBuilder targetOriginal(com.grafana.foundation.cog.Builder<ConnectionCoordinates> targetOriginal) {
    ConnectionCoordinates targetOriginalResource = targetOriginal.build();
        this.internal.targetOriginal = targetOriginalResource;
        return this;
    }
    public CanvasConnection build() {
        return this.internal;
    }
}
