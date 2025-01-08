// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;

public class YAxisConfigBuilder implements com.grafana.foundation.cog.Builder<YAxisConfig> {
    protected final YAxisConfig internal;
    
    public YAxisConfigBuilder() {
        this.internal = new YAxisConfig();
    }
    public YAxisConfigBuilder unit(String unit) {
    this.internal.unit = unit;
        return this;
    }
    
    public YAxisConfigBuilder reverse(Boolean reverse) {
    this.internal.reverse = reverse;
        return this;
    }
    
    public YAxisConfigBuilder decimals(Float decimals) {
    this.internal.decimals = decimals;
        return this;
    }
    
    public YAxisConfigBuilder min(Float min) {
    this.internal.min = min;
        return this;
    }
    
    public YAxisConfigBuilder axisPlacement(AxisPlacement axisPlacement) {
    this.internal.axisPlacement = axisPlacement;
        return this;
    }
    
    public YAxisConfigBuilder axisColorMode(AxisColorMode axisColorMode) {
    this.internal.axisColorMode = axisColorMode;
        return this;
    }
    
    public YAxisConfigBuilder axisLabel(String axisLabel) {
    this.internal.axisLabel = axisLabel;
        return this;
    }
    
    public YAxisConfigBuilder axisWidth(Double axisWidth) {
    this.internal.axisWidth = axisWidth;
        return this;
    }
    
    public YAxisConfigBuilder axisSoftMin(Double axisSoftMin) {
    this.internal.axisSoftMin = axisSoftMin;
        return this;
    }
    
    public YAxisConfigBuilder axisSoftMax(Double axisSoftMax) {
    this.internal.axisSoftMax = axisSoftMax;
        return this;
    }
    
    public YAxisConfigBuilder axisGridShow(Boolean axisGridShow) {
    this.internal.axisGridShow = axisGridShow;
        return this;
    }
    
    public YAxisConfigBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
    this.internal.scaleDistribution = scaleDistribution.build();
        return this;
    }
    
    public YAxisConfigBuilder max(Float max) {
    this.internal.max = max;
        return this;
    }
    
    public YAxisConfigBuilder axisCenteredZero(Boolean axisCenteredZero) {
    this.internal.axisCenteredZero = axisCenteredZero;
        return this;
    }
    public YAxisConfig build() {
        return this.internal;
    }
}
