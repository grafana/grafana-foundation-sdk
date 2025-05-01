// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import java.util.List;
import java.util.LinkedList;

public class CanvasElementOptionsBuilder implements com.grafana.foundation.cog.Builder<CanvasElementOptions> {
    protected final CanvasElementOptions internal;
    
    public CanvasElementOptionsBuilder() {
        this.internal = new CanvasElementOptions();
    }
    public CanvasElementOptionsBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public CanvasElementOptionsBuilder type(String type) {
        this.internal.type = type;
        return this;
    }
    
    public CanvasElementOptionsBuilder config(Object config) {
        this.internal.config = config;
        return this;
    }
    
    public CanvasElementOptionsBuilder constraint(com.grafana.foundation.cog.Builder<Constraint> constraint) {
    Constraint constraintResource = constraint.build();
        this.internal.constraint = constraintResource;
        return this;
    }
    
    public CanvasElementOptionsBuilder placement(com.grafana.foundation.cog.Builder<Placement> placement) {
    Placement placementResource = placement.build();
        this.internal.placement = placementResource;
        return this;
    }
    
    public CanvasElementOptionsBuilder background(com.grafana.foundation.cog.Builder<BackgroundConfig> background) {
    BackgroundConfig backgroundResource = background.build();
        this.internal.background = backgroundResource;
        return this;
    }
    
    public CanvasElementOptionsBuilder border(com.grafana.foundation.cog.Builder<LineConfig> border) {
    LineConfig borderResource = border.build();
        this.internal.border = borderResource;
        return this;
    }
    
    public CanvasElementOptionsBuilder connections(List<com.grafana.foundation.cog.Builder<CanvasConnection>> connections) {
        List<CanvasConnection> connectionsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<CanvasConnection> r1 : connections) {
                CanvasConnection connectionsDepth1 = r1.build();
                connectionsResources.add(connectionsDepth1); 
        }
        this.internal.connections = connectionsResources;
        return this;
    }
    public CanvasElementOptions build() {
        return this.internal;
    }
}
