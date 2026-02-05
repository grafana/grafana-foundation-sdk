// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;


public class PlacementBuilder implements com.grafana.foundation.cog.Builder<Placement> {
    protected final Placement internal;
    
    public PlacementBuilder() {
        this.internal = new Placement();
    }
    public PlacementBuilder top(Double top) {
        this.internal.top = top;
        return this;
    }
    
    public PlacementBuilder left(Double left) {
        this.internal.left = left;
        return this;
    }
    
    public PlacementBuilder right(Double right) {
        this.internal.right = right;
        return this;
    }
    
    public PlacementBuilder bottom(Double bottom) {
        this.internal.bottom = bottom;
        return this;
    }
    
    public PlacementBuilder width(Double width) {
        this.internal.width = width;
        return this;
    }
    
    public PlacementBuilder height(Double height) {
        this.internal.height = height;
        return this;
    }
    
    public PlacementBuilder rotation(Double rotation) {
        this.internal.rotation = rotation;
        return this;
    }
    public Placement build() {
        return this.internal;
    }
}
