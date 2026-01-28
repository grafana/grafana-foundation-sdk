// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;


public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder inlineEditing(Boolean inlineEditing) {
        this.internal.inlineEditing = inlineEditing;
        return this;
    }
    
    public OptionsBuilder showAdvancedTypes(Boolean showAdvancedTypes) {
        this.internal.showAdvancedTypes = showAdvancedTypes;
        return this;
    }
    
    public OptionsBuilder panZoom(Boolean panZoom) {
        this.internal.panZoom = panZoom;
        return this;
    }
    
    public OptionsBuilder root(com.grafana.foundation.cog.Builder<CanvasOptionsRoot> root) {
    CanvasOptionsRoot rootResource = root.build();
        this.internal.root = rootResource;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
