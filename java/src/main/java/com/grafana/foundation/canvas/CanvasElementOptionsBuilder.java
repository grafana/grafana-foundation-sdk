// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import java.util.List;

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
    this.internal.constraint = constraint.build();
        return this;
    }
    
    public CanvasElementOptionsBuilder placement(com.grafana.foundation.cog.Builder<Placement> placement) {
    this.internal.placement = placement.build();
        return this;
    }
    
    public CanvasElementOptionsBuilder background(com.grafana.foundation.cog.Builder<BackgroundConfig> background) {
    this.internal.background = background.build();
        return this;
    }
    
    public CanvasElementOptionsBuilder border(com.grafana.foundation.cog.Builder<LineConfig> border) {
    this.internal.border = border.build();
        return this;
    }
    
    public CanvasElementOptionsBuilder connections(com.grafana.foundation.cog.Builder<List<CanvasConnection>> connections) {
    this.internal.connections = connections.build();
        return this;
    }
    public CanvasElementOptions build() {
        return this.internal;
    }
}
