// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class FieldColorBuilder implements com.grafana.foundation.cog.Builder<FieldColor> {
    protected final FieldColor internal;
    
    public FieldColorBuilder() {
        this.internal = new FieldColor();
    }
    public FieldColorBuilder mode(FieldColorModeId mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public FieldColorBuilder fixedColor(String fixedColor) {
    this.internal.fixedColor = fixedColor;
        return this;
    }
    
    public FieldColorBuilder seriesBy(FieldColorSeriesByMode seriesBy) {
    this.internal.seriesBy = seriesBy;
        return this;
    }
    public FieldColor build() {
        return this.internal;
    }
}
