// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import java.util.List;
import java.util.LinkedList;

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
    
    public CanvasOptionsRootBuilder elements(List<com.grafana.foundation.cog.Builder<CanvasElementOptions>> elements) {
        List<CanvasElementOptions> elementsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<CanvasElementOptions> r1 : elements) {
                CanvasElementOptions elementsDepth1 = r1.build();
                elementsResources.add(elementsDepth1); 
        }
        this.internal.elements = elementsResources;
        return this;
    }
    public CanvasOptionsRoot build() {
        return this.internal;
    }
}
