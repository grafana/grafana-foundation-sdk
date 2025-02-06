// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class AxisConfigBuilder implements com.grafana.foundation.cog.Builder<AxisConfig> {
    protected final AxisConfig internal;
    
    public AxisConfigBuilder() {
        this.internal = new AxisConfig();
    }
    public AxisConfigBuilder axisPlacement(AxisPlacement axisPlacement) {
        this.internal.axisPlacement = axisPlacement;
        return this;
    }
    
    public AxisConfigBuilder axisColorMode(AxisColorMode axisColorMode) {
        this.internal.axisColorMode = axisColorMode;
        return this;
    }
    
    public AxisConfigBuilder axisLabel(String axisLabel) {
        this.internal.axisLabel = axisLabel;
        return this;
    }
    
    public AxisConfigBuilder axisWidth(Double axisWidth) {
        this.internal.axisWidth = axisWidth;
        return this;
    }
    
    public AxisConfigBuilder axisSoftMin(Double axisSoftMin) {
        this.internal.axisSoftMin = axisSoftMin;
        return this;
    }
    
    public AxisConfigBuilder axisSoftMax(Double axisSoftMax) {
        this.internal.axisSoftMax = axisSoftMax;
        return this;
    }
    
    public AxisConfigBuilder axisGridShow(Boolean axisGridShow) {
        this.internal.axisGridShow = axisGridShow;
        return this;
    }
    
    public AxisConfigBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
        this.internal.scaleDistribution = scaleDistribution.build();
        return this;
    }
    
    public AxisConfigBuilder axisCenteredZero(Boolean axisCenteredZero) {
        this.internal.axisCenteredZero = axisCenteredZero;
        return this;
    }
    public AxisConfig build() {
        return this.internal;
    }
}
