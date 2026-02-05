// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.nodegraph;


public class ArcOptionBuilder implements com.grafana.foundation.cog.Builder<ArcOption> {
    protected final ArcOption internal;
    
    public ArcOptionBuilder() {
        this.internal = new ArcOption();
    }
    public ArcOptionBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public ArcOptionBuilder color(String color) {
        this.internal.color = color;
        return this;
    }
    public ArcOption build() {
        return this.internal;
    }
}
