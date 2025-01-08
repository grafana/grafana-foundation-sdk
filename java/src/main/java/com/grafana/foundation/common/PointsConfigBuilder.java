// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class PointsConfigBuilder implements com.grafana.foundation.cog.Builder<PointsConfig> {
    protected final PointsConfig internal;
    
    public PointsConfigBuilder() {
        this.internal = new PointsConfig();
    }
    public PointsConfigBuilder showPoints(VisibilityMode showPoints) {
    this.internal.showPoints = showPoints;
        return this;
    }
    
    public PointsConfigBuilder pointSize(Double pointSize) {
    this.internal.pointSize = pointSize;
        return this;
    }
    
    public PointsConfigBuilder pointColor(String pointColor) {
    this.internal.pointColor = pointColor;
        return this;
    }
    
    public PointsConfigBuilder pointSymbol(String pointSymbol) {
    this.internal.pointSymbol = pointSymbol;
        return this;
    }
    public PointsConfig build() {
        return this.internal;
    }
}
