// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import java.util.List;

public class CanvasOptionsRootBuilder implements com.grafana.foundation.cog.Builder<CanvasOptionsRoot> {
    protected final CanvasOptionsRoot internal;
    
    public CanvasOptionsRootBuilder() {
        this.internal = new CanvasOptionsRoot();
    this.internal.type = "frame";
    }
    public CanvasOptionsRootBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public CanvasOptionsRootBuilder elements(com.grafana.foundation.cog.Builder<List<CanvasElementOptions>> elements) {
    this.internal.elements = elements.build();
        return this;
    }
    public CanvasOptionsRoot build() {
        return this.internal;
    }
}
