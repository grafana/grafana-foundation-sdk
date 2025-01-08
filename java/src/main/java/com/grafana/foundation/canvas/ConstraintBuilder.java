// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;


public class ConstraintBuilder implements com.grafana.foundation.cog.Builder<Constraint> {
    protected final Constraint internal;
    
    public ConstraintBuilder() {
        this.internal = new Constraint();
    }
    public ConstraintBuilder horizontal(HorizontalConstraint horizontal) {
    this.internal.horizontal = horizontal;
        return this;
    }
    
    public ConstraintBuilder vertical(VerticalConstraint vertical) {
    this.internal.vertical = vertical;
        return this;
    }
    public Constraint build() {
        return this.internal;
    }
}
