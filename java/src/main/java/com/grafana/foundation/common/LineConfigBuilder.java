// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class LineConfigBuilder implements com.grafana.foundation.cog.Builder<LineConfig> {
    protected final LineConfig internal;
    
    public LineConfigBuilder() {
        this.internal = new LineConfig();
    }
    public LineConfigBuilder lineColor(String lineColor) {
        this.internal.lineColor = lineColor;
        return this;
    }
    
    public LineConfigBuilder lineWidth(Double lineWidth) {
        this.internal.lineWidth = lineWidth;
        return this;
    }
    
    public LineConfigBuilder lineInterpolation(LineInterpolation lineInterpolation) {
        this.internal.lineInterpolation = lineInterpolation;
        return this;
    }
    
    public LineConfigBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle) {
        this.internal.lineStyle = lineStyle.build();
        return this;
    }
    
    public LineConfigBuilder spanNulls(BoolOrFloat64 spanNulls) {
        this.internal.spanNulls = spanNulls;
        return this;
    }
    public LineConfig build() {
        return this.internal;
    }
}
