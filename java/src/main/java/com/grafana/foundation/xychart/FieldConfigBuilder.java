// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.grafana.foundation.common.HideSeriesConfig;
import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.LineStyle;

public class FieldConfigBuilder implements com.grafana.foundation.cog.Builder<FieldConfig> {
    protected final FieldConfig internal;
    
    public FieldConfigBuilder() {
        this.internal = new FieldConfig();
    }
    public FieldConfigBuilder show(XYShowMode show) {
        this.internal.show = show;
        return this;
    }
    
    public FieldConfigBuilder pointSize(com.grafana.foundation.cog.Builder<XychartFieldConfigPointSize> pointSize) {
    XychartFieldConfigPointSize pointSizeResource = pointSize.build();
        this.internal.pointSize = pointSizeResource;
        return this;
    }
    
    public FieldConfigBuilder pointShape(PointShape pointShape) {
        this.internal.pointShape = pointShape;
        return this;
    }
    
    public FieldConfigBuilder pointStrokeWidth(Integer pointStrokeWidth) {
        if (!(pointStrokeWidth >= 0)) {
            throw new IllegalArgumentException("pointStrokeWidth must be >= 0");
        }
        this.internal.pointStrokeWidth = pointStrokeWidth;
        return this;
    }
    
    public FieldConfigBuilder fillOpacity(Integer fillOpacity) {
        if (!(fillOpacity <= 100)) {
            throw new IllegalArgumentException("fillOpacity must be <= 100");
        }
        this.internal.fillOpacity = fillOpacity;
        return this;
    }
    
    public FieldConfigBuilder lineWidth(Integer lineWidth) {
        if (!(lineWidth >= 0)) {
            throw new IllegalArgumentException("lineWidth must be >= 0");
        }
        this.internal.lineWidth = lineWidth;
        return this;
    }
    
    public FieldConfigBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
    HideSeriesConfig hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
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
    
    public FieldConfigBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle) {
    LineStyle lineStyleResource = lineStyle.build();
        this.internal.lineStyle = lineStyleResource;
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
