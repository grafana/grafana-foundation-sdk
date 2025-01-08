// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import java.util.List;

public class LineStyleBuilder implements com.grafana.foundation.cog.Builder<LineStyle> {
    protected final LineStyle internal;
    
    public LineStyleBuilder() {
        this.internal = new LineStyle();
    }
    public LineStyleBuilder fill(LineStyleFill fill) {
    this.internal.fill = fill;
        return this;
    }
    
    public LineStyleBuilder dash(List<Double> dash) {
    this.internal.dash = dash;
        return this;
    }
    public LineStyle build() {
        return this.internal;
    }
}
