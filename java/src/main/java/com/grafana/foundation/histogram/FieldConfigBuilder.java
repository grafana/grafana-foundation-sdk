// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.histogram;

import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.HideSeriesConfig;
import com.grafana.foundation.common.GraphGradientMode;

public class FieldConfigBuilder implements com.grafana.foundation.cog.Builder<FieldConfig> {
    protected final FieldConfig internal;
    
    public FieldConfigBuilder() {
        this.internal = new FieldConfig();
    }
    public FieldConfigBuilder lineWidth(Integer lineWidth) {
        if (!(lineWidth <= 10)) {
            throw new IllegalArgumentException("lineWidth must be <= 10");
        }
        this.internal.lineWidth = lineWidth;
        return this;
    }
    
    public FieldConfigBuilder fillOpacity(Integer fillOpacity) {
        if (!(fillOpacity <= 100)) {
            throw new IllegalArgumentException("fillOpacity must be <= 100");
        }
        this.internal.fillOpacity = fillOpacity;
        return this;
    }
    
    public FieldConfigBuilder axisPlacement(AxisPlacement axisPlacement) {
        this.internal.axisPlacement = axisPlacement;
        return this;
    }
    
    public FieldConfigBuilder axisColorMode(AxisColorMode axisColorMode) {
        this.internal.axisColorMode = axisColorMode;
        return this;
    }
    
    public FieldConfigBuilder axisLabel(String axisLabel) {
        this.internal.axisLabel = axisLabel;
        return this;
    }
    
    public FieldConfigBuilder axisWidth(Double axisWidth) {
        this.internal.axisWidth = axisWidth;
        return this;
    }
    
    public FieldConfigBuilder axisSoftMin(Double axisSoftMin) {
        this.internal.axisSoftMin = axisSoftMin;
        return this;
    }
    
    public FieldConfigBuilder axisSoftMax(Double axisSoftMax) {
        this.internal.axisSoftMax = axisSoftMax;
        return this;
    }
    
    public FieldConfigBuilder axisGridShow(Boolean axisGridShow) {
        this.internal.axisGridShow = axisGridShow;
        return this;
    }
    
    public FieldConfigBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
    ScaleDistributionConfig scaleDistributionResource = scaleDistribution.build();
        this.internal.scaleDistribution = scaleDistributionResource;
        return this;
    }
    
    public FieldConfigBuilder axisCenteredZero(Boolean axisCenteredZero) {
        this.internal.axisCenteredZero = axisCenteredZero;
        return this;
    }
    
    public FieldConfigBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
    HideSeriesConfig hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
        return this;
    }
    
    public FieldConfigBuilder gradientMode(GraphGradientMode gradientMode) {
        this.internal.gradientMode = gradientMode;
        return this;
    }
    
    public FieldConfigBuilder axisBorderShow(Boolean axisBorderShow) {
        this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
    public FieldConfig build() {
        return this.internal;
    }
}
