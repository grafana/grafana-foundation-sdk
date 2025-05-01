// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.grafana.foundation.common.ScaleDimensionConfig;
import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.LineStyle;
import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.HideSeriesConfig;
import com.grafana.foundation.common.AxisPlacement;
import com.grafana.foundation.common.AxisColorMode;
import com.grafana.foundation.common.ScaleDistributionConfig;
import com.grafana.foundation.common.TextDimensionConfig;

public class ScatterSeriesConfigBuilder implements com.grafana.foundation.cog.Builder<ScatterSeriesConfig> {
    protected final ScatterSeriesConfig internal;
    
    public ScatterSeriesConfigBuilder() {
        this.internal = new ScatterSeriesConfig();
    }
    public ScatterSeriesConfigBuilder x(String x) {
        this.internal.x = x;
        return this;
    }
    
    public ScatterSeriesConfigBuilder y(String y) {
        this.internal.y = y;
        return this;
    }
    
    public ScatterSeriesConfigBuilder show(ScatterShow show) {
        this.internal.show = show;
        return this;
    }
    
    public ScatterSeriesConfigBuilder pointSize(com.grafana.foundation.cog.Builder<ScaleDimensionConfig> pointSize) {
    ScaleDimensionConfig pointSizeResource = pointSize.build();
        this.internal.pointSize = pointSizeResource;
        return this;
    }
    
    public ScatterSeriesConfigBuilder pointColor(com.grafana.foundation.cog.Builder<ColorDimensionConfig> pointColor) {
    ColorDimensionConfig pointColorResource = pointColor.build();
        this.internal.pointColor = pointColorResource;
        return this;
    }
    
    public ScatterSeriesConfigBuilder lineColor(com.grafana.foundation.cog.Builder<ColorDimensionConfig> lineColor) {
    ColorDimensionConfig lineColorResource = lineColor.build();
        this.internal.lineColor = lineColorResource;
        return this;
    }
    
    public ScatterSeriesConfigBuilder lineWidth(Integer lineWidth) {
        if (!(lineWidth >= 0)) {
            throw new IllegalArgumentException("lineWidth must be >= 0");
        }
        this.internal.lineWidth = lineWidth;
        return this;
    }
    
    public ScatterSeriesConfigBuilder lineStyle(com.grafana.foundation.cog.Builder<LineStyle> lineStyle) {
    LineStyle lineStyleResource = lineStyle.build();
        this.internal.lineStyle = lineStyleResource;
        return this;
    }
    
    public ScatterSeriesConfigBuilder label(VisibilityMode label) {
        this.internal.label = label;
        return this;
    }
    
    public ScatterSeriesConfigBuilder hideFrom(com.grafana.foundation.cog.Builder<HideSeriesConfig> hideFrom) {
    HideSeriesConfig hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisPlacement(AxisPlacement axisPlacement) {
        this.internal.axisPlacement = axisPlacement;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisColorMode(AxisColorMode axisColorMode) {
        this.internal.axisColorMode = axisColorMode;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisLabel(String axisLabel) {
        this.internal.axisLabel = axisLabel;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisWidth(Double axisWidth) {
        this.internal.axisWidth = axisWidth;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisSoftMin(Double axisSoftMin) {
        this.internal.axisSoftMin = axisSoftMin;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisSoftMax(Double axisSoftMax) {
        this.internal.axisSoftMax = axisSoftMax;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisGridShow(Boolean axisGridShow) {
        this.internal.axisGridShow = axisGridShow;
        return this;
    }
    
    public ScatterSeriesConfigBuilder scaleDistribution(com.grafana.foundation.cog.Builder<ScaleDistributionConfig> scaleDistribution) {
    ScaleDistributionConfig scaleDistributionResource = scaleDistribution.build();
        this.internal.scaleDistribution = scaleDistributionResource;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisCenteredZero(Boolean axisCenteredZero) {
        this.internal.axisCenteredZero = axisCenteredZero;
        return this;
    }
    
    public ScatterSeriesConfigBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public ScatterSeriesConfigBuilder labelValue(com.grafana.foundation.cog.Builder<TextDimensionConfig> labelValue) {
    TextDimensionConfig labelValueResource = labelValue.build();
        this.internal.labelValue = labelValueResource;
        return this;
    }
    
    public ScatterSeriesConfigBuilder axisBorderShow(Boolean axisBorderShow) {
        this.internal.axisBorderShow = axisBorderShow;
        return this;
    }
    public ScatterSeriesConfig build() {
        return this.internal;
    }
}
